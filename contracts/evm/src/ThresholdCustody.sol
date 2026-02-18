// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {MultiSignerERC7913} from "@openzeppelin/contracts/utils/cryptography/signers/MultiSignerERC7913.sol";

import {IWithdraw} from "./interfaces/IWithdraw.sol";
import {IDeposit} from "./interfaces/IDeposit.sol";

contract ThresholdCustody is IWithdraw, IDeposit, ReentrancyGuard, EIP712, MultiSignerERC7913 {
    using SafeERC20 for IERC20;

    error NotSigner();
    error InvalidQuorum();
    error InvalidUser();
    error WithdrawalExpired();
    error SignerAlreadyApproved();
    error WithdrawalNotExpired();
    error InvalidSignature();
    error EmptySignersArray();
    error DeadlineExpired();

    event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals);

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool finalized;
        uint64 createdAt;
        uint64 requiredQuorum;
    }

    bytes32 public constant ADD_SIGNERS_TYPEHASH =
        keccak256("AddSigners(address[] newSigners,uint256 newQuorum,uint256 nonce,uint256 deadline)");
    bytes32 public constant REMOVE_SIGNERS_TYPEHASH =
        keccak256("RemoveSigners(address[] signersToRemove,uint256 newQuorum,uint256 nonce,uint256 deadline)");

    uint256 public constant OPERATION_EXPIRY = 1 hours;

    mapping(bytes32 withdrawalId => WithdrawalRequest request) public withdrawals;
    mapping(bytes32 withdrawalId => mapping(address signer => bool hasApproved)) public withdrawalApprovals;
    uint256 public signerNonce;

    constructor(address[] memory initialSigners, uint64 quorum_)
        EIP712("ThresholdCustody", "1")
        MultiSignerERC7913(_toAddressBytesArray(initialSigners), quorum_)
    {
        require(initialSigners.length != 0, EmptySignersArray());
        require(quorum_ != 0 && quorum_ <= initialSigners.length, InvalidQuorum());

    }

    modifier onlySigner() {
        require(isSigner(msg.sender), NotSigner());
        _;
    }

    function isSigner(address signer) public view returns (bool) {
        return isSigner(_toBytes(signer));
    }

    function addSigners(
        address[] calldata newSigners,
        uint64 newThreshold,
        uint256 deadline,
        bytes calldata signatures
    ) external onlySigner {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(newSigners.length != 0, EmptySignersArray());

        bytes32 structHash = keccak256(
            abi.encode(ADD_SIGNERS_TYPEHASH, _hashAddressArray(newSigners), newThreshold, signerNonce, deadline)
        );
        bytes32 digest = _hashTypedDataV4(structHash);

        require(_rawSignatureValidation(digest, signatures), InvalidSignature());

        signerNonce++;

        _addSigners(_toAddressBytesArray(newSigners));
        _setThreshold(newThreshold);
    }

    function removeSigners(
        address[] calldata signersToRemove,
        uint64 newThreshold,
        uint256 deadline,
        bytes calldata signatures
    ) external onlySigner {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(signersToRemove.length != 0, EmptySignersArray());

        bytes32 structHash = keccak256(
            abi.encode(REMOVE_SIGNERS_TYPEHASH, _hashAddressArray(signersToRemove), newThreshold, signerNonce, deadline)
        );
        bytes32 digest = _hashTypedDataV4(structHash);

        require(_rawSignatureValidation(digest, signatures), InvalidSignature());

        signerNonce++;

        _removeSigners(_toAddressBytesArray(signersToRemove));
        _setThreshold(newThreshold);
    }

    function deposit(address token, uint256 amount) external payable override nonReentrant {
        require(amount != 0, ZeroAmount());

        if (token == address(0)) {
            require(msg.value == amount, MsgValueMismatch());
        } else {
            require(msg.value == 0, NonZeroMsgValueForERC20());
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
        require(amount != 0, ZeroAmount());

        bytes32 withdrawalId = _getWithdrawalId(user, token, amount, nonce);
        require(withdrawals[withdrawalId].createdAt == 0, WithdrawalAlreadyExists());

        withdrawals[withdrawalId] = WithdrawalRequest({
            user: user,
            token: token,
            amount: amount,
            finalized: false,
            requiredQuorum: threshold(),
            createdAt: uint64(block.timestamp)
        });

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
        return withdrawalId;
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        address signer = msg.sender;

        require(request.createdAt != 0, WithdrawalNotFound());
        require(!request.finalized, WithdrawalAlreadyFinalized());
        require(block.timestamp <= request.createdAt + OPERATION_EXPIRY, WithdrawalExpired());
        require(!withdrawalApprovals[withdrawalId][signer], SignerAlreadyApproved());

        withdrawalApprovals[withdrawalId][signer] = true;
        uint256 validApprovals = _countValidApprovals(withdrawalId);

        emit WithdrawalApproved(withdrawalId, signer, validApprovals);

        if (validApprovals >= request.requiredQuorum) {
            _executeWithdrawal(request);
            emit WithdrawFinalized(withdrawalId, true);
        }
    }

    function rejectWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];

        require(request.createdAt != 0, WithdrawalNotFound());
        require(!request.finalized, WithdrawalAlreadyFinalized());
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
            require(address(this).balance >= amount, InsufficientLiquidity());
            (bool success,) = user.call{value: amount}("");
            require(success, ETHTransferFailed());
        } else {
            require(IERC20(token).balanceOf(address(this)) >= amount, InsufficientLiquidity());
            IERC20(token).safeTransfer(user, amount);
        }

        request.user = address(0);
        request.token = address(0);
        request.amount = 0;
    }

    function _countValidApprovals(bytes32 withdrawalId) internal view returns (uint256 count) {
        bytes[] memory allSigners = getSigners(0, type(uint64).max);

        for (uint256 i = 0; i < allSigners.length; i++) {
            address s = _bytesToAddress(allSigners[i]);
            if (withdrawalApprovals[withdrawalId][s]) count++;
        }
    }

    function _hashAddressArray(address[] calldata arr) internal pure returns (bytes32) {
        bytes32[] memory encoded = new bytes32[](arr.length);
        for (uint256 i = 0; i < arr.length; i++) {
            encoded[i] = bytes32(uint256(uint160(arr[i])));
        }
        return keccak256(abi.encodePacked(encoded));
    }

    function _getWithdrawalId(address user, address token, uint256 amount, uint256 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(block.chainid, address(this), user, token, amount, nonce));
    }

    // Helpers for conversion
    function _toBytes(address a) internal pure returns (bytes memory) {
        return abi.encodePacked(a);
    }

    function _toAddressBytesArray(address[] memory addrs) internal pure returns (bytes[] memory) {
        bytes[] memory b = new bytes[](addrs.length);
        for(uint i=0; i<addrs.length; i++) {
            b[i] = _toBytes(addrs[i]);
        }
        return b;
    }

    function _bytesToAddress(bytes memory b) internal pure returns (address) {
        return address(bytes20(b));
    }
}
