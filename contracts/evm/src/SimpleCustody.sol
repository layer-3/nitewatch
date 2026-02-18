// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {IWithdraw} from "./interfaces/IWithdraw.sol";
import {IDeposit} from "./interfaces/IDeposit.sol";

contract SimpleCustody is IWithdraw, IDeposit, AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant NEODAX_ROLE = keccak256("NEODAX_ROLE");
    bytes32 public constant NITEWATCH_ROLE = keccak256("NITEWATCH_ROLE");

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool exists;
        bool finalized;
    }

    mapping(bytes32 id => WithdrawalRequest request) public withdrawals;

    constructor(address admin, address neodax, address nitewatch) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(NEODAX_ROLE, neodax);
        _grantRole(NITEWATCH_ROLE, nitewatch);
    }

    function deposit(address token, uint256 amount) external payable override nonReentrant {
        if (amount == 0) revert ZeroAmount();
        uint256 received = amount;
        if (token == address(0)) {
            if (msg.value != amount) revert MsgValueMismatch();
        } else {
            if (msg.value != 0) revert NonZeroMsgValueForERC20();
            uint256 balanceBefore = IERC20(token).balanceOf(address(this));
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
            received = IERC20(token).balanceOf(address(this)) - balanceBefore;
        }
        emit Deposited(msg.sender, token, received);
    }

    function startWithdraw(address user, address token, uint256 amount, uint256 nonce)
        external
        override
        onlyRole(NEODAX_ROLE)
        nonReentrant
        returns (bytes32 withdrawalId)
    {
        if (amount == 0) revert ZeroAmount();
        withdrawalId = keccak256(abi.encode(block.chainid, address(this), user, token, amount, nonce));

        if (withdrawals[withdrawalId].exists) revert WithdrawalAlreadyExists();

        withdrawals[withdrawalId] =
            WithdrawalRequest({user: user, token: token, amount: amount, exists: true, finalized: false});

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlyRole(NITEWATCH_ROLE) nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();

        request.finalized = true;
        address user = request.user;
        address token = request.token;
        uint256 amount = request.amount;

        // Clear storage to refund gas, but keep 'exists' and 'finalized'
        request.user = address(0);
        request.token = address(0);
        request.amount = 0;

        if (token == address(0)) {
            if (address(this).balance < amount) revert InsufficientLiquidity();
            (bool success,) = user.call{value: amount}("");
            if (!success) revert ETHTransferFailed();
        } else {
            if (IERC20(token).balanceOf(address(this)) < amount) revert InsufficientLiquidity();
            IERC20(token).safeTransfer(user, amount);
        }

        emit WithdrawFinalized(withdrawalId, true);
    }

    function rejectWithdraw(bytes32 withdrawalId) external override onlyRole(NITEWATCH_ROLE) nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();

        request.finalized = true;

        emit WithdrawFinalized(withdrawalId, false);
    }
}
