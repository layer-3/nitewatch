// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICustody} from "./interfaces/ICustody.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract QuorumCustody is ICustody, ReentrancyGuard {
    using SafeERC20 for IERC20;

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
        require(initialSigner != address(0), "QuorumCustody: invalid signer");

        signers.push(initialSigner);
        isSigner[initialSigner] = true;
        quorum = 1;
        emit SignerAdded(initialSigner, 1);
    }

    modifier onlySigner() {
        require(isSigner[msg.sender], "QuorumCustody: caller is not a signer");
        _;
    }

    // =========================================================================
    // Signer Management (each call is a vote; executes when quorum is reached)
    // =========================================================================

    function addSigner(address signer, uint256 _quorum) external onlySigner {
        require(signer != address(0), "QuorumCustody: invalid signer");
        require(!isSigner[signer], "QuorumCustody: already signer");
        require(_quorum > 0 && _quorum <= signers.length + 1, "QuorumCustody: invalid quorum");

        bytes32 opHash = keccak256(abi.encode("addSigner", signer, _quorum, signerNonce));
        require(!operationApproved[opHash][msg.sender], "QuorumCustody: already approved");
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
        require(isSigner[signer], "QuorumCustody: not a signer");
        require(signers.length > 1, "QuorumCustody: cannot remove last signer");
        require(_quorum > 0 && _quorum <= signers.length - 1, "QuorumCustody: invalid quorum");

        bytes32 opHash = keccak256(abi.encode("removeSigner", signer, _quorum, signerNonce));
        require(!operationApproved[opHash][msg.sender], "QuorumCustody: already approved");
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
        require(amount > 0, "QuorumCustody: amount must be greater than 0");
        uint256 received = amount;
        if (token == address(0)) {
            require(msg.value == amount, "QuorumCustody: msg.value mismatch");
        } else {
            require(msg.value == 0, "QuorumCustody: non-zero msg.value for ERC20");
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
        require(user != address(0), "QuorumCustody: invalid user");
        require(amount > 0, "QuorumCustody: amount must be greater than 0");
        withdrawalId = keccak256(abi.encode(block.chainid, address(this), user, token, amount, nonce));

        require(!withdrawals[withdrawalId].exists, "QuorumCustody: withdrawal already exists");

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
        require(request.exists, "QuorumCustody: withdrawal not found");
        require(!request.finalized, "QuorumCustody: withdrawal already finalized");
        require(block.timestamp <= request.createdAt + WITHDRAWAL_EXPIRY, "QuorumCustody: withdrawal expired");
        require(!withdrawalApprovals[withdrawalId][msg.sender], "QuorumCustody: signer already approved");

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
        require(request.exists, "QuorumCustody: withdrawal not found");
        require(!request.finalized, "QuorumCustody: withdrawal already finalized");
        require(block.timestamp > request.createdAt + WITHDRAWAL_EXPIRY, "QuorumCustody: withdrawal not expired");

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
            require(address(this).balance >= amount, "QuorumCustody: insufficient ETH liquidity");
            (bool success,) = user.call{value: amount}("");
            require(success, "QuorumCustody: ETH transfer failed");
        } else {
            require(IERC20(token).balanceOf(address(this)) >= amount, "QuorumCustody: insufficient ERC20 liquidity");
            IERC20(token).safeTransfer(user, amount);
        }

        emit WithdrawFinalized(withdrawalId, true);
    }
}
