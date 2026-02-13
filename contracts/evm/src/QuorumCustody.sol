// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICustody} from "./interfaces/ICustody.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract QuorumCustody is ICustody, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ---- errors ----
    error InvalidSigner();
    error NotSigner();
    error AlreadySigner();
    error InvalidQuorum();
    error AlreadyApproved();
    error NotASigner();
    error CannotRemoveLastSigner();
    error InvalidUser();
    error WithdrawalExpired();
    error SignerAlreadyApproved();
    error WithdrawalNotExpired();

    uint256 public constant WITHDRAWAL_EXPIRY = 1 hours;

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool exists;
        bool finalized;
        uint256 approvalCount;
        uint256 requiredQuorum;
        uint256 createdAt;
    }

    mapping(bytes32 => WithdrawalRequest) public withdrawals;
    mapping(bytes32 => mapping(address => bool)) public withdrawalApprovals;

    // Generic quorum approval for signer operations
    mapping(bytes32 => uint256) public operationApprovals;
    mapping(bytes32 => mapping(address => bool)) public operationApproved;
    uint256 public signerNonce;

    address[] public signers;
    mapping(address => bool) public isSigner;
    uint256 public quorum;

    event SignerAdded(address indexed signer, uint256 newQuorum);
    event SignerRemoved(address indexed signer, uint256 newQuorum);
    event QuorumChanged(uint256 oldQuorum, uint256 newQuorum);
    event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals);

    constructor(address initialSigner) {
        if (initialSigner == address(0)) revert InvalidSigner();

        signers.push(initialSigner);
        isSigner[initialSigner] = true;
        quorum = 1;
        emit SignerAdded(initialSigner, 1);
    }

    modifier onlySigner() {
        if (!isSigner[msg.sender]) revert NotSigner();
        _;
    }

    // =========================================================================
    // Signer Management (each call is a vote; executes when quorum is reached)
    // =========================================================================

    function addSigner(address signer, uint256 _quorum) external onlySigner {
        if (signer == address(0)) revert InvalidSigner();
        if (isSigner[signer]) revert AlreadySigner();
        if (_quorum == 0 || _quorum > signers.length + 1) revert InvalidQuorum();

        bytes32 opHash = keccak256(abi.encode("addSigner", signer, _quorum, signerNonce));
        if (operationApproved[opHash][msg.sender]) revert AlreadyApproved();
        operationApproved[opHash][msg.sender] = true;
        operationApprovals[opHash]++;

        if (operationApprovals[opHash] >= quorum) {
            signerNonce++;
            signers.push(signer);
            isSigner[signer] = true;
            uint256 oldQuorum = quorum;
            quorum = _quorum;
            emit SignerAdded(signer, _quorum);
            if (_quorum != oldQuorum) emit QuorumChanged(oldQuorum, _quorum);
        }
    }

    function removeSigner(address signer, uint256 _quorum) external onlySigner {
        if (!isSigner[signer]) revert NotASigner();
        if (signers.length <= 1) revert CannotRemoveLastSigner();
        if (_quorum == 0 || _quorum > signers.length - 1) revert InvalidQuorum();

        bytes32 opHash = keccak256(abi.encode("removeSigner", signer, _quorum, signerNonce));
        if (operationApproved[opHash][msg.sender]) revert AlreadyApproved();
        operationApproved[opHash][msg.sender] = true;
        operationApprovals[opHash]++;

        if (operationApprovals[opHash] >= quorum) {
            signerNonce++;
            isSigner[signer] = false;
            uint256 len = signers.length;
            for (uint256 i = 0; i < len; i++) {
                if (signers[i] == signer) {
                    signers[i] = signers[len - 1];
                    signers.pop();
                    break;
                }
            }
            uint256 oldQuorum = quorum;
            quorum = _quorum;
            emit SignerRemoved(signer, _quorum);
            if (_quorum != oldQuorum) emit QuorumChanged(oldQuorum, _quorum);
        }
    }

    function getSignerCount() external view returns (uint256) {
        return signers.length;
    }

    // =========================================================================
    // Deposit
    // =========================================================================

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

    // =========================================================================
    // Withdrawal
    // =========================================================================

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
            approvalCount: 0,
            requiredQuorum: quorum,
            createdAt: block.timestamp
        });

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();
        if (block.timestamp > request.createdAt + WITHDRAWAL_EXPIRY) revert WithdrawalExpired();
        if (withdrawalApprovals[withdrawalId][msg.sender]) revert SignerAlreadyApproved();

        withdrawalApprovals[withdrawalId][msg.sender] = true;
        request.approvalCount += 1;

        emit WithdrawalApproved(withdrawalId, msg.sender, request.approvalCount);

        if (request.approvalCount >= request.requiredQuorum) {
            _executeWithdrawal(withdrawalId, request);
        }
    }

    /// @notice Rejects an expired withdrawal. Non-expired withdrawals cannot be rejected;
    ///         they simply expire if they don't reach quorum within WITHDRAWAL_EXPIRY.
    function rejectWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        if (!request.exists) revert WithdrawalNotFound();
        if (request.finalized) revert WithdrawalAlreadyFinalized();
        if (block.timestamp <= request.createdAt + WITHDRAWAL_EXPIRY) revert WithdrawalNotExpired();

        request.finalized = true;
        emit WithdrawFinalized(withdrawalId, false);
    }

    function _executeWithdrawal(bytes32 withdrawalId, WithdrawalRequest storage request) internal {
        request.finalized = true;
        address user = request.user;
        address token = request.token;
        uint256 amount = request.amount;

        // Clear storage
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
}
