package main

import (
	"context"
	"crypto/ecdsa"
	exampleabi "event-poc/abi/example"
	"event-poc/intercept"
	"flag"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var rpcURL = flag.String("rpc", "http://localhost:8545", "Ethereum RPC URL")
var wsURL = flag.String("ws", "ws://localhost:8546", "Ethereum WebSocket URL")
var hexpk = flag.String("pk", "", "Private key")
var withHTTPLog = flag.Bool("httplog", false, "Log HTTP requests and responses")
var withWSLog = flag.Bool("wslog", false, "Log WebSocket messages")

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	flag.Parse()

	ctx := context.Background()

	// Initialize client
	options := []rpc.ClientOption{}
	if *withHTTPLog {
		options = append(options, rpc.WithHTTPClient(&http.Client{
			Transport: &intercept.Transport{},
		}))
	}
	if *withWSLog {
		options = append(options, rpc.WithWebsocketDialer(websocket.Dialer{}))
	}
	rpcClient, err := rpc.DialOptions(
		ctx,
		*rpcURL,
		options...,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Ethereum RPC")
	}
	defer rpcClient.Close()
	ethRPC := ethclient.NewClient(rpcClient)
	wsClient, err := rpc.DialOptions(
		ctx,
		*wsURL,
		options...,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Ethereum WebSocket")
	}
	ethWS := ethclient.NewClient(wsClient)
	chainID, err := ethRPC.ChainID(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get chain ID")
	}
	pk, err := crypto.HexToECDSA(*hexpk)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse private key")
	}
	aClient := &Authenticator{
		pk:      pk,
		chainID: chainID,
		client:  ethRPC,
	}

	// Deploy contract
	contractAddress, err := aClient.authAndWaitDeployed(
		ctx,
		func(auth *bind.TransactOpts) (*types.Transaction, error) {
			_, tx, _, err := exampleabi.DeployEventExample(auth, ethRPC)
			return tx, err
		},
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to deploy contract")
	}
	log.Info().Str("contract", contractAddress.Hex()).Msg("Contract deployed")
	// Create type-safe binding to contract
	eventExampleContract, err := exampleabi.NewEventExample(contractAddress, ethWS)
	if err != nil {
		panic(err)
	}

	// Trigger event
	if err = aClient.authAndWaitMined(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return eventExampleContract.LaunchEvent(auth, "Hello, world!")
	}); err != nil {
		log.Fatal().Err(err).Msg("Failed to trigger event")
	}

	// Subscribe to event logs
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ready := make(chan struct{})
	done := make(chan struct{})

	// Method 1: Using SubscribeFilterLogs from go-ethereum
	// go watchWithEthereum(ctx, contractAddress, ws, eventExampleContract, ready, done)

	// Method 2: Using WatchNewEvent from the contract binding
	// go watchWithContractBinding(ctx, eventExampleContract, ready, done)

	// Method 3: Using raw JSON-RPC calls
	go watchWithCall(ctx, contractAddress, ethWS, eventExampleContract, ready, done)

	select {
	case <-ready:
	case <-ctx.Done():
		log.Fatal().Err(ctx.Err()).Msg("Timeout waiting for subscription")
	}

	// Trigger event
	if err = aClient.authAndWaitMined(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return eventExampleContract.LaunchEvent(auth, "Hello, world!")
	}); err != nil {
		log.Fatal().Err(err).Msg("Failed to trigger event")
	}

	select {
	case <-done:
	case <-ctx.Done():
		log.Fatal().Err(ctx.Err()).Msg("Timeout waiting for event")
	}
}

type Authenticator struct {
	pk      *ecdsa.PrivateKey
	chainID *big.Int
	client  *ethclient.Client
}

func (a *Authenticator) authenticate(
	ctx context.Context,
) (*bind.TransactOpts, error) {
	gasPrice, err := a.client.SuggestGasPrice(ctx)
	if err != nil {
		slog.Error("Failed to suggest gas price", slog.String("err", err.Error()))
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(a.pk, a.chainID)
	if err != nil {
		return nil, err
	}
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}

func (a *Authenticator) authAndWaitDeployed(
	ctx context.Context,
	act func(auth *bind.TransactOpts) (*types.Transaction, error),
) (common.Address, error) {
	auth, err := a.authenticate(ctx)
	if err != nil {
		return common.Address{}, err
	}
	tx, err := act(auth)
	if err != nil {
		return common.Address{}, err
	}
	return bind.WaitDeployed(ctx, a.client, tx)
}

func (a *Authenticator) authAndWaitMined(
	ctx context.Context,
	act func(auth *bind.TransactOpts) (*types.Transaction, error),
) error {
	auth, err := a.authenticate(ctx)
	if err != nil {
		return err
	}
	tx, err := act(auth)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, a.client, tx)
	if err != nil {
		return err
	}
	return nil
}

func watchWithEthereum(
	ctx context.Context,
	contractAddress common.Address,
	ws *ethclient.Client,
	eventExampleContract *exampleabi.EventExample,
	ready, done chan struct{},
) {
	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	sub, err := ws.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	close(ready)

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		case l := <-logs:
			// Note: using the RPC client for parsing logs instead of WS doesn't matter.
			e, err := eventExampleContract.ParseNewEvent(l)
			if err != nil {
				log.Err(err).Msg("Failed to parse event")
				return
			}

			fmt.Println("Event received:", e.Message)
			close(done)
		}
	}
}

func watchWithContractBinding(
	ctx context.Context,
	eventExampleContract *exampleabi.EventExample,
	ready, done chan struct{},
) {
	events := make(chan *exampleabi.EventExampleNewEvent)
	sub, err := eventExampleContract.WatchNewEvent(&bind.WatchOpts{
		Context: ctx,
	}, events)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	close(ready)

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		case e := <-events:
			fmt.Println("Event received:", e.Message)
			close(done)
		}
	}
}

func watchWithCall(
	ctx context.Context,
	contractAddress common.Address,
	ws *ethclient.Client,
	eventExampleContract *exampleabi.EventExample,
	ready, done chan struct{},
) {
	ethWS := ws.Client()
	//HERE: Replace interface{} by types.Log and the parsing will fail.
	logs := make(chan interface{}, 1)
	sub, err := ethWS.EthSubscribe(ctx, logs, "logs", map[string]interface{}{
		"address": []string{contractAddress.Hex()},
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to subscribe to newHeads")
	}
	defer sub.Unsubscribe()

	close(ready)

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		case l := <-logs:
			fmt.Println("Log received:", l)
			// e, err := eventExampleContract.ParseNewEvent(l)
			// if err != nil {
			// 	log.Err(err).Msg("Failed to parse event")
			// 	return
			// }

			// fmt.Println("Event received:", e.Message)
			close(done)
		}
	}
}
