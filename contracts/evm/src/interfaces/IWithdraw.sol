// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

interface IWithdraw {
    error WithdrawalAlreadyExists();
    error WithdrawalNotFound();
    error WithdrawalAlreadyFinalized();
    error InsufficientLiquidity();
    error ETHTransferFailed();

    /// @notice Emitted when a withdrawal is initiated by NeoDAX.
    event WithdrawStarted(
        bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce
    );

    /// @notice Emitted when a withdrawal is finalized by Nitewatch.
    event WithdrawFinalized(bytes32 indexed withdrawalId, bool success);

    /// @notice Initiates a withdrawal request (called by NeoDAX).
    /// @dev Locks the withdrawal details and emits an event for Nitewatch to verify.
    /// @param user The recipient of the withdrawal.
    /// @param token The ERC20 token address, or address(0) for native ETH.
    /// @param amount The amount to withdraw.
    /// @param nonce A unique nonce to prevent replay attacks.
    /// @return withdrawalId The unique identifier for this withdrawal request.
    function startWithdraw(address user, address token, uint256 amount, uint256 nonce)
        external
        returns (bytes32 withdrawalId);

    /// @notice Finalizes a withdrawal request (called by Nitewatch).
    /// @dev Verifies the withdrawalId exists and executes the transfer.
    /// @param withdrawalId The unique identifier of the withdrawal to finalize.
    function finalizeWithdraw(bytes32 withdrawalId) external;

    /// @notice Rejects a withdrawal request (called by Nitewatch).
    /// @dev Verifies the withdrawalId exists and marks it as rejected.
    /// @param withdrawalId The unique identifier of the withdrawal to reject.
    function rejectWithdraw(bytes32 withdrawalId) external;
}
