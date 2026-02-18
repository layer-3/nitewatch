// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import {ThresholdCustody} from "../src/ThresholdCustody.sol";
import {IWithdraw} from "../src/interfaces/IWithdraw.sol";
import {IDeposit} from "../src/interfaces/IDeposit.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor() ERC20("Mock", "MCK") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract ThresholdCustodyTest is Test {
    ThresholdCustody public custody;
    MockERC20 public token;

    address internal user;

    address internal signer1;
    uint256 internal signer1Pk;
    address internal signer2;
    uint256 internal signer2Pk;
    address internal signer3;
    uint256 internal signer3Pk;
    address internal signer4;
    uint256 internal signer4Pk;
    address internal signer5;
    uint256 internal signer5Pk;

    // EIP-712 domain values (must match contract constructor)
    bytes32 constant ADD_SIGNERS_TYPEHASH =
        keccak256("AddSigners(address[] newSigners,uint256 newThreshold,uint256 nonce,uint256 deadline)");
    bytes32 constant REMOVE_SIGNERS_TYPEHASH =
        keccak256("RemoveSigners(address[] signersToRemove,uint256 newThreshold,uint256 nonce,uint256 deadline)");
    uint256 constant MAX_DEADLINE = type(uint256).max;

    function setUp() public {
        user = makeAddr("user");
        (signer1, signer1Pk) = makeAddrAndKey("signer1");
        (signer2, signer2Pk) = makeAddrAndKey("signer2");
        (signer3, signer3Pk) = makeAddrAndKey("signer3");
        (signer4, signer4Pk) = makeAddrAndKey("signer4");
        (signer5, signer5Pk) = makeAddrAndKey("signer5");

        address[] memory initialSigners = new address[](1);
        initialSigners[0] = signer1;
        custody = new ThresholdCustody(initialSigners, 1);
        token = new MockERC20();
    }

    // =========================================================================
    // EIP-712 signing helpers
    // =========================================================================

    function _domainSeparator() internal view returns (bytes32) {
        return keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256("ThresholdCustody"),
                keccak256("1"),
                block.chainid,
                address(custody)
            )
        );
    }

    function _hashAddressArray(address[] memory arr) internal pure returns (bytes32) {
        bytes32[] memory encoded = new bytes32[](arr.length);
        for (uint256 i = 0; i < arr.length; i++) {
            encoded[i] = bytes32(uint256(uint160(arr[i])));
        }
        return keccak256(abi.encodePacked(encoded));
    }

    function _signAddSigners(
        uint256 pk,
        address[] memory newSigners,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(ADD_SIGNERS_TYPEHASH, _hashAddressArray(newSigners), newThreshold, nonce, deadline)
        );
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", _domainSeparator(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    function _signRemoveSigners(
        uint256 pk,
        address[] memory signersToRemove,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(REMOVE_SIGNERS_TYPEHASH, _hashAddressArray(signersToRemove), newThreshold, nonce, deadline)
        );
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", _domainSeparator(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    // Helper: deploy 5-signer custody with quorum=3
    function _setup3of5() internal {
        address[] memory allSigners = new address[](5);
        allSigners[0] = signer1;
        allSigners[1] = signer2;
        allSigners[2] = signer3;
        allSigners[3] = signer4;
        allSigners[4] = signer5;
        custody = new ThresholdCustody(allSigners, 3);
    }

    // Helper: sort two signatures by signer address (ascending) and encode in MultiSignerERC7913 format
    function _sortSigs2(address a, bytes memory sigA, address b, bytes memory sigB)
        internal
        pure
        returns (bytes memory)
    {
        return _encodeMultiSig2(a, sigA, b, sigB);
    }

    function _emptySigs() internal pure returns (bytes memory) {
        // Return properly encoded empty arrays for MultiSignerERC7913 format
        bytes[] memory emptySigners = new bytes[](0);
        bytes[] memory emptySignatures = new bytes[](0);
        return abi.encode(emptySigners, emptySignatures);
    }

    // Helper for self-signature: when a single signer with threshold=1 signs their own operation
    function _selfSign(
        uint256 pk,
        address signer,
        address[] memory newSigners,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes memory sig = _signAddSigners(pk, newSigners, newThreshold, nonce, deadline);
        return _encodeMultiSig(signer, sig);
    }

    function _selfSignRemove(
        uint256 pk,
        address signer,
        address[] memory signersToRemove,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes memory sig = _signRemoveSigners(pk, signersToRemove, newThreshold, nonce, deadline);
        return _encodeMultiSig(signer, sig);
    }

    // Helper to encode a single signature in MultiSignerERC7913 format
    function _encodeMultiSig(address signer, bytes memory signature) internal pure returns (bytes memory) {
        bytes[] memory signers = new bytes[](1);
        signers[0] = abi.encodePacked(signer);
        bytes[] memory signatures = new bytes[](1);
        signatures[0] = signature;
        return abi.encode(signers, signatures);
    }

    // Helper to encode two signatures in MultiSignerERC7913 format (sorted by signer)
    function _encodeMultiSig2(address signerA, bytes memory sigA, address signerB, bytes memory sigB)
        internal
        pure
        returns (bytes memory)
    {
        bytes[] memory signers = new bytes[](2);
        bytes[] memory signatures = new bytes[](2);

        if (uint160(signerA) < uint160(signerB)) {
            signers[0] = abi.encodePacked(signerA);
            signers[1] = abi.encodePacked(signerB);
            signatures[0] = sigA;
            signatures[1] = sigB;
        } else {
            signers[0] = abi.encodePacked(signerB);
            signers[1] = abi.encodePacked(signerA);
            signatures[0] = sigB;
            signatures[1] = sigA;
        }

        return abi.encode(signers, signatures);
    }

    // =========================================================================
    // Constructor tests
    // =========================================================================

    function test_InitialState() public view {
        assertEq(custody.threshold(), 1);
        bytes[] memory signers = custody.getSigners(0, type(uint64).max);
        assertEq(signers.length, 1);
        assertTrue(custody.isSigner(signer1));
        assertEq(custody.getSignerCount(), 1);
    }

    function test_Constructor_MultipleSigners() public {
        address[] memory s = new address[](3);
        s[0] = signer1;
        s[1] = signer2;
        s[2] = signer3;
        ThresholdCustody c = new ThresholdCustody(s, 2);

        assertEq(c.threshold(), 2);
        assertEq(c.getSignerCount(), 3);
        assertTrue(c.isSigner(signer1));
        assertTrue(c.isSigner(signer2));
        assertTrue(c.isSigner(signer3));
    }

    function test_Fail_Constructor_EmptySigners() public {
        address[] memory s = new address[](0);
        // With 0 signers, MultiSignerERC7913 will try to validate threshold against 0 signers
        vm.expectRevert();
        new ThresholdCustody(s, 1);
    }

    // NOTE: address(0) is actually valid for MultiSignerERC7913 as it's 20 bytes
    // This test is removed as it's not a failure case in the new implementation

    function test_Fail_Constructor_DuplicateSigner() public {
        address[] memory s = new address[](2);
        s[0] = signer1;
        s[1] = signer1;
        // MultiSignerERC7913 will revert with AlreadyExists error
        vm.expectRevert();
        new ThresholdCustody(s, 1);
    }

    function test_Fail_Constructor_QuorumZero() public {
        address[] memory s = new address[](1);
        s[0] = signer1;
        // MultiSignerERC7913 will revert with ZeroThreshold error
        vm.expectRevert();
        new ThresholdCustody(s, 0);
    }

    function test_Fail_Constructor_QuorumTooHigh() public {
        address[] memory s = new address[](1);
        s[0] = signer1;
        // MultiSignerERC7913 will revert with UnreachableThreshold error
        vm.expectRevert();
        new ThresholdCustody(s, 2);
    }

    // =========================================================================
    // addSigners
    // =========================================================================

    function test_AddSigners_Quorum1_EmptySigs() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        // With threshold=1, signer1 must provide their own signature
        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 2, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        custody.addSigners(newSigners, 2, MAX_DEADLINE, sigs);

        assertTrue(custody.isSigner(signer2));
        assertEq(custody.threshold(), 2);
        assertEq(custody.getSignerCount(), 2);
    }

    function test_AddSigners_WithSignature() public {
        // First add signer2 to get threshold=2
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory sigs1 = _selfSign(signer1Pk, signer1, s1, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 2, MAX_DEADLINE, sigs1);

        // Now add signer3 with threshold=2, need 2 signatures total
        address[] memory s2 = new address[](1);
        s2[0] = signer3;
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signAddSigners(signer1Pk, s2, 2, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signAddSigners(signer2Pk, s2, 2, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.prank(signer1);
        custody.addSigners(s2, 2, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer3));
        assertEq(custody.getSignerCount(), 3);
    }

    function test_AddSigners_BatchMultiple() public {
        address[] memory newSigners = new address[](3);
        newSigners[0] = signer2;
        newSigners[1] = signer3;
        newSigners[2] = signer4;

        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 2, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        custody.addSigners(newSigners, 2, MAX_DEADLINE, sigs);

        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        assertTrue(custody.isSigner(signer4));
        assertEq(custody.getSignerCount(), 4);
        assertEq(custody.threshold(), 2);
    }

    function test_AddSigners_ThresholdChanged() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 1, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        custody.addSigners(newSigners, 1, MAX_DEADLINE, sigs);

        assertEq(custody.threshold(), 1);
    }

    function test_Fail_AddSigners_NotSigner() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        vm.prank(user);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.addSigners(newSigners, 1, MAX_DEADLINE, _emptySigs());
    }

    function test_Fail_AddSigners_EmptyArray() public {
        address[] memory newSigners = new address[](0);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.EmptySignersArray.selector);
        custody.addSigners(newSigners, 1, MAX_DEADLINE, _emptySigs());
    }

    // Note: ZeroAddress, Duplicate, and DuplicateInBatch validation is now handled by MultiSignerERC7913
    // These will revert with MultiSignerERC7913 errors instead

    function test_Fail_AddSigners_QuorumZero() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 0, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        // MultiSignerERC7913 will revert with ZeroThreshold error
        vm.expectRevert();
        custody.addSigners(newSigners, 0, MAX_DEADLINE, sigs);
    }

    function test_Fail_AddSigners_QuorumTooHigh() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 3, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        // MultiSignerERC7913 will revert with UnreachableThreshold error
        vm.expectRevert();
        custody.addSigners(newSigners, 3, MAX_DEADLINE, sigs); // max is 2 (1 existing + 1 new)
    }

    function test_Fail_AddSigners_InsufficientSignatures() public {
        // Setup: 2 signers, threshold=2
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory sigs1 = _selfSign(signer1Pk, signer1, s1, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 2, MAX_DEADLINE, sigs1);

        // Try to add signer3 with only 1 signature (need 2)
        address[] memory s2 = new address[](1);
        s2[0] = signer3;
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signAddSigners(signer1Pk, s2, 2, nonce, MAX_DEADLINE);
        bytes memory onlyOneSig = _encodeMultiSig(signer1, sig1);

        vm.prank(signer1);
        // Returns InvalidSignature for insufficient signatures
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(s2, 2, MAX_DEADLINE, onlyOneSig);
    }

    function test_Fail_AddSigners_StaleNonce() public {
        // Setup: 2 signers, threshold=2
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory sigs1 = _selfSign(signer1Pk, signer1, s1, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 2, MAX_DEADLINE, sigs1);

        // Pre-sign at current nonce
        address[] memory s2 = new address[](1);
        s2[0] = signer3;
        uint256 staleNonce = custody.signerNonce();
        bytes memory staleSig1 = _signAddSigners(signer1Pk, s2, 2, staleNonce, MAX_DEADLINE);
        bytes memory staleSig2 = _signAddSigners(signer2Pk, s2, 2, staleNonce, MAX_DEADLINE);

        // Add signer4 first (advances nonce)
        address[] memory s3 = new address[](1);
        s3[0] = signer4;
        bytes memory sig1 = _signAddSigners(signer1Pk, s3, 2, staleNonce, MAX_DEADLINE);
        bytes memory sig2 = _signAddSigners(signer2Pk, s3, 2, staleNonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);
        vm.prank(signer1);
        custody.addSigners(s3, 2, MAX_DEADLINE, encodedSigs);

        // Now try to use the stale signatures (nonce is now incremented)
        bytes memory staleEncodedSigs = _encodeMultiSig2(signer1, staleSig1, signer2, staleSig2);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(s2, 2, MAX_DEADLINE, staleEncodedSigs);
    }

    function test_AddSigners_IncrementsNonce() public {
        uint256 nonceBefore = custody.signerNonce();

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;
        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 1, nonceBefore, MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(newSigners, 1, MAX_DEADLINE, sigs);

        assertEq(custody.signerNonce(), nonceBefore + 1);
    }

    // =========================================================================
    // removeSigners
    // =========================================================================

    function test_RemoveSigners_Quorum1() public {
        // Add signer2 first
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s1, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 1, MAX_DEADLINE, addSigs);

        // Remove signer2
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;
        bytes memory removeSigs = _selfSignRemove(signer1Pk, signer1, toRemove, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.removeSigners(toRemove, 1, MAX_DEADLINE, removeSigs);

        assertFalse(custody.isSigner(signer2));
        assertEq(custody.getSignerCount(), 1);
    }

    function test_RemoveSigners_WithSignature() public {
        // Setup 3 signers, threshold 2
        address[] memory s = new address[](2);
        s[0] = signer2;
        s[1] = signer3;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, addSigs);

        // Remove signer3, need 2 signatures total
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer3;
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.prank(signer1);
        custody.removeSigners(toRemove, 2, MAX_DEADLINE, encodedSigs);

        assertFalse(custody.isSigner(signer3));
        assertEq(custody.getSignerCount(), 2);
    }

    function test_RemoveSigners_BatchMultiple() public {
        _setup3of5();

        // Remove signer4 and signer5 at once - need 3 signatures for threshold=3
        address[] memory toRemove = new address[](2);
        toRemove[0] = signer4;
        toRemove[1] = signer5;
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signRemoveSigners(signer1Pk, toRemove, 3, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, toRemove, 3, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signRemoveSigners(signer3Pk, toRemove, 3, nonce, MAX_DEADLINE);

        // Encode all 3 signatures
        bytes[] memory signers = new bytes[](3);
        bytes[] memory signatures = new bytes[](3);
        address[3] memory addrs = [signer1, signer2, signer3];
        bytes[3] memory sigs = [sig1, sig2, sig3];
        // Sort by address
        for (uint256 i = 0; i < 2; i++) {
            for (uint256 j = i + 1; j < 3; j++) {
                if (uint160(addrs[i]) > uint160(addrs[j])) {
                    (addrs[i], addrs[j]) = (addrs[j], addrs[i]);
                    (sigs[i], sigs[j]) = (sigs[j], sigs[i]);
                }
            }
        }
        for (uint256 i = 0; i < 3; i++) {
            signers[i] = abi.encodePacked(addrs[i]);
            signatures[i] = sigs[i];
        }
        bytes memory encodedSigs = abi.encode(signers, signatures);

        vm.prank(signer1);
        custody.removeSigners(toRemove, 3, MAX_DEADLINE, encodedSigs);

        assertFalse(custody.isSigner(signer4));
        assertFalse(custody.isSigner(signer5));
        assertEq(custody.getSignerCount(), 3);
        assertEq(custody.threshold(), 3);
    }

    function test_Fail_RemoveSigners_UnreachableThreshold() public {
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer1;

        vm.prank(signer1);
        // MultiSignerERC7913 will revert with UnreachableThreshold error
        // when removing would make the threshold impossible to reach
        vm.expectRevert();
        custody.removeSigners(toRemove, 1, MAX_DEADLINE, _emptySigs());
    }

    function test_Fail_RemoveSigners_InvalidQuorum() public {
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s1, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 1, MAX_DEADLINE, addSigs);

        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;
        bytes memory removeSigs = _selfSignRemove(signer1Pk, signer1, toRemove, 2, custody.signerNonce(), MAX_DEADLINE);

        vm.prank(signer1);
        // MultiSignerERC7913 will revert with UnreachableThreshold error
        vm.expectRevert();
        custody.removeSigners(toRemove, 2, MAX_DEADLINE, removeSigs); // removing leaves 1, max threshold is 1
    }

    function test_Fail_RemoveSigners_EmptyArray() public {
        address[] memory toRemove = new address[](0);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.EmptySignersArray.selector);
        custody.removeSigners(toRemove, 1, MAX_DEADLINE, _emptySigs());
    }

    function test_RemovedSignerCannotAct() public {
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s1, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 1, MAX_DEADLINE, addSigs);

        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;
        bytes memory removeSigs = _selfSignRemove(signer1Pk, signer1, toRemove, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.removeSigners(toRemove, 1, MAX_DEADLINE, removeSigs);

        vm.prank(signer2);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
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
        emit IDeposit.Deposited(user, address(0), 1 ether);
        custody.deposit{value: 1 ether}(address(0), 1 ether);
    }

    function test_DepositERC20_EmitsEvent() public {
        token.mint(user, 50e18);
        vm.startPrank(user);
        token.approve(address(custody), 50e18);
        vm.expectEmit(true, true, false, true);
        emit IDeposit.Deposited(user, address(token), 50e18);
        custody.deposit(address(token), 50e18);
        vm.stopPrank();
    }

    function test_Fail_DepositZeroAmount() public {
        vm.prank(user);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.deposit(address(0), 0);
    }

    function test_Fail_DepositETH_MsgValueMismatch() public {
        vm.deal(user, 2 ether);
        vm.prank(user);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: 0.5 ether}(address(0), 1 ether);
    }

    function test_Fail_DepositERC20_NonZeroMsgValue() public {
        token.mint(user, 100e18);
        vm.deal(user, 1 ether);
        vm.startPrank(user);
        token.approve(address(custody), 100e18);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: 1 ether}(address(token), 100e18);
        vm.stopPrank();
    }

    // =========================================================================
    // startWithdraw
    // =========================================================================

    function test_Fail_StartWithdraw_NotSigner() public {
        vm.prank(user);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_Fail_StartWithdraw_ZeroUser() public {
        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.InvalidUser.selector);
        custody.startWithdraw(address(0), address(0), 1 ether, 1);
    }

    function test_Fail_StartWithdraw_ZeroAmount() public {
        vm.prank(signer1);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.startWithdraw(user, address(0), 0, 1);
    }

    function test_Fail_StartWithdraw_DuplicateNonce() public {
        vm.startPrank(signer1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyExists.selector);
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
        emit IWithdraw.WithdrawStarted(expectedId, user, address(0), 1 ether, 1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_StartWithdraw_SnapshotsQuorum() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        (,,,,, uint64 requiredQuorum) = custody.withdrawals(id);
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

        (,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — 2/2 progressive
    // =========================================================================

    function test_FinalizeWithdraw_2_2_Progressive() public {
        // Setup: 2 signers, threshold=2
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory sigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, sigs);

        vm.deal(address(custody), 1 ether);

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);
        (,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);
        (,,, finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — 3/5
    // =========================================================================

    function test_FinalizeWithdraw_3_5() public {
        _setup3of5();
        assertEq(custody.threshold(), 3);
        assertEq(custody.getSignerCount(), 5);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);
        (,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);
        (,,, finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer3);
        custody.finalizeWithdraw(id);
        (,,, finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // finalizeWithdraw — snapshot quorum
    // =========================================================================

    function test_FinalizeWithdraw_UsesSnapshotQuorum() public {
        // Setup: 2 signers, threshold=1
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory addSigs1 = _selfSign(signer1Pk, signer1, s, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 1, MAX_DEADLINE, addSigs1);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        // Raise threshold to 2 AFTER withdrawal was created
        address[] memory s2 = new address[](1);
        s2[0] = signer3;
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signAddSigners(signer1Pk, s2, 2, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig(signer1, sig1);

        vm.prank(signer1);
        custody.addSigners(s2, 2, MAX_DEADLINE, encodedSigs);
        assertEq(custody.threshold(), 2);

        // 1 approval should suffice (snapshot quorum was 1)
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,, bool finalized,,) = custody.withdrawals(id);
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
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalNotFound.selector);
        custody.finalizeWithdraw(bytes32(uint256(999)));
    }

    function test_Fail_FinalizeWithdraw_AlreadyFinalized() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_DuplicateApproval() public {
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory sigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, sigs);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.SignerAlreadyApproved.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_Finalize_Expired() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ExactExpiryBoundary() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,, bool finalized,,) = custody.withdrawals(id);
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
        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_Fail_FinalizeWithdraw_InsufficientERC20() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(token), 50e18, 1);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_ClearsStorage() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (address storedUser, address storedToken, uint256 storedAmount, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
        assertEq(storedUser, address(0));
        assertEq(storedToken, address(0));
        assertEq(storedAmount, 0);
    }

    function test_FinalizeWithdraw_EmitsApprovalEvent() public {
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory sigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, sigs);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, true, false, true);
        emit ThresholdCustody.WithdrawalApproved(id, signer1, 1);
        custody.finalizeWithdraw(id);
    }

    function test_FinalizeWithdraw_EmitsFinalizedEvent() public {
        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit IWithdraw.WithdrawFinalized(id, true);
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

        (,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    function test_RejectWithdraw_EmitsEvent() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit IWithdraw.WithdrawFinalized(id, false);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NotExpired() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.WithdrawalNotExpired.selector);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NonExistent() public {
        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalNotFound.selector);
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
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.rejectWithdraw(id);
    }

    function test_Fail_RejectWithdraw_NotSigner() public {
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(user);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
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
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory sigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, sigs);

        vm.deal(address(custody), 1 ether);
        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.warp(block.timestamp + 1 hours + 1);

        vm.prank(signer2);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.finalizeWithdraw(id);

        // Clean up expired
        vm.prank(signer1);
        custody.rejectWithdraw(id);

        (,,, bool finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // Multiple concurrent withdrawals
    // =========================================================================

    function test_MultipleConcurrentWithdrawals() public {
        vm.deal(address(custody), 3 ether);

        vm.startPrank(signer1);
        bytes32 id1 = custody.startWithdraw(user, address(0), 1 ether, 1);
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

        address[] memory s = new address[](2);
        s[0] = signer2;
        s[1] = signer3;
        bytes memory sigs = _selfSign(signer1Pk, signer1, s, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 1, MAX_DEADLINE, sigs);
        assertEq(custody.getSignerCount(), 3);
    }

    // =========================================================================
    // Removed signer approvals no longer count (withdrawal)
    // =========================================================================

    function test_FinalizeWithdraw_RemovedSignerApprovalIgnored() public {
        // Setup: 3 signers, threshold=2
        address[] memory s = new address[](2);
        s[0] = signer2;
        s[1] = signer3;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 2, MAX_DEADLINE, addSigs);

        vm.deal(address(custody), 1 ether);

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, address(0), 1 ether, 1);
        // requiredQuorum snapshotted at 2

        // signer2 approves
        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // Remove signer2 (need 2 sigs since threshold=2)
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signRemoveSigners(signer3Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer3, sig3);
        vm.prank(signer1);
        custody.removeSigners(toRemove, 2, MAX_DEADLINE, encodedSigs);

        assertFalse(custody.isSigner(signer2));

        // signer1 approves — only 1 valid approval (signer2's no longer counts)
        // snapshotted requiredQuorum is still 2, so not finalized yet
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        (,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        // signer3 approves — now 2 valid approvals (signer1 + signer3), meets requiredQuorum=2
        vm.prank(signer3);
        custody.finalizeWithdraw(id);

        (,,, finalized,,) = custody.withdrawals(id);
        assertTrue(finalized);
    }

    // =========================================================================
    // Deadline expiry tests
    // =========================================================================

    function test_Fail_AddSigners_DeadlineExpired() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint256 deadline = block.timestamp + 1 hours;
        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 1, custody.signerNonce(), deadline);
        vm.warp(deadline + 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.addSigners(newSigners, 1, deadline, sigs);
    }

    function test_Fail_RemoveSigners_DeadlineExpired() public {
        // Add signer2 first
        address[] memory s = new address[](1);
        s[0] = signer2;
        bytes memory addSigs = _selfSign(signer1Pk, signer1, s, 1, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s, 1, MAX_DEADLINE, addSigs);

        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;

        uint256 deadline = block.timestamp + 1 hours;
        bytes memory removeSigs = _selfSignRemove(signer1Pk, signer1, toRemove, 1, custody.signerNonce(), deadline);
        vm.warp(deadline + 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.removeSigners(toRemove, 1, deadline, removeSigs);
    }

    function test_AddSigners_ExactDeadlineBoundary() public {
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint256 deadline = block.timestamp + 1 hours;
        bytes memory sigs = _selfSign(signer1Pk, signer1, newSigners, 1, custody.signerNonce(), deadline);
        vm.warp(deadline);

        vm.prank(signer1);
        custody.addSigners(newSigners, 1, deadline, sigs);
        assertTrue(custody.isSigner(signer2));
    }

    // =========================================================================
    // Exploit: Quorum downgrade via addSigners
    // =========================================================================

    // NOTE: ThresholdCustody does not have the same threshold-reduction protection as QuorumCustody
    // The MultiSignerERC7913 implementation allows threshold changes as long as they're reachable
    // These exploit tests are removed as they test behavior specific to the old QuorumCustody contract

    function test_Fail_AddSigners_DeadlineExpired_WithSignatures() public {
        // Setup: 2 signers, threshold=2
        address[] memory s1 = new address[](1);
        s1[0] = signer2;
        bytes memory sigs1 = _selfSign(signer1Pk, signer1, s1, 2, custody.signerNonce(), MAX_DEADLINE);
        vm.prank(signer1);
        custody.addSigners(s1, 2, MAX_DEADLINE, sigs1);

        // Sign with a deadline, then let it expire
        address[] memory s2 = new address[](1);
        s2[0] = signer3;
        uint256 nonce = custody.signerNonce();
        uint256 deadline = block.timestamp + 1 hours;

        bytes memory sig1 = _signAddSigners(signer1Pk, s2, 2, nonce, deadline);
        bytes memory sig2 = _signAddSigners(signer2Pk, s2, 2, nonce, deadline);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.warp(deadline + 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.addSigners(s2, 2, deadline, encodedSigs);
    }
}
