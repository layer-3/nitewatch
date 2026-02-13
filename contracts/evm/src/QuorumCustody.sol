// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IWithdraw} from "./interfaces/IWithdraw.sol";
import {IDeposit} from "./interfaces/IDeposit.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract QuorumCustody is IWithdraw, IDeposit, ReentrancyGuard, EIP712 {
    using SafeERC20 for IERC20;

    error InvalidSigner();
    error NotSigner();
    error AlreadySigner();
    error InvalidQuorum();
    error NotASigner();
    error CannotRemoveLastSigner();
    error InvalidUser();
    error WithdrawalExpired();
    error SignerAlreadyApproved();
    error WithdrawalNotExpired();
    error InvalidSignature();
    error SignaturesNotSorted();
    error InsufficientSignatures();
    error EmptySignersArray();
    error SignerIsCaller();
    error DeadlineExpired();

    bytes32 public constant ADD_SIGNERS_TYPEHASH =
        keccak256("AddSigners(address[] newSigners,uint256 newQuorum,uint256 nonce,uint256 deadline)");
    bytes32 public constant REMOVE_SIGNERS_TYPEHASH =
        keccak256("RemoveSigners(address[] signersToRemove,uint256 newQuorum,uint256 nonce,uint256 deadline)");

    uint256 public constant OPERATION_EXPIRY = 1 hours;

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool exists;
        bool finalized;
        uint256 requiredQuorum;
        uint256 createdAt;
    }

    mapping(bytes32 => WithdrawalRequest) public withdrawals;
    mapping(bytes32 => mapping(address => bool)) public withdrawalApprovals;
    uint256 public signerNonce;

    address[] public signers;
    mapping(address => bool) public isSigner;
    uint256 public quorum;

    event SignerAdded(address indexed signer, uint256 newQuorum);
    event SignerRemoved(address indexed signer, uint256 newQuorum);
    event QuorumChanged(uint256 oldQuorum, uint256 newQuorum);
    event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals);

    constructor(address[] memory initialSigners, uint256 _quorum) EIP712("QuorumCustody", "1") {
        if (initialSigners.length == 0) revert EmptySignersArray();
        if (_quorum == 0 || _quorum > initialSigners.length) revert InvalidQuorum();
        for (uint256 i = 0; i < initialSigners.length; i++) {
            _addSigner(initialSigners[i], _quorum);
        }
        quorum = _quorum;
    }

    modifier onlySigner() {
        if (!isSigner[msg.sender]) revert NotSigner();
        _;
    }

    function addSigners(address[] calldata newSigners, uint256 newQuorum, uint256 deadline, bytes[] calldata signatures)
        external
        onlySigner
    {
        if (block.timestamp > deadline) revert DeadlineExpired();
        if (newSigners.length == 0) revert EmptySignersArray();
        if (newQuorum == 0 || newQuorum < quorum || newQuorum > signers.length + newSigners.length) {
            revert InvalidQuorum();
        }

        _verifySignatures(
            keccak256(
                abi.encode(ADD_SIGNERS_TYPEHASH, _hashAddressArray(newSigners), newQuorum, signerNonce, deadline)
            ),
            signatures
        );

        signerNonce++;
        uint256 oldQuorum = quorum;
        quorum = newQuorum;
        for (uint256 i = 0; i < newSigners.length; i++) {
            _addSigner(newSigners[i], newQuorum);
        }
        if (newQuorum != oldQuorum) emit QuorumChanged(oldQuorum, newQuorum);
    }

    function removeSigners(
        address[] calldata signersToRemove,
        uint256 newQuorum,
        uint256 deadline,
        bytes[] calldata signatures
    ) external onlySigner {
        if (block.timestamp > deadline) revert DeadlineExpired();
        if (signersToRemove.length == 0) revert EmptySignersArray();
        if (signersToRemove.length >= signers.length) revert CannotRemoveLastSigner();
        uint256 remainingCount = signers.length - signersToRemove.length;
        uint256 minQuorum = quorum < remainingCount ? quorum : remainingCount;
        if (newQuorum == 0 || newQuorum < minQuorum || newQuorum > remainingCount) revert InvalidQuorum();

        _verifySignatures(
            keccak256(
                abi.encode(
                    REMOVE_SIGNERS_TYPEHASH, _hashAddressArray(signersToRemove), newQuorum, signerNonce, deadline
                )
            ),
            signatures
        );

        signerNonce++;
        uint256 oldQuorum = quorum;
        quorum = newQuorum;
        for (uint256 i = 0; i < signersToRemove.length; i++) {
            address s = signersToRemove[i];
            if (!isSigner[s]) revert NotASigner();
            isSigner[s] = false;
            uint256 len = signers.length;
            for (uint256 j = 0; j < len; j++) {
                if (signers[j] == s) {
                    signers[j] = signers[len - 1];
                    signers.pop();
                    break;
                }
            }
            emit SignerRemoved(s, newQuorum);
        }
        if (newQuorum != oldQuorum) emit QuorumChanged(oldQuorum, newQuorum);
    }

    function getSignerCount() external view returns (uint256) {
        return signers.length;
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
        onlySigner
        nonReentrant
        returns (bytes32 withdrawalId)
    {
        if (user == address(0)) revert InvalidUser();
        if (amount == 0) revert ZeroAmount();
        withdrawalId = keccak256(abi.encode(block.chainid, address(this), user, token, amount, nonce));
        if (withdrawals[withdrawalId].exists) revert WithdrawalAlreadyExists();

        withdrawals[withdrawalId] = WithdrawalRequest({
            user: user,
            token: token,
            amount: amount,
            exists: true,
            finalized: false,
            requiredQuorum: quorum,
            createdAt: block.timestamp
        });
        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();
        if (block.timestamp > request.createdAt + OPERATION_EXPIRY) revert WithdrawalExpired();
        if (withdrawalApprovals[withdrawalId][msg.sender]) revert SignerAlreadyApproved();

        withdrawalApprovals[withdrawalId][msg.sender] = true;
        uint256 validApprovals = _countValidApprovals(withdrawalId);
        emit WithdrawalApproved(withdrawalId, msg.sender, validApprovals);

        if (validApprovals >= request.requiredQuorum) {
            _executeWithdrawal(withdrawalId, request);
        }
    }

    function rejectWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();
        if (block.timestamp <= request.createdAt + OPERATION_EXPIRY) revert WithdrawalNotExpired();

        request.finalized = true;
        emit WithdrawFinalized(withdrawalId, false);
    }

    // --- Internal ---

    function _addSigner(address s, uint256 newQuorum) internal {
        if (s == address(0)) revert InvalidSigner();
        if (isSigner[s]) revert AlreadySigner();
        signers.push(s);
        isSigner[s] = true;
        emit SignerAdded(s, newQuorum);
    }

    function _executeWithdrawal(bytes32 withdrawalId, WithdrawalRequest storage request) internal {
        request.finalized = true;
        address user = request.user;
        address token = request.token;
        uint256 amount = request.amount;
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

    function _countValidApprovals(bytes32 withdrawalId) internal view returns (uint256 count) {
        uint256 len = signers.length;
        for (uint256 i = 0; i < len; i++) {
            if (withdrawalApprovals[withdrawalId][signers[i]]) count++;
        }
    }

    function _verifySignatures(bytes32 structHash, bytes[] calldata signatures) internal view {
        bytes32 digest = _hashTypedDataV4(structHash);
        uint256 validApprovals = 1;
        address lastSigner = address(0);
        for (uint256 i = 0; i < signatures.length; i++) {
            address recovered = ECDSA.recover(digest, signatures[i]);
            if (uint160(recovered) <= uint160(lastSigner)) revert SignaturesNotSorted();
            if (!isSigner[recovered]) revert InvalidSignature();
            if (recovered == msg.sender) revert SignerIsCaller();
            lastSigner = recovered;
            validApprovals++;
        }
        if (validApprovals < quorum) revert InsufficientSignatures();
    }

    function _hashAddressArray(address[] calldata arr) internal pure returns (bytes32) {
        bytes32[] memory encoded = new bytes32[](arr.length);
        for (uint256 i = 0; i < arr.length; i++) {
            encoded[i] = bytes32(uint256(uint160(arr[i])));
        }
        return keccak256(abi.encodePacked(encoded));
    }
}
