// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {SimpleCustody} from "../src/SimpleCustody.sol";
import {IWithdraw} from "../src/interfaces/IWithdraw.sol";
import {IDeposit} from "../src/interfaces/IDeposit.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor() ERC20("Mock", "MCK") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract SimpleCustodyTest is Test {
    SimpleCustody public custody;
    MockERC20 public token;

    address internal admin;
    address internal neodax;
    address internal nitewatch;
    address internal user;

    bytes32 public constant NEODAX_ROLE = keccak256("NEODAX_ROLE");
    bytes32 public constant NITEWATCH_ROLE = keccak256("NITEWATCH_ROLE");

    function setUp() public {
        admin = makeAddr("admin");
        neodax = makeAddr("neodax");
        nitewatch = makeAddr("nitewatch");
        user = makeAddr("user");

        vm.startPrank(admin);
        custody = new SimpleCustody(admin, neodax, nitewatch);
        vm.stopPrank();

        token = new MockERC20();
    }

    // ---- deposit tests ----

    function test_depositETH() public {
        vm.deal(user, 1 ether);
        vm.prank(user);

        vm.expectEmit(true, true, false, true);
        emit IDeposit.Deposited(user, address(0), 1 ether);

        custody.deposit{value: 1 ether}(address(0), 1 ether);

        assertEq(address(custody).balance, 1 ether);
    }

    function test_depositERC20() public {
        token.mint(user, 100e18);

        vm.startPrank(user);
        token.approve(address(custody), 100e18);

        vm.expectEmit(true, true, false, true);
        emit IDeposit.Deposited(user, address(token), 100e18);

        custody.deposit(address(token), 100e18);
        vm.stopPrank();

        assertEq(token.balanceOf(address(custody)), 100e18);
        assertEq(token.balanceOf(user), 0);
    }

    function test_depositETH_wrongAmount() public {
        vm.deal(user, 2 ether);
        vm.prank(user);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: 1 ether}(address(0), 2 ether);
    }

    function test_depositERC20_withValue() public {
        token.mint(user, 100e18);
        vm.deal(user, 1 ether);

        vm.startPrank(user);
        token.approve(address(custody), 100e18);

        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: 1 ether}(address(token), 100e18);
        vm.stopPrank();
    }

    // ---- startWithdraw tests ----

    function test_startWithdraw() public {
        uint256 nonce = 1;
        uint256 amount = 1 ether;

        vm.startPrank(neodax);

        bytes32 expectedId = keccak256(abi.encode(block.chainid, address(custody), user, address(0), amount, nonce));

        vm.expectEmit(true, true, false, true);
        emit IWithdraw.WithdrawStarted(expectedId, user, address(0), amount, nonce);

        bytes32 id = custody.startWithdraw(user, address(0), amount, nonce);

        assertEq(id, expectedId);

        // Use a tuple to decode the struct or access fields individually if public getter
        // Since `withdrawals` is public mapping, we can access fields via multiple return values
        (address u, address t, uint256 a, bool exists, bool finalized) = custody.withdrawals(id);
        assertEq(u, user);
        assertEq(t, address(0));
        assertEq(a, amount);
        assertTrue(exists);
        assertFalse(finalized);

        vm.stopPrank();
    }

    function test_deposit_zeroAmount() public {
        vm.prank(user);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.deposit{value: 0}(address(0), 0);
    }

    function test_startWithdraw_zeroAmount() public {
        vm.startPrank(neodax);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.startWithdraw(user, address(0), 0, 1);
        vm.stopPrank();
    }

    function test_finalizeWithdraw_storageCleared() public {
        // Setup: deposit first
        vm.deal(address(custody), 5 ether);

        // Start withdraw
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Finalize
        vm.startPrank(nitewatch);

        custody.finalizeWithdraw(id);

        (address u, address t, uint256 a, bool exists, bool finalized) = custody.withdrawals(id);
        assertEq(u, address(0)); // Cleared
        assertEq(t, address(0)); // Cleared
        assertEq(a, 0); // Cleared
        assertTrue(exists); // Preserved
        assertTrue(finalized); // Preserved

        vm.stopPrank();
    }

    function test_startWithdraw_unauthorized() public {
        vm.startPrank(user); // not neodax
        vm.expectRevert(); // Should revert with AccessControl error, but checking generic revert for simplicity or specific error
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.stopPrank();
    }

    function test_startWithdraw_duplicate() public {
        vm.startPrank(neodax);
        custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.expectRevert(IWithdraw.WithdrawalAlreadyExists.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.stopPrank();
    }

    // ---- finalizeWithdraw tests ----

    function test_finalizeWithdrawETH() public {
        // Setup: deposit first
        vm.deal(address(this), 5 ether); // test contract funding
        // We need to fund the custody contract
        vm.deal(address(custody), 5 ether);

        // Start withdraw
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Finalize
        vm.startPrank(nitewatch);

        uint256 preBalance = user.balance;

        vm.expectEmit(true, true, false, true);
        emit IWithdraw.WithdrawFinalized(id, true);

        custody.finalizeWithdraw(id);

        assertEq(user.balance, preBalance + 1 ether);

        (,,,, bool finalized) = custody.withdrawals(id);
        assertTrue(finalized);

        vm.stopPrank();
    }

    function test_finalizeWithdrawERC20() public {
        // Setup: deposit first
        token.mint(address(custody), 100e18);

        // Start withdraw
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        // Finalize
        vm.startPrank(nitewatch);

        uint256 preBalance = token.balanceOf(user);

        vm.expectEmit(true, true, false, true);
        emit IWithdraw.WithdrawFinalized(id, true);

        custody.finalizeWithdraw(id);

        assertEq(token.balanceOf(user), preBalance + 50e18);

        (,,,, bool finalized) = custody.withdrawals(id);
        assertTrue(finalized);

        vm.stopPrank();
    }

    function test_finalizeWithdraw_unauthorized() public {
        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.startPrank(user); // not nitewatch
        vm.expectRevert();
        custody.finalizeWithdraw(id);
        vm.stopPrank();
    }

    function test_finalizeWithdraw_notFound() public {
        vm.startPrank(nitewatch);
        vm.expectRevert(IWithdraw.WithdrawalNotFound.selector);
        custody.finalizeWithdraw(bytes32(uint256(999))); // random id
        vm.stopPrank();
    }

    function test_finalizeWithdraw_alreadyFinalized() public {
        vm.deal(address(custody), 2 ether);

        vm.prank(neodax);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(nitewatch);
        custody.finalizeWithdraw(id);

        vm.startPrank(nitewatch);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.finalizeWithdraw(id);
        vm.stopPrank();
    }
}
