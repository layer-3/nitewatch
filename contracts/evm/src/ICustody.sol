// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface ICustody {
    event Deposited(address indexed user, address indexed token, uint256 amount);
    event Withdrawn(address indexed user, address indexed token, uint256 amount, uint256 nonce);

    function deposit(address token, uint256 amount) external payable;

    function withdraw(
        address user,
        address token,
        uint256 amount,
        uint256 nonce,
        uint256 expiry,
        bytes[] calldata signatures
    ) external;
}
