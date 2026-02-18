// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import {IWithdraw} from "./interfaces/IWithdraw.sol";
import {IDeposit} from "./interfaces/IDeposit.sol";

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

    event SignerAdded(address indexed signer, uint64 newQuorum);
    event SignerRemoved(address indexed signer, uint64 newQuorum);
    event QuorumChanged(uint64 oldQuorum, uint64 newQuorum);
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

    address[] public signers;
    mapping(address signer => bool isSigner) public isSigner;
    uint64 public quorum;

    constructor(address[] memory initialSigners, uint64 quorum_) EIP712("QuorumCustody", "1") {
        require(initialSigners.length != 0, EmptySignersArray());
        require(quorum_ != 0 && quorum_ <= initialSigners.length, InvalidQuorum());

        uint256 signersLength = initialSigners.length;
        for (uint256 i = 0; i < signersLength; i++) {
            _addSigner(initialSigners[i], quorum_);
        }

        quorum = quorum_;
    }

    modifier onlySigner() {
        require(isSigner[msg.sender], NotSigner());
        _;
    }

    function addSigners(address[] calldata newSigners, uint64 newQuorum, uint256 deadline, bytes[] calldata signatures)
        external
        onlySigner
    {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(newSigners.length != 0, EmptySignersArray());
        require(
            newQuorum != 0 && newQuorum >= quorum && newQuorum <= signers.length + newSigners.length,
            InvalidQuorum()
        );

        _verifySignatures(
            keccak256(
                abi.encode(ADD_SIGNERS_TYPEHASH, _hashAddressArray(newSigners), newQuorum, signerNonce, deadline)
            ),
            signatures
        );

        signerNonce++;
        for (uint256 i = 0; i < newSigners.length; i++) {
            _addSigner(newSigners[i], newQuorum);
        }
        if (quorum != newQuorum) {
            uint64 oldQuorum = quorum;
            quorum = newQuorum;
            emit QuorumChanged(oldQuorum, newQuorum);
        }
    }

    function removeSigners(
        address[] calldata signersToRemove,
        uint64 newQuorum,
        uint256 deadline,
        bytes[] calldata signatures
    ) external onlySigner {
        require(block.timestamp <= deadline, DeadlineExpired());
        require(signersToRemove.length != 0, EmptySignersArray());
        require(signersToRemove.length < signers.length, CannotRemoveLastSigner());
        uint256 remainingCount = signers.length - signersToRemove.length;
        uint64 minQuorum = quorum < remainingCount ? quorum : uint64(remainingCount);
        require(newQuorum != 0 && newQuorum >= minQuorum && newQuorum <= remainingCount, InvalidQuorum());

        _verifySignatures(
            keccak256(
                abi.encode(
                    REMOVE_SIGNERS_TYPEHASH, _hashAddressArray(signersToRemove), newQuorum, signerNonce, deadline
                )
            ),
            signatures
        );

        signerNonce++;
        uint256 signersLen = signers.length;
        for (uint256 i = 0; i < signersToRemove.length; i++) {
            address s = signersToRemove[i];
            require(isSigner[s], NotASigner());
            isSigner[s] = false;
            for (uint256 j = 0; j < signersLen; j++) {
                if (signers[j] == s) {
                    signers[j] = signers[signersLen - 1];
                    signers.pop();
                    signersLen--;
                    break;
                }
            }
            emit SignerRemoved(s, newQuorum);
        }
        if (quorum != newQuorum) {
            uint64 oldQuorum = quorum;
            quorum = newQuorum;
            emit QuorumChanged(oldQuorum, newQuorum);
        }
    }

    function getSignerCount() external view returns (uint256) {
        return signers.length;
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
            requiredQuorum: quorum,
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

    function _addSigner(address s, uint64 newQuorum) internal {
        require(s != address(0), InvalidSigner());
        require(!isSigner[s], AlreadySigner());
        signers.push(s);
        isSigner[s] = true;
        emit SignerAdded(s, newQuorum);
    }

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
            require(uint160(recovered) > uint160(lastSigner), SignaturesNotSorted());
            require(isSigner[recovered], InvalidSignature());
            require(recovered != msg.sender, SignerIsCaller());
            lastSigner = recovered;
            validApprovals++;
        }
        require(validApprovals >= quorum, InsufficientSignatures());
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
}
