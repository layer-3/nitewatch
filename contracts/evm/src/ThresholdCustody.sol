// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {MultiSignerERC7913} from "@openzeppelin/contracts/utils/cryptography/signers/MultiSignerERC7913.sol";

import {IWithdraw} from "./interfaces/IWithdraw.sol";
import {IDeposit} from "./interfaces/IDeposit.sol";
import {Utils} from "./Utils.sol";

bytes32 constant SET_THRESHOLD_TYPEHASH =
    keccak256("SetThreshold(uint256 newThreshold,uint256 nonce,uint256 deadline)");
bytes32 constant ADD_SIGNERS_TYPEHASH =
    keccak256("AddSigners(address[] newSigners,uint256 newThreshold,uint256 nonce,uint256 deadline)");
bytes32 constant REMOVE_SIGNERS_TYPEHASH =
    keccak256("RemoveSigners(address[] signersToRemove,uint256 newThreshold,uint256 nonce,uint256 deadline)");

uint256 constant OPERATION_EXPIRY = 1 hours;

contract ThresholdCustody is IWithdraw, IDeposit, ReentrancyGuard, EIP712, MultiSignerERC7913 {
    using SafeERC20 for IERC20;
    using {Utils.hashArrayed, Utils.toAddressBytesArray} for address[];
    using {Utils.toBytes} for address;
    using {Utils.toAddress} for bytes;

    error EmptySignersArray();
    error DeadlineExpired();
    error InvalidSignature();
    error InvalidThreshold();
    error NotSigner();
    error InvalidUser();
    error SignerAlreadyApproved();
    error WithdrawalNotExpired();

    event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals);

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool finalized;
        uint64 createdAt;
        uint64 requiredThreshold;
    }

    mapping(bytes32 withdrawalId => WithdrawalRequest request) public withdrawals;
    mapping(bytes32 withdrawalId => mapping(address signer => bool hasApproved)) public withdrawalApprovals;
    uint256 public signerNonce;

    constructor(address[] memory initialSigners, uint64 threshold)
        EIP712("ThresholdCustody", "1")
        MultiSignerERC7913(initialSigners.toAddressBytesArray(), threshold)
    {
        require(initialSigners.length != 0, EmptySignersArray());
    }

    modifier onlySigner() {
        require(isSigner(msg.sender), NotSigner());
        _;
    }

    function isSigner(address signer) public view returns (bool) {
        return isSigner(signer.toBytes());
    }

    function setThreshold(uint64 newThreshold, uint256 deadline, bytes calldata signatures) external {
        require(block.timestamp <= deadline, DeadlineExpired());

        bytes32 structHash = keccak256(abi.encode(SET_THRESHOLD_TYPEHASH, newThreshold, signerNonce, deadline));
        bytes32 digest = _hashTypedDataV4(structHash);

        require(_rawSignatureValidation(digest, signatures), InvalidSignature());

        signerNonce++;

        _setThreshold(newThreshold);
    }

    function addSigners(address[] calldata newSigners, uint64 newThreshold, uint256 deadline, bytes calldata signatures)
        external
    {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(newSigners.length != 0, EmptySignersArray());

        bytes32 structHash =
            keccak256(abi.encode(ADD_SIGNERS_TYPEHASH, newSigners.hashArrayed(), newThreshold, signerNonce, deadline));
        bytes32 digest = _hashTypedDataV4(structHash);

        require(_rawSignatureValidation(digest, signatures), InvalidSignature());

        signerNonce++;

        _addSigners(newSigners.toAddressBytesArray());
        _setThreshold(newThreshold);
    }

    function removeSigners(
        address[] calldata signersToRemove,
        uint64 newThreshold,
        uint256 deadline,
        bytes calldata signatures
    ) external {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(signersToRemove.length != 0, EmptySignersArray());

        bytes32 structHash = keccak256(
            abi.encode(REMOVE_SIGNERS_TYPEHASH, signersToRemove.hashArrayed(), newThreshold, signerNonce, deadline)
        );
        bytes32 digest = _hashTypedDataV4(structHash);

        require(_rawSignatureValidation(digest, signatures), InvalidSignature());

        signerNonce++;

        _removeSigners(signersToRemove.toAddressBytesArray());
        _setThreshold(newThreshold);
    }

    function deposit(address token, uint256 amount) external payable override nonReentrant {
        require(amount != 0, IDeposit.ZeroAmount());

        if (token == address(0)) {
            require(msg.value == amount, IDeposit.InvalidMsgValue());
        } else {
            require(msg.value == 0, IDeposit.InvalidMsgValue());
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        }

        emit Deposited(msg.sender, token, amount);
    }

    function startWithdraw(address user, address token, uint256 amount, uint256 nonce)
        external
        override
        onlySigner
        nonReentrant
        returns (bytes32)
    {
        require(user != address(0), InvalidUser());
        require(amount != 0, IDeposit.ZeroAmount());

        bytes32 withdrawalId = Utils.getWithdrawalId(address(this), user, token, amount, nonce);
        require(withdrawals[withdrawalId].createdAt == 0, IWithdraw.WithdrawalAlreadyExists());

        withdrawals[withdrawalId] = WithdrawalRequest({
            user: user,
            token: token,
            amount: amount,
            finalized: false,
            requiredThreshold: threshold(),
            createdAt: uint64(block.timestamp)
        });

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
        return withdrawalId;
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        address signer = msg.sender;

        require(request.createdAt != 0, IWithdraw.WithdrawalNotFound());
        require(!request.finalized, IWithdraw.WithdrawalAlreadyFinalized());
        require(block.timestamp <= request.createdAt + OPERATION_EXPIRY, DeadlineExpired());
        require(!withdrawalApprovals[withdrawalId][signer], SignerAlreadyApproved());

        withdrawalApprovals[withdrawalId][signer] = true;
        uint256 validApprovals = _countValidApprovals(withdrawalId);

        emit WithdrawalApproved(withdrawalId, signer, validApprovals);

        if (validApprovals >= request.requiredThreshold) {
            _executeWithdrawal(request);
            emit WithdrawFinalized(withdrawalId, true);
        }
    }

    function rejectWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];

        require(request.createdAt != 0, IWithdraw.WithdrawalNotFound());
        require(!request.finalized, IWithdraw.WithdrawalAlreadyFinalized());
        require(block.timestamp > request.createdAt + OPERATION_EXPIRY, WithdrawalNotExpired());

        request.finalized = true;
        emit WithdrawFinalized(withdrawalId, false);
    }

    // --- Internal ---

    function _executeWithdrawal(WithdrawalRequest storage request) internal {
        address user = request.user;
        address token = request.token;
        uint256 amount = request.amount;

        request.finalized = true;

        if (token == address(0)) {
            require(address(this).balance >= amount, IWithdraw.InsufficientLiquidity());
            (bool success,) = user.call{value: amount}("");
            require(success, IWithdraw.ETHTransferFailed());
        } else {
            require(IERC20(token).balanceOf(address(this)) >= amount, IWithdraw.InsufficientLiquidity());
            IERC20(token).safeTransfer(user, amount);
        }

        request.user = address(0);
        request.token = address(0);
        request.amount = 0;
    }

    function _countValidApprovals(bytes32 withdrawalId) internal view returns (uint256 count) {
        bytes[] memory allSigners = getSigners(0, type(uint64).max);

        for (uint256 i = 0; i < allSigners.length; i++) {
            address s = allSigners[i].toAddress();
            if (withdrawalApprovals[withdrawalId][s]) count++;
        }
    }
}
