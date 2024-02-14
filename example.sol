// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EventExample {
    event NewEvent(string _message);

    function launchEvent(string memory _message) public {
        emit NewEvent(_message);
    }
}
