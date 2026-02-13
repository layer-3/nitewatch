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

    address internal user;
    address internal signer1;
    address internal signer2;
    address internal signer3;
    address internal signer4;
    address internal signer5;

    function setUp() public {
        user = makeAddr("user");
        signer1 = makeAddr("signer1");
        signer2 = makeAddr("signer2");
        signer3 = makeAddr("signer3");
        signer4 = makeAddr("signer4");
        signer5 = makeAddr("signer5");

        custody = new QuorumCustody(signer1);
        token = new MockERC20();
    }

    // Helper: set up 5 signers with quorum=3
    function _setup3of5() internal {
        // quorum=1 → auto-executes
        vm.prank(signer1);
        custody.addSigner(signer2, 1);
        vm.prank(signer1);
        custody.addSigner(signer3, 2); // quorum becomes 2

        // quorum=2 → needs two voters
        vm.prank(signer1);
        custody.addSigner(signer4, 2);
        vm.prank(signer2);
        custody.addSigner(signer4, 2); // executes

        vm.prank(signer1);
        custody.addSigner(signer5, 3);
        vm.prank(signer2);
        custody.addSigner(signer5, 3); // executes, quorum becomes 3
    }

    // =========================================================================
    // Initial state
    // =========================================================================

    function test_InitialState() public view {
        assertEq(custody.quorum(), 1);
        assertEq(custody.signers(0), signer1);
        assertTrue(custody.isSigner(signer1));
        assertEq(custody.getSignerCount(), 1);
    }

    function test_Fail_Constructor_ZeroSigner() public {
        vm.expectRevert(QuorumCustody.InvalidSigner.selector);
        new QuorumCustody(address(0));
    }

    // =========================================================================
    // addSigner
    // =========================================================================

    function test_AddSigner_AutoExecutesAtQuorum1() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2);

        assertTrue(custody.isSigner(signer2));
        assertEq(custody.quorum(), 2);
        assertEq(custody.getSignerCount(), 2);
    }

    function test_AddSigner_PendingAtHigherQuorum() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2); // auto-executes, quorum=2

        vm.prank(signer1);
        custody.addSigner(signer3, 2); // 1 of 2, pending

        assertFalse(custody.isSigner(signer3));
        assertEq(custody.getSignerCount(), 2);
    }

    function test_AddSigner_ExecutesAtQuorum() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2); // auto-executes, quorum=2

        vm.prank(signer1);
        custody.addSigner(signer3, 2); // 1 of 2

        vm.prank(signer2);
        custody.addSigner(signer3, 2); // 2 of 2, executes

        assertTrue(custody.isSigner(signer3));
        assertEq(custody.getSignerCount(), 3);
    }

    function test_AddSigner_EmitsSignerAdded() public {
        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit QuorumCustody.SignerAdded(signer2, 2);
        custody.addSigner(signer2, 2);
    }

    function test_AddSigner_EmitsQuorumChanged() public {
        vm.prank(signer1);
        vm.expectEmit(false, false, false, true);
        emit QuorumCustody.QuorumChanged(1, 2);
        custody.addSigner(signer2, 2);
    }

    function test_AddSigner_NoQuorumChangedWhenSame() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1); // quorum stays 1

        // No QuorumChanged event — just verify quorum didn't change
        assertEq(custody.quorum(), 1);
    }

    function test_Fail_AddSigner_NotSigner() public {
        vm.prank(user);
        vm.expectRevert(QuorumCustody.NotSigner.selector);
        custody.addSigner(signer2, 1);
    }

    function test_Fail_AddSigner_ZeroAddress() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.InvalidSigner.selector);
        custody.addSigner(address(0), 1);
    }

    function test_Fail_AddSigner_Duplicate() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.AlreadySigner.selector);
        custody.addSigner(signer1, 1);
    }

    function test_Fail_AddSigner_QuorumZero() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.InvalidQuorum.selector);
        custody.addSigner(signer2, 0);
    }

    function test_Fail_AddSigner_QuorumTooHigh() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.InvalidQuorum.selector);
        custody.addSigner(signer2, 3); // max is signers.length+1 = 2
    }

    function test_Fail_AddSigner_AlreadyApproved() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2); // auto, quorum=2

        vm.prank(signer1);
        custody.addSigner(signer3, 2); // 1 of 2

        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.AlreadyApproved.selector);
        custody.addSigner(signer3, 2); // duplicate vote
    }

    function test_AddSigner_NonceInvalidatesOldApprovals() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2); // auto, quorum=2

        // signer1 votes for signer3 under nonce=1
        vm.prank(signer1);
        custody.addSigner(signer3, 2);

        // A different addSigner executes, incrementing nonce to 2
        vm.prank(signer1);
        custody.addSigner(signer4, 2);
        vm.prank(signer2);
        custody.addSigner(signer4, 2); // executes, nonce becomes 2

        // signer1's old vote for signer3 used nonce=1, now nonce=2 → fresh start
        // signer1 can vote again with the new nonce
        vm.prank(signer1);
        custody.addSigner(signer3, 2); // new opHash, 1 of 2

        assertFalse(custody.isSigner(signer3)); // still pending
    }

    // =========================================================================
    // removeSigner
    // =========================================================================

    function test_RemoveSigner_AutoExecutesAtQuorum1() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1); // quorum stays 1

        vm.prank(signer1);
        custody.removeSigner(signer2, 1);

        assertFalse(custody.isSigner(signer2));
        assertEq(custody.getSignerCount(), 1);
    }

    function test_RemoveSigner_RequiresQuorum() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2); // quorum=2

        // Add signer3 so we can remove signer2
        vm.prank(signer1);
        custody.addSigner(signer3, 2);
        vm.prank(signer2);
        custody.addSigner(signer3, 2); // executes

        // Remove signer2 (needs 2 votes)
        vm.prank(signer1);
        custody.removeSigner(signer2, 2);
        assertTrue(custody.isSigner(signer2)); // not yet

        vm.prank(signer3);
        custody.removeSigner(signer2, 2); // executes
        assertFalse(custody.isSigner(signer2));
        assertEq(custody.getSignerCount(), 2);
    }

    function test_RemoveSigner_EmitsEvent() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit QuorumCustody.SignerRemoved(signer2, 1);
        custody.removeSigner(signer2, 1);
    }

    function test_Fail_RemoveSigner_NotASigner() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.NotASigner.selector);
        custody.removeSigner(signer2, 1);
    }

    function test_Fail_RemoveSigner_LastSigner() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.CannotRemoveLastSigner.selector);
        custody.removeSigner(signer1, 1);
    }

    function test_Fail_RemoveSigner_InvalidQuorum() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1);

        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.InvalidQuorum.selector);
        custody.removeSigner(signer2, 2); // removing leaves 1, max quorum is 1
    }

    function test_RemovedSignerCannotAct() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1);

        vm.prank(signer1);
        custody.removeSigner(signer2, 1);

        vm.prank(signer2);
        vm.expectRevert(QuorumCustody.NotSigner.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    // =========================================================================
    // Deposit
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
        vm.expectRevert(ICustody.ZeroAmount.selector);
        custody.deposit(address(0), 0);
    }

    function test_Fail_DepositETH_MsgValueMismatch() public {
        vm.deal(user, 2 ether);
        vm.prank(user);
        vm.expectRevert(ICustody.MsgValueMismatch.selector);
        custody.deposit{value: 0.5 ether}(address(0), 1 ether);
    }

    function test_Fail_DepositERC20_NonZeroMsgValue() public {
        token.mint(user, 100e18);
        vm.deal(user, 1 ether);
        vm.startPrank(user);
        token.approve(address(custody), 100e18);
        vm.expectRevert(ICustody.NonZeroMsgValueForERC20.selector);
        custody.deposit{value: 1 ether}(address(token), 100e18);
        vm.stopPrank();
    }

    // =========================================================================
    // startWithdraw
    // =========================================================================

    function test_Fail_StartWithdraw_NotSigner() public {
        vm.prank(user);
        vm.expectRevert(QuorumCustody.NotSigner.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_Fail_StartWithdraw_ZeroUser() public {
        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.InvalidUser.selector);
        custody.startWithdraw(address(0), address(0), 1 ether, 1);
    }

    function test_Fail_StartWithdraw_ZeroAmount() public {
        vm.prank(signer1);
        vm.expectRevert(ICustody.ZeroAmount.selector);
        custody.startWithdraw(user, address(0), 0, 1);
    }

    function test_Fail_StartWithdraw_DuplicateNonce() public {
        vm.startPrank(signer1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.expectRevert(ICustody.WithdrawalAlreadyExists.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.stopPrank();
    }

    function test_StartWithdraw_SameParamsDifferentNonce() public {
        vm.startPrank(signer1);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        vm.stopPrank();
        assertTrue(id1 != id2);
    }

    function test_StartWithdraw_EmitsEvent() public {
        vm.prank(signer1);
        vm.expectEmit(true, true, true, true);
        bytes32 expectedId =
            keccak256(abi.encode(block.chainid, address(custody), user, address(0), 1 ether, uint256(1)));
        emit ICustody.WithdrawStarted(expectedId, user, address(0), 1 ether, 1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_StartWithdraw_SnapshotsQuorum() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        (,,,,,, uint256 requiredQuorum,) = custody.withdrawals(id);
        assertEq(requiredQuorum, 1);
    }

    // =========================================================================
    // finalizeWithdraw — 1/1
    // =========================================================================

    function test_FinalizeWithdraw_1_1() public {
        vm.deal(address(custody), 1 ether);

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — 2/2 progressive
    // =========================================================================

    function test_FinalizeWithdraw_2_2_Progressive() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);
        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);
        (,,,, finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — 3/5
    // =========================================================================

    function test_FinalizeWithdraw_3_5() public {
        _setup3of5();
        assertEq(custody.quorum(), 3);
        assertEq(custody.getSignerCount(), 5);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);
        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);
        (,,,, finalized,,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer3);
        custody.finalizeWithdraw(id);
        (,,,, finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — snapshot quorum
    // =========================================================================

    function test_FinalizeWithdraw_UsesSnapshotQuorum() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 1); // quorum=1

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Raise quorum to 2 AFTER withdrawal was created
        vm.prank(signer1);
        custody.addSigner(signer3, 2);
        assertEq(custody.quorum(), 2);

        // 1 approval should suffice (snapshot quorum was 1)
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — edge cases
    // =========================================================================

    function test_Fail_FinalizeWithdraw_NotSigner() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(user);
        vm.expectRevert(QuorumCustody.NotSigner.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert(ICustody.WithdrawalNotFound.selector);
        custody.finalizeWithdraw(bytes32(uint256(999)));
    }

    function test_Fail_FinalizeWithdraw_AlreadyFinalized() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(ICustody.WithdrawalAlreadyFinalized.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_DuplicateApproval() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.SignerAlreadyApproved.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_Finalize_Expired() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.WithdrawalExpired.selector);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ExactExpiryBoundary() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_FinalizeWithdraw_ERC20() public {
        token.mint(address(custody), 50e18);

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        assertEq(token.balanceOf(user), 50e18);
        assertEq(token.balanceOf(address(custody)), 0);
    }

    function test_Fail_FinalizeWithdraw_InsufficientETH() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectRevert(ICustody.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_InsufficientERC20() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        vm.prank(signer1);
        vm.expectRevert(ICustody.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ClearsStorage() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (address storedUser, address storedToken, uint256 storedAmount, bool exists, bool finalized,,,) =
            custody.withdrawals(id);
        assertTrue(exists);
        assertTrue(finalized);
        assertEq(storedUser, address(0));
        assertEq(storedToken, address(0));
        assertEq(storedAmount, 0);
    }

    function test_FinalizeWithdraw_EmitsApprovalEvent() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, true, false, true);
        emit QuorumCustody.WithdrawalApproved(id, signer1, 1);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_EmitsFinalizedEvent() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit ICustody.WithdrawFinalized(id, true);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ETH_UserReceivesBalance() public {
        vm.deal(address(custody), 5 ether);
        uint256 balanceBefore = user.balance;

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 2 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        assertEq(user.balance, balanceBefore + 2 ether);
        assertEq(address(custody).balance, 3 ether);
    }

    // =========================================================================
    // rejectWithdraw (expired-only cleanup)
    // =========================================================================

    function test_RejectWithdraw_Expired() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_RejectWithdraw_EmitsEvent() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit ICustody.WithdrawFinalized(id, false);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NotExpired() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectRevert(QuorumCustody.WithdrawalNotExpired.selector);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert(ICustody.WithdrawalNotFound.selector);
        custody.rejectWithdraw(bytes32(uint256(999)));
    }

    function test_Fail_RejectWithdraw_AlreadyFinalized() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        vm.expectRevert(ICustody.WithdrawalAlreadyFinalized.selector);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NotSigner() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(user);
        vm.expectRevert(QuorumCustody.NotSigner.selector);
        custody.rejectWithdraw(id);
    }

    // =========================================================================
    // Lifecycle: reject expired, then re-create
    // =========================================================================

    function test_RejectExpiredThenRecreate() public {
        vm.deal(address(custody), 1 ether);

        vm.prank(signer1);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        custody.rejectWithdraw(id1);

        vm.prank(signer1);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        assertTrue(id1 != id2);

        vm.prank(signer1);
        custody.finalizeWithdraw(id2);

        assertEq(user.balance, 1 ether);
    }

    // =========================================================================
    // Partial approval then expiry
    // =========================================================================

    function test_PartialApprovalThenExpiry() public {
        vm.prank(signer1);
        custody.addSigner(signer2, 2);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer2);
        vm.expectRevert(QuorumCustody.WithdrawalExpired.selector);
        custody.finalizeWithdraw(id);

        // Clean up expired
        vm.prank(signer1);
        custody.rejectWithdraw(id);

        (,,,, bool finalized,,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // Multiple concurrent withdrawals
    // =========================================================================

    function test_MultipleConcurrentWithdrawals() public {
        vm.deal(address(custody), 3 ether);

        vm.startPrank(signer1);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);
        bytes32 id2 = custody.startWithdraw(user, address(0), 1 ether, 2);
        bytes32 id3 = custody.startWithdraw(user, address(0), 1 ether, 3);
        vm.stopPrank();

        vm.prank(signer1);
        custody.finalizeWithdraw(id1);

        // id2 left to expire
        vm.prank(signer1);
        custody.finalizeWithdraw(id3);

        assertEq(user.balance, 2 ether);
        assertEq(address(custody).balance, 1 ether);
    }

    // =========================================================================
    // getSignerCount
    // =========================================================================

    function test_GetSignerCount() public {
        assertEq(custody.getSignerCount(), 1);

        vm.prank(signer1);
        custody.addSigner(signer2, 1);
        assertEq(custody.getSignerCount(), 2);

        vm.prank(signer1);
        custody.addSigner(signer3, 1);
        assertEq(custody.getSignerCount(), 3);
    }
}
