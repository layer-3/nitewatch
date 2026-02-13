// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IDeposit {
    // ---- errors ----
    error MsgValueMismatch();
    error NonZeroMsgValueForERC20();

    /// @notice Emitted when a user deposits funds into custody.
    event Deposited(address indexed user, address indexed token, uint256 amount);

    /// @notice Deposit ERC20 tokens or native ETH into custody.
    /// @dev For native ETH deposits, `token` must be address(0) and `msg.value` must equal `amount`.
    /// @param token The ERC20 token address, or address(0) for native ETH.
    /// @param amount The amount to deposit.
    function deposit(address token, uint256 amount) external payable;
}
