// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {QuorumCustody} from "../src/QuorumCustody.sol";
import {ICustody} from "../src/interfaces/ICustody.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor() ERC20("Mock", "MCK") {}
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract QuorumCustodyTest is Test {
    QuorumCustody public custody;
    MockERC20 public token;

    address internal admin;
    address internal neodax;
    address internal user;
    address internal signer1;
    address internal signer2;
    address internal signer3;
    address internal signer4;
    address internal signer5;

    function setUp() public {
        admin = makeAddr("admin");
        neodax = makeAddr("neodax");
        user = makeAddr("user");

        signer1 = makeAddr("signer1");
        signer2 = makeAddr("signer2");
        signer3 = makeAddr("signer3");
        signer4 = makeAddr("signer4");
        signer5 = makeAddr("signer5");

        vm.startPrank(admin);
        custody = new QuorumCustody(admin, neodax, signer1);
        vm.stopPrank();

        token = new MockERC20();
    }

    function test_InitialState() public view {
        assertEq(custody.quorum(), 1);
        assertEq(custody.signers(0), signer1);
        assertTrue(custody.isSigner(signer1));
    }

    function test_FinalizeWithdraw_1_1_AsSigner() public {
        vm.deal(address(custody), 1 ether);
        
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_FinalizeWithdraw_2_2_Progressive() public {
        vm.prank(admin);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Signer 1 approves
        vm.prank(signer1);
        custody.finalizeWithdraw(id);
        
        // Not finalized yet
        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);
        
        // Signer 2 approves
        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // Finalized
        (,,,, finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_FinalizeWithdraw_3_5() public {
        vm.startPrank(admin);
        // 1/1 -> 1/2
        custody.addSigner(signer2, 1);
        // 1/2 -> 2/3
        custody.addSigner(signer3, 2);
        // 2/3 -> 2/4
        custody.addSigner(signer4, 2);
        // 2/4 -> 3/5
        custody.addSigner(signer5, 3);
        vm.stopPrank();

        assertEq(custody.quorum(), 3);
        assertEq(custody.signers(4), signer5);

        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id); // 1/3
        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer2);
        custody.finalizeWithdraw(id); // 2/3
        (,,,, finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer3);
        custody.finalizeWithdraw(id); // 3/3 -> execute
        (,,,, finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_Fail_DuplicateApproval() public {
        vm.prank(admin);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: signer already approved");
        custody.finalizeWithdraw(id);
    }
    
    function test_RejectWithdraw_Expired() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Advance time past expiry (3 days + 1 sec)
        vm.warp(block.timestamp + 3 days + 1);

        // Signer 1 rejects
        vm.prank(signer1);
        custody.rejectWithdraw(id);
        
        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized); // Finalized as rejected (or boolean logic in my contract sets finalized=true)
    }

    function test_Fail_Finalize_Expired() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Advance time past expiry
        vm.warp(block.timestamp + 3 days + 1);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: withdrawal expired");
        custody.finalizeWithdraw(id);
    }

    // =========================================================================
    // Constructor edge cases
    // =========================================================================

    function test_Fail_Constructor_ZeroAdmin() public {
        vm.expectRevert("QuorumCustody: invalid admin");
        new QuorumCustody(address(0), neodax, signer1);
    }

    function test_Fail_Constructor_ZeroSigner() public {
        vm.expectRevert("QuorumCustody: invalid signer");
        new QuorumCustody(admin, neodax, address(0));
    }

    // =========================================================================
    // Deposit edge cases
    // =========================================================================

    function test_DepositETH() public {
        vm.deal(user, 1 ether);
        vm.prank(user);
        custody.deposit{value: 1 ether}(address(0), 1 ether);
        assertEq(address(custody).balance, 1 ether);
    }

    function test_DepositERC20() public {
        token.mint(user, 100e18);
        vm.startPrank(user);
        token.approve(address(custody), 100e18);
        custody.deposit(address(token), 100e18);
        vm.stopPrank();
        assertEq(token.balanceOf(address(custody)), 100e18);
    }

    function test_DepositETH_EmitsEvent() public {
        vm.deal(user, 1 ether);
        vm.prank(user);
        vm.expectEmit(true, true, false, true);
        emit ICustody.Deposited(user, address(0), 1 ether);
        custody.deposit{value: 1 ether}(address(0), 1 ether);
    }

    function test_DepositERC20_EmitsEvent() public {
        token.mint(user, 50e18);
        vm.startPrank(user);
        token.approve(address(custody), 50e18);
        vm.expectEmit(true, true, false, true);
        emit ICustody.Deposited(user, address(token), 50e18);
        custody.deposit(address(token), 50e18);
        vm.stopPrank();
    }

    function test_Fail_DepositZeroAmount() public {
        vm.prank(user);
        vm.expectRevert("QuorumCustody: amount must be greater than 0");
        custody.deposit(address(0), 0);
    }

    function test_Fail_DepositETH_MsgValueMismatch() public {
        vm.deal(user, 2 ether);
        vm.prank(user);
        vm.expectRevert("QuorumCustody: msg.value mismatch");
        custody.deposit{value: 0.5 ether}(address(0), 1 ether);
    }

    function test_Fail_DepositERC20_NonZeroMsgValue() public {
        token.mint(user, 100e18);
        vm.deal(user, 1 ether);
        vm.startPrank(user);
        token.approve(address(custody), 100e18);
        vm.expectRevert("QuorumCustody: non-zero msg.value for ERC20");
        custody.deposit{value: 1 ether}(address(token), 100e18);
        vm.stopPrank();
    }

    // =========================================================================
    // addSigner edge cases
    // =========================================================================

    function test_Fail_AddSigner_NotAdmin() public {
        vm.prank(user);
        vm.expectRevert();
        custody.addSigner(signer2, 1);
    }

    function test_Fail_AddSigner_ZeroAddress() public {
        vm.prank(admin);
        vm.expectRevert("QuorumCustody: invalid signer");
        custody.addSigner(address(0), 1);
    }

    function test_Fail_AddSigner_Duplicate() public {
        vm.prank(admin);
        vm.expectRevert("QuorumCustody: already signer");
        custody.addSigner(signer1, 1);
    }

    function test_Fail_AddSigner_QuorumZero() public {
        vm.prank(admin);
        vm.expectRevert("QuorumCustody: invalid quorum");
        custody.addSigner(signer2, 0);
    }

    function test_Fail_AddSigner_QuorumTooHigh() public {
        // signers.length is 1, adding signer2 makes it 2, so quorum max is 2
        vm.prank(admin);
        vm.expectRevert("QuorumCustody: invalid quorum");
        custody.addSigner(signer2, 3);
    }

    function test_AddSigner_EmitsEvent() public {
        vm.prank(admin);
        vm.expectEmit(true, false, false, true);
        emit QuorumCustody.SignerAdded(signer2, 2);
        custody.addSigner(signer2, 2);
    }

    // =========================================================================
    // startWithdraw edge cases
    // =========================================================================

    function test_Fail_StartWithdraw_NotNeodax() public {
        vm.prank(user);
        vm.expectRevert();
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_Fail_StartWithdraw_ZeroAmount() public {
        vm.prank(neodax);
        vm.expectRevert("QuorumCustody: amount must be greater than 0");
        custody.startWithdraw(user, address(0), 0, 1);
    }

    function test_Fail_StartWithdraw_DuplicateNonce() public {
        vm.deal(address(custody), 2 ether);
        vm.startPrank(neodax);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.expectRevert("QuorumCustody: withdrawal already exists");
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.stopPrank();
    }

    function test_StartWithdraw_SameParamsDifferentNonce() public {
        vm.startPrank(neodax);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        vm.stopPrank();
        assertTrue(id1 != id2);
    }

    function test_StartWithdraw_EmitsEvent() public {
        vm.prank(neodax);
        vm.expectEmit(true, true, true, true);
        bytes32 expectedId = keccak256(abi.encode(block.chainid, address(custody), user, address(0), 1 ether, uint256(1)));
        emit ICustody.WithdrawStarted(expectedId, user, address(0), 1 ether, 1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    // =========================================================================
    // finalizeWithdraw edge cases
    // =========================================================================

    function test_Fail_FinalizeWithdraw_NotSigner() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(user);
        vm.expectRevert("QuorumCustody: only signer can finalize");
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: withdrawal not found");
        custody.finalizeWithdraw(bytes32(uint256(999)));
    }

    function test_Fail_FinalizeWithdraw_AlreadyFinalized() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: withdrawal already finalized");
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ERC20() public {
        token.mint(address(custody), 50e18);

        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        assertEq(token.balanceOf(user), 50e18);
        assertEq(token.balanceOf(address(custody)), 0);
    }

    function test_Fail_FinalizeWithdraw_InsufficientETH() public {
        // No ETH in custody
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: insufficient ETH liquidity");
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_InsufficientERC20() public {
        // No tokens in custody
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: insufficient ERC20 liquidity");
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ExactExpiryBoundary() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Warp to exactly the expiry boundary (should still succeed: <= check)
        vm.warp(block.timestamp + 3 days);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_FinalizeWithdraw_ClearsStorage() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (address storedUser, address storedToken, uint256 storedAmount, bool exists, bool finalized,,) = custody.withdrawals(id);
        assertTrue(exists);
        assertTrue(finalized);
        assertEq(storedUser, address(0));
        assertEq(storedToken, address(0));
        assertEq(storedAmount, 0);
    }

    function test_FinalizeWithdraw_EmitsApprovalEvent() public {
        vm.prank(admin);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, true, false, true);
        emit QuorumCustody.WithdrawalApproved(id, signer1, 1);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_EmitsFinalizedEvent() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit ICustody.WithdrawFinalized(id, true);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ETH_UserReceivesBalance() public {
        vm.deal(address(custody), 5 ether);
        uint256 balanceBefore = user.balance;

        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 2 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        assertEq(user.balance, balanceBefore + 2 ether);
        assertEq(address(custody).balance, 3 ether);
    }

    // =========================================================================
    // rejectWithdraw edge cases
    // =========================================================================

    function test_Fail_RejectWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: withdrawal not found");
        custody.rejectWithdraw(bytes32(uint256(999)));
    }

    function test_Fail_RejectWithdraw_AlreadyFinalized() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert("QuorumCustody: withdrawal already finalized");
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_UnauthorizedBeforeExpiry() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Random user cannot reject before expiry
        vm.prank(user);
        vm.expectRevert("QuorumCustody: unauthorized rejection");
        custody.rejectWithdraw(id);
    }

    function test_RejectWithdraw_AdminBeforeExpiry() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(admin);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_RejectWithdraw_SignerBeforeExpiry() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_RejectWithdraw_NeodaxAfterExpiry() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 3 days + 1);

        vm.prank(neodax);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_Fail_RejectWithdraw_RandomUserAfterExpiry() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 3 days + 1);

        vm.prank(user);
        vm.expectRevert("QuorumCustody: unauthorized rejection");
        custody.rejectWithdraw(id);
    }

    function test_RejectWithdraw_ExactExpiryBoundary_StillNotExpired() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Exactly at boundary: block.timestamp == createdAt + WITHDRAWAL_EXPIRY
        // Contract checks `>` for expiry, so this is NOT expired
        vm.warp(block.timestamp + 3 days);

        // Random user should NOT be able to reject (not expired yet)
        // But neodax doesn't have signer or admin role by default, so even neodax can't reject non-expired
        // Only signer/admin can reject before expiry
        vm.prank(user);
        vm.expectRevert("QuorumCustody: unauthorized rejection");
        custody.rejectWithdraw(id);
    }

    function test_RejectWithdraw_EmitsEvent() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit ICustody.WithdrawFinalized(id, false);
        custody.rejectWithdraw(id);
    }

    // =========================================================================
    // Quorum lifecycle: reject then re-create with same params + new nonce
    // =========================================================================

    function test_RejectThenRecreateWithdrawal() public {
        vm.deal(address(custody), 1 ether);

        vm.prank(neodax);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.rejectWithdraw(id1);

        // Same params but different nonce succeeds
        vm.prank(neodax);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        assertTrue(id1 != id2);

        vm.prank(signer1);
        custody.finalizeWithdraw(id2);

        assertEq(user.balance, 1 ether);
    }

    // =========================================================================
    // Partial approval then expiry (approvals don't carry over time)
    // =========================================================================

    function test_PartialApprovalThenExpiry() public {
        vm.prank(admin);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // First signer approves
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        // Time passes, expires before second approval
        vm.warp(block.timestamp + 3 days + 1);

        vm.prank(signer2);
        vm.expectRevert("QuorumCustody: withdrawal expired");
        custody.finalizeWithdraw(id);

        // Can still reject the expired withdrawal
        vm.prank(signer1);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // Multiple concurrent withdrawals
    // =========================================================================

    function test_MultipleConcurrentWithdrawals() public {
        vm.deal(address(custody), 3 ether);

        vm.startPrank(neodax);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        bytes32 id3 = custody.startWithdraw(user, address(0), 1 ether, 3);
        vm.stopPrank();

        // Finalize id1 and id3, reject id2
        vm.prank(signer1);
        custody.finalizeWithdraw(id1);
        vm.prank(signer1);
        custody.rejectWithdraw(id2);
        vm.prank(signer1);
        custody.finalizeWithdraw(id3);

        assertEq(user.balance, 2 ether);
        assertEq(address(custody).balance, 1 ether);
    }
}
