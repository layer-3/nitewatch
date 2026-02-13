// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICustody} from "./interfaces/ICustody.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract QuorumCustody is ICustody, AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant NEODAX_ROLE = keccak256("NEODAX_ROLE");
    uint256 public constant WITHDRAWAL_EXPIRY = 3 days;

    struct WithdrawalRequest {
        address user;
        address token;
        uint256 amount;
        bool exists;
        bool finalized;
        uint256 approvalCount;
        uint256 createdAt;
    }

    mapping(bytes32 => WithdrawalRequest) public withdrawals;
    mapping(bytes32 => mapping(address => bool)) public approvals;
    
    address[] public signers;
    mapping(address => bool) public isSigner;
    uint256 public quorum;

    event SignerAdded(address indexed signer, uint256 newQuorum);
    event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals);

    constructor(address admin, address neodax, address initialSigner) {
        require(initialSigner != address(0), "QuorumCustody: invalid signer");
        require(admin != address(0), "QuorumCustody: invalid admin");
        
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(NEODAX_ROLE, neodax);

        signers.push(initialSigner);
        isSigner[initialSigner] = true;
        quorum = 1;
        emit SignerAdded(initialSigner, 1);
    }

    modifier onlySigner() {
        _checkSigner();
        _;
    }

    function _checkSigner() internal view {
        require(isSigner[msg.sender], "QuorumCustody: only signer can finalize");
    }

    function addSigner(address signer, uint256 _quorum) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(signer != address(0), "QuorumCustody: invalid signer");
        require(!isSigner[signer], "QuorumCustody: already signer");
        require(_quorum > 0 && _quorum <= signers.length + 1, "QuorumCustody: invalid quorum");

        signers.push(signer);
        isSigner[signer] = true;
        quorum = _quorum;
        emit SignerAdded(signer, _quorum);
    }

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

    function startWithdraw(address user, address token, uint256 amount, uint256 nonce)
        external
        override
        onlyRole(NEODAX_ROLE)
        nonReentrant
        returns (bytes32 withdrawalId)
    {
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
            createdAt: block.timestamp
        });

        emit WithdrawStarted(withdrawalId, user, token, amount, nonce);
    }

    function finalizeWithdraw(bytes32 withdrawalId) external override onlySigner nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        require(request.exists, "QuorumCustody: withdrawal not found");
        require(!request.finalized, "QuorumCustody: withdrawal already finalized");
        require(block.timestamp <= request.createdAt + WITHDRAWAL_EXPIRY, "QuorumCustody: withdrawal expired");
        require(!approvals[withdrawalId][msg.sender], "QuorumCustody: signer already approved");

        approvals[withdrawalId][msg.sender] = true;
        request.approvalCount += 1;
        
        emit WithdrawalApproved(withdrawalId, msg.sender, request.approvalCount);

        if (request.approvalCount >= quorum) {
            _executeWithdrawal(withdrawalId, request);
        }
    }

    function rejectWithdraw(bytes32 withdrawalId) external override nonReentrant {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        require(request.exists, "QuorumCustody: withdrawal not found");
        require(!request.finalized, "QuorumCustody: withdrawal already finalized");
        
        // Allow rejection if expired
        bool isExpired = block.timestamp > request.createdAt + WITHDRAWAL_EXPIRY;
        
        if (!isExpired) {
             // If not expired, only allow signers to reject (maybe explicit rejection logic could be added, 
             // but for now let's say only expiry allows "public" rejection or just stick to signers/admin)
             // For safety, let's restrict non-expired rejection to signers or admin.
             require(isSigner[msg.sender] || hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "QuorumCustody: unauthorized rejection");
        }
        
        // If expired, anyone can trigger rejection (or restrict to roles, but expiration usually implies invalidity)
        // Let's restrict to relevant parties to prevent griefing if that's a concern, but usually expiry cleanup is open or role based.
        // Given the prompt "if expiry is reach the withdraw is rejected", implies a condition.
        // Let's stick to safe roles: Admin, Signers, or Neodax (initiator).
        if (isExpired) {
             require(
                isSigner[msg.sender] || 
                hasRole(DEFAULT_ADMIN_ROLE, msg.sender) || 
                hasRole(NEODAX_ROLE, msg.sender), 
                "QuorumCustody: unauthorized rejection"
            );
        }

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
