// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICustody} from "./interfaces/ICustody.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract SimpleCustody is ICustody, AccessControl, ReentrancyGuard {
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

    mapping(bytes32 => WithdrawalRequest) public withdrawals;

    constructor(address admin, address neodax, address nitewatch) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(NEODAX_ROLE, neodax);
        _grantRole(NITEWATCH_ROLE, nitewatch);
    }

    function deposit(address token, uint256 amount) external payable override nonReentrant {
        require(amount > 0, "SimpleCustody: amount must be greater than 0");
        uint256 received = amount;
        if (token == address(0)) {
            require(msg.value == amount, "SimpleCustody: msg.value mismatch");
        } else {
            require(msg.value == 0, "SimpleCustody: non-zero msg.value for ERC20");
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
        require(amount > 0, "SimpleCustody: amount must be greater than 0");
        withdrawalId = keccak256(abi.encode(block.chainid, address(this), user, token, amount, nonce));

        require(!withdrawals[withdrawalId].exists, "SimpleCustody: withdrawal already exists");

        withdrawals[withdrawalId] =
            WithdrawalRequest({user: user, token: token, amount: amount, exists: true, finalized: false});

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlyRole(NITEWATCH_ROLE) nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        require(request.exists, "SimpleCustody: withdrawal not found");
        require(!request.finalized, "SimpleCustody: withdrawal already finalized");

        request.finalized = true;
        address user = request.user;
        address token = request.token;
        uint256 amount = request.amount;

        // Clear storage to refund gas, but keep 'exists' and 'finalized'
        request.user = address(0);
        request.token = address(0);
        request.amount = 0;

        if (token == address(0)) {
            require(address(this).balance >= amount, "SimpleCustody: insufficient ETH liquidity");
            (bool success,) = user.call{value: amount}("");
            require(success, "SimpleCustody: ETH transfer failed");
        } else {
            require(IERC20(token).balanceOf(address(this)) >= amount, "SimpleCustody: insufficient ERC20 liquidity");
            IERC20(token).safeTransfer(user, amount);
        }

        emit WithdrawFinalized(withdrawalId, true);
    }

    function rejectWithdraw(bytes32 withdrawalId) external override onlyRole(NITEWATCH_ROLE) nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        require(request.exists, "SimpleCustody: withdrawal not found");
        require(!request.finalized, "SimpleCustody: withdrawal already finalized");

        request.finalized = true;

        emit WithdrawFinalized(withdrawalId, false);
    }
}
