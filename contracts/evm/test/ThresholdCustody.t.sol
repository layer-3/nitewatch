// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MultiSignerERC7913} from "@openzeppelin/contracts/utils/cryptography/signers/MultiSignerERC7913.sol";

import {
    ThresholdCustody,
    SET_THRESHOLD_TYPEHASH,
    ADD_SIGNERS_TYPEHASH,
    REMOVE_SIGNERS_TYPEHASH,
    OPERATION_EXPIRY
} from "../src/ThresholdCustody.sol";
import {IWithdraw} from "../src/interfaces/IWithdraw.sol";
import {IDeposit} from "../src/interfaces/IDeposit.sol";
import {Utils} from "../src/Utils.sol";
import {TestThresholdCustody} from "./TestThresholdCustody.sol";

using {Utils.toBytes} for address;

contract MockERC20 is ERC20 {
    constructor() ERC20("Mock", "MCK") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract ThresholdCustodyTest_Base is Test {
    using {Utils.hashArrayed} for address[];

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

    address internal notSigner;
    uint256 internal notSignerPk;

    uint256 constant MAX_DEADLINE = type(uint256).max;

    address[] oneSigner = new address[](1);
    address[] twoSigners = new address[](2);
    address[] threeSigners = new address[](3);
    address[] fiveSigners = new address[](5);

    function setUp() public virtual {
        user = makeAddr("user");
        (signer1, signer1Pk) = makeAddrAndKey("signer1");
        (signer2, signer2Pk) = makeAddrAndKey("signer2");
        (signer3, signer3Pk) = makeAddrAndKey("signer3");
        (signer4, signer4Pk) = makeAddrAndKey("signer4");
        (signer5, signer5Pk) = makeAddrAndKey("signer5");
        (notSigner, notSignerPk) = makeAddrAndKey("notSigner");

        oneSigner[0] = signer1;
        twoSigners[0] = signer1;
        twoSigners[1] = signer2;
        threeSigners[0] = signer1;
        threeSigners[1] = signer2;
        threeSigners[2] = signer3;
        fiveSigners[0] = signer1;
        fiveSigners[1] = signer2;
        fiveSigners[2] = signer3;
        fiveSigners[3] = signer4;
        fiveSigners[4] = signer5;
        token = new MockERC20();
    }

    // =========================================================================
    // EIP-712 signing helpers
    // =========================================================================

    function _domainSeparator(address contractAddress) internal view returns (bytes32) {
        return keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256("ThresholdCustody"),
                keccak256("1"),
                block.chainid,
                address(contractAddress)
            )
        );
    }

    function _domainSeparator() internal view returns (bytes32) {
        return _domainSeparator(address(custody));
    }

    function _signSetThreshold(uint256 pk, uint256 newThreshold, uint256 nonce, uint256 deadline)
        internal
        view
        returns (bytes memory)
    {
        bytes32 structHash = keccak256(abi.encode(SET_THRESHOLD_TYPEHASH, newThreshold, nonce, deadline));
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", _domainSeparator(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    function _signAddSigners(
        uint256 pk,
        address[] memory newSigners,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(ADD_SIGNERS_TYPEHASH, newSigners.hashArrayed(), newThreshold, nonce, deadline)
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
            abi.encode(REMOVE_SIGNERS_TYPEHASH, signersToRemove.hashArrayed(), newThreshold, nonce, deadline)
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
        allSigners[4] = notSigner;
        custody = new ThresholdCustody(allSigners, 3);
    }

    function _emptySigs() internal pure returns (bytes memory) {
        bytes[] memory emptySigners = new bytes[](0);
        bytes[] memory emptySignatures = new bytes[](0);
        return abi.encode(emptySigners, emptySignatures);
    }

    function _signSingleSetThreshold(uint256 pk, uint256 newThreshold, uint256 nonce, uint256 deadline)
        internal
        view
        returns (bytes memory)
    {
        address signer = vm.addr(pk);
        bytes memory sig = _signSetThreshold(pk, newThreshold, nonce, deadline);
        return _encodeMultiSig(signer, sig);
    }

    function _signSingleAdd(
        uint256 pk,
        address[] memory newSigners,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        address signer = vm.addr(pk);
        bytes memory sig = _signAddSigners(pk, newSigners, newThreshold, nonce, deadline);
        return _encodeMultiSig(signer, sig);
    }

    function _signSingleRemove(
        uint256 pk,
        address[] memory signersToRemove,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        address signer = vm.addr(pk);
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

    // do not sort as MultiSignerERC7913 accepts not sorted as well
    function _encodeMultiSig3(
        address signerA,
        bytes memory sigA,
        address signerB,
        bytes memory sigB,
        address signerC,
        bytes memory sigC
    ) internal pure returns (bytes memory) {
        bytes[] memory signers = new bytes[](3);
        bytes[] memory signatures = new bytes[](3);

        signers[0] = abi.encodePacked(signerA);
        signers[1] = abi.encodePacked(signerB);
        signers[2] = abi.encodePacked(signerC);

        signatures[0] = sigA;
        signatures[1] = sigB;
        signatures[2] = sigC;

        return abi.encode(signers, signatures);
    }

    function _checkStats(uint256 signersCount, uint64 threshold, uint256 newNonce) internal view {
        assertEq(custody.threshold(), threshold);
        assertEq(custody.getSignerCount(), signersCount);
        assertEq(custody.signerNonce(), newNonce);
    }

    function _validateWithdrawalData(
        bytes32 id,
        address expectedUser,
        address expectedToken,
        uint256 expectedAmount,
        bool expectedFinalized,
        uint64 expectedThreshold,
        uint64 expectedCreatedAt
    ) internal view virtual {
        (
            address storedUser,
            address storedToken,
            uint256 storedAmount,
            bool storedFinalized,
            uint64 storedCreatedAt,
            uint64 storedThreshold
        ) = custody.withdrawals(id);
        assertEq(storedUser, expectedUser);
        assertEq(storedToken, expectedToken);
        assertEq(storedAmount, expectedAmount);
        assertEq(storedFinalized, expectedFinalized);
        assertEq(storedThreshold, expectedThreshold);
        assertEq(storedCreatedAt, expectedCreatedAt);
    }
}

// =========================================================================
// Constructor tests
// =========================================================================
contract ThresholdCustodyTest_Constructor is ThresholdCustodyTest_Base {
    function test_singleSigner() public {
        address[] memory s = new address[](1);
        s[0] = signer1;
        ThresholdCustody c = new ThresholdCustody(s, 1);

        assertEq(c.threshold(), 1);
        bytes[] memory signers = c.getSigners(0, type(uint64).max);
        assertEq(signers.length, 1);
        assertEq(Utils.toAddress(signers[0]), signer1);
        assertTrue(c.isSigner(signer1));
        assertEq(c.getSignerCount(), 1);
    }

    function test_multipleSigners() public {
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

    function test_revert_emptySigners() public {
        address[] memory s = new address[](0);
        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 0, 1)
        );
        new ThresholdCustody(s, 1);
    }

    function test_revert_duplicateSigners() public {
        address[] memory s = new address[](2);
        s[0] = signer1;
        s[1] = signer1;
        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913AlreadyExists.selector, signer1.toBytes())
        );
        new ThresholdCustody(s, 1);
    }

    function test_revert_quorumZero() public {
        address[] memory s = new address[](1);
        s[0] = signer1;
        vm.expectRevert(MultiSignerERC7913.MultiSignerERC7913ZeroThreshold.selector);
        new ThresholdCustody(s, 0);
    }

    function test_revert_quorumNotReachable() public {
        address[] memory s = new address[](1);
        s[0] = signer1;
        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 1, 2)
        );
        new ThresholdCustody(s, 2);
    }
}

// =========================================================================
// setThreshold
// =========================================================================
contract ThresholdCustodyTest_SetThreshold is ThresholdCustodyTest_Base {
    function test_success_1_of_3_signature_increase() public {
        custody = new ThresholdCustody(threeSigners, 1);

        uint64 newThreshold = 2; // increase
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);

        custody.setThreshold(newThreshold, MAX_DEADLINE, sigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signature_decrease() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 1; // decrease
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_2() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signSetThreshold(signer3Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer3, sig3);

        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_2_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signSetThreshold(signer3Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer3, sig3);

        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_revert_emptySignatures() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 1;
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, _emptySigs());
    }

    function test_revert_signatureFromNotSigner() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        // signature from NOT signer
        bytes memory notSignerSig = _signSetThreshold(notSignerPk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, notSigner, notSignerSig);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSingleSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, sig2);
    }

    function test_revert_duplicatedSignature_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_duplicatedSignature_thresholdReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig3(signer1, sig1, signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_incorrectSignature() public {
        custody = new ThresholdCustody(threeSigners, 1);

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        // corrupt the signature by changing one byte
        sig1[10] = bytes1(sig1[10] ^ 0x01);
        bytes memory encodedSigs = _encodeMultiSig(signer1, sig1);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_zeroNewThreshold() public {
        custody = new ThresholdCustody(oneSigner, 1);

        uint64 newThreshold = 0;
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(MultiSignerERC7913.MultiSignerERC7913ZeroThreshold.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_unreachableNewThreshold() public {
        custody = new ThresholdCustody(oneSigner, 1);

        uint64 newThreshold = 3; // cannot be reached with 1 signer
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 1, 3)
        );
        custody.setThreshold(newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_deadlinePassed() public {
        custody = new ThresholdCustody(oneSigner, 1);

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleSetThreshold(signer1Pk, newThreshold, nonce, block.timestamp - 1);

        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.setThreshold(newThreshold, block.timestamp - 1, sigs);
    }

    function test_revert_outdatedNonce() public {
        custody = new ThresholdCustody(threeSigners, 1);

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signSingleSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);

        // First call succeeds
        custody.setThreshold(newThreshold, MAX_DEADLINE, sig1);

        // Second call with same nonce should revert
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, sig1);
    }

    function test_revert_futureNonce() public {
        custody = new ThresholdCustody(oneSigner, 1);

        uint64 newThreshold = 2; // change
        uint256 futureNonce = custody.signerNonce() + 42;
        bytes memory sigs = _signSingleSetThreshold(signer1Pk, newThreshold, futureNonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.setThreshold(newThreshold, MAX_DEADLINE, sigs);
    }
}

// =========================================================================
// addSigners
// =========================================================================
contract ThresholdCustodyTest_AddSigners is ThresholdCustodyTest_Base {
    function test_success_1_of_1_signature() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        _checkStats(2, newThreshold, ++nonce);
    }

    function test_success_onlyAddSigners() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](2);
        newSigners[0] = signer2;
        newSigners[1] = signer3;

        uint64 newThreshold = 1; // no change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        _checkStats(3, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_2() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signAddSigners(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signAddSigners(signer2Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        assertTrue(custody.isSigner(signer4));
        _checkStats(4, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signAddSigners(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signAddSigners(signer3Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer3, sig3);

        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        assertTrue(custody.isSigner(signer4));
        _checkStats(4, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_2_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signAddSigners(signer2Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signAddSigners(signer3Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer3, sig3);

        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertTrue(custody.isSigner(signer3));
        assertTrue(custody.isSigner(signer4));
        _checkStats(4, newThreshold, ++nonce);
    }

    function test_revert_emptySignatures() public {
        custody = new ThresholdCustody(oneSigner, 1);
        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3;
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, _emptySigs());
    }

    function test_revert_signatureFromNotSigner() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signAddSigners(signer2Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        // signature from NOT signer
        bytes memory notSignerSig = _signAddSigners(notSignerPk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, notSigner, notSignerSig);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sig2);
    }

    function test_revert_duplicatedSignature_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signAddSigners(signer2Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_duplicatedSignature_thresholdReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer4;

        uint64 newThreshold = 3; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signAddSigners(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signAddSigners(signer2Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig3(signer1, sig1, signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_incorrectSignature() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signAddSigners(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);
        // corrupt the signature by changing one byte
        sig1[10] = bytes1(sig1[10] ^ 0x01);
        bytes memory encodedSigs = _encodeMultiSig(signer1, sig1);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_emptyNewArray() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](0);

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.EmptySignersArray.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_newArrayIncludesExistingSigner() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer1; // already existing signer

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913AlreadyExists.selector, signer1.toBytes())
        );
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_newArrayIncludesDuplicatedSigner() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](2);
        newSigners[0] = signer2;
        newSigners[1] = signer2; // duplicated in the new array

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913AlreadyExists.selector, signer2.toBytes())
        );
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_zeroNewThreshold() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 0;
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(MultiSignerERC7913.MultiSignerERC7913ZeroThreshold.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_unreachableNewThreshold() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 3; // cannot be reached with 2 signers
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 2, 3)
        );
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_deadlinePassed() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, block.timestamp - 1);

        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.addSigners(newSigners, newThreshold, block.timestamp - 1, sigs);
    }

    function test_revert_outdatedNonce() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 2; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, nonce, MAX_DEADLINE);

        // First call succeeds
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);

        // Second call with same nonce should revert
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_futureNonce() public {
        custody = new ThresholdCustody(oneSigner, 1);

        address[] memory newSigners = new address[](1);
        newSigners[0] = signer2;

        uint64 newThreshold = 2; // change
        uint256 futureNonce = custody.signerNonce() + 42;
        bytes memory sigs = _signSingleAdd(signer1Pk, newSigners, newThreshold, futureNonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.addSigners(newSigners, newThreshold, MAX_DEADLINE, sigs);
    }
}

// =========================================================================
// removeSigners
// =========================================================================
contract ThresholdCustodyTest_RemoveSigners is ThresholdCustodyTest_Base {
    function test_success_onlyRemoveSigners() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 2; // do NOT change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertFalse(custody.isSigner(signer3));
        _checkStats(2, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_2() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertFalse(custody.isSigner(signer3));
        _checkStats(2, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_1_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signRemoveSigners(signer3Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer3, sig3);

        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertFalse(custody.isSigner(signer3));
        _checkStats(2, newThreshold, ++nonce);
    }

    function test_success_2_of_3_signatures_2_and_3() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig3 = _signRemoveSigners(signer3Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer3, sig3);

        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);

        assertTrue(custody.isSigner(signer1));
        assertTrue(custody.isSigner(signer2));
        assertFalse(custody.isSigner(signer3));
        _checkStats(2, newThreshold, ++nonce);
    }

    function test_revert_emptySignatures() public {
        custody = new ThresholdCustody(twoSigners, 1);
        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 3;
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, _emptySigs());
    }

    function test_revert_signatureFromNotSigner() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        // signature from NOT signer
        bytes memory notSignerSig = _signRemoveSigners(notSignerPk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, notSigner, notSignerSig);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signSingleRemove(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, sig2);
    }

    function test_revert_duplicatedSignature_thresholdNotReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig2(signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_duplicatedSignature_thresholdReached() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        // duplicate signature from signer2
        bytes memory encodedSigs = _encodeMultiSig3(signer1, sig1, signer2, sig2, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_incorrectSignature() public {
        custody = new ThresholdCustody(twoSigners, 1);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();

        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        // corrupt the signature by changing one byte
        sig1[10] = bytes1(sig1[10] ^ 0x01);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_emptyToRemoveArray() public {
        custody = new ThresholdCustody(twoSigners, 2);

        address[] memory signersToRemove = new address[](0);

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(ThresholdCustody.EmptySignersArray.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_removeArrayDoesNotIncludeExistingSigner() public {
        custody = new ThresholdCustody(twoSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3; // not a signer

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913NonexistentSigner.selector, signer3.toBytes())
        );
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_removeArrayIncludesDuplicateSigner() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](2);
        signersToRemove[0] = signer3;
        signersToRemove[1] = signer3; // duplicate

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913NonexistentSigner.selector, signer3.toBytes())
        );
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_removingAllSigners() public {
        custody = new ThresholdCustody(twoSigners, 2);

        address[] memory signersToRemove = new address[](2);
        signersToRemove[0] = signer1;
        signersToRemove[1] = signer2;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 0, 2)
        );
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_zeroNewThreshold() public {
        custody = new ThresholdCustody(twoSigners, 1);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 0;
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleRemove(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(MultiSignerERC7913.MultiSignerERC7913ZeroThreshold.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_unreachableNewThreshold() public {
        custody = new ThresholdCustody(twoSigners, 1);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 2; // cannot be reached with 1 signer
        uint256 nonce = custody.signerNonce();
        bytes memory sigs = _signSingleRemove(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);

        vm.expectRevert(
            abi.encodeWithSelector(MultiSignerERC7913.MultiSignerERC7913UnreachableThreshold.selector, 1, 2)
        );
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, sigs);
    }

    function test_revert_deadlinePassed() public {
        custody = new ThresholdCustody(twoSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.removeSigners(signersToRemove, newThreshold, block.timestamp - 1, encodedSigs);
    }

    function test_revert_outdatedNonce() public {
        custody = new ThresholdCustody(threeSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer3;

        uint64 newThreshold = 1; // change
        uint256 nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        // First call succeeds
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);

        // Second call with same nonce should revert
        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }

    function test_revert_futureNonce() public {
        custody = new ThresholdCustody(twoSigners, 2);

        address[] memory signersToRemove = new address[](1);
        signersToRemove[0] = signer2;

        uint64 newThreshold = 1; // change
        uint256 futureNonce = custody.signerNonce() + 42;
        bytes memory sig1 = _signRemoveSigners(signer1Pk, signersToRemove, newThreshold, futureNonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, signersToRemove, newThreshold, futureNonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.expectRevert(ThresholdCustody.InvalidSignature.selector);
        custody.removeSigners(signersToRemove, newThreshold, MAX_DEADLINE, encodedSigs);
    }
}

// =========================================================================
// Deposit
// =========================================================================
contract ThresholdCustodyTest_Deposit is ThresholdCustodyTest_Base {
    uint256 startEthBalance = 1 ether;
    uint256 startErc20Balance = 1e18;

    function setUp() public override {
        super.setUp();

        custody = new ThresholdCustody(oneSigner, 1);

        vm.deal(address(user), startEthBalance);
        token.mint(address(user), startErc20Balance);

        vm.prank(user);
        token.approve(address(custody), type(uint256).max);
    }

    function test_success_eth() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        custody.deposit{value: depositAmount}(address(0), depositAmount);
        assertEq(address(custody).balance, depositAmount);
        assertEq(user.balance, startEthBalance - depositAmount);
    }

    function test_success_erc20() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        custody.deposit(address(token), depositAmount);
        assertEq(token.balanceOf(address(custody)), depositAmount);
        assertEq(token.balanceOf(user), startErc20Balance - depositAmount);
    }

    function test_eth_emitsEvent() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        vm.expectEmit(true, true, false, true);
        emit IDeposit.Deposited(user, address(0), depositAmount);
        custody.deposit{value: depositAmount}(address(0), depositAmount);
    }

    function test_erc20_emitsEvent() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        vm.expectEmit(true, true, false, true);
        emit IDeposit.Deposited(user, address(token), depositAmount);
        custody.deposit(address(token), depositAmount);
    }

    function test_revert_eth_amountZero() public {
        vm.prank(user);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.deposit(address(0), 0);
    }

    function test_revert_eth_msgValueZero() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit(address(0), depositAmount);
    }

    function test_revert_eth_amountMsgValueMismatch() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: depositAmount / 2}(address(0), depositAmount);
    }

    function test_revert_erc20_amountZero() public {
        vm.prank(user);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.deposit(address(token), 0);
    }

    function test_revert_erc20_nonZeroMsgValue() public {
        uint256 depositAmount = 1e17;

        vm.prank(user);
        vm.expectRevert(IDeposit.InvalidMsgValue.selector);
        custody.deposit{value: 1 ether}(address(token), depositAmount);
    }
}

// =========================================================================
// startWithdraw
// =========================================================================
contract ThresholdCustodyTest_StartWithdraw is ThresholdCustodyTest_Base {
    function setUp() public override {
        super.setUp();

        custody = new ThresholdCustody(oneSigner, 1);

        vm.deal(address(custody), 1 ether);
    }

    function test_success() public {
        address token = address(0);
        uint256 amount = 1 ether;
        uint256 nonce = 1;
        uint256 timestamp = block.timestamp;

        vm.prank(signer1);
        bytes32 id = custody.startWithdraw(user, token, amount, nonce);

        bytes32 expectedId = Utils.getWithdrawalId(address(custody), user, token, amount, nonce);
        assertEq(id, expectedId);
        _validateWithdrawalData(id, user, token, amount, false, 1, uint64(timestamp));
    }

    function test_success_emitsEvent() public {
        address token = address(0);
        uint256 amount = 1 ether;
        uint256 nonce = 1;

        vm.prank(signer1);
        vm.expectEmit(true, true, true, true);
        bytes32 expectedId = Utils.getWithdrawalId(address(custody), user, token, amount, nonce);
        emit IWithdraw.WithdrawStarted(expectedId, user, token, amount, nonce);
        custody.startWithdraw(user, token, amount, nonce);
    }

    function test_success_sameParamsDifferentNonce() public {
        address token = address(0);
        uint256 amount = 1 ether;
        uint256 nonce = 1;
        uint256 timestamp = block.timestamp;

        vm.startPrank(signer1);
        bytes32 id1 = custody.startWithdraw(user, token, amount, nonce);
        bytes32 id2 = custody.startWithdraw(user, token, amount, nonce + 1);
        vm.stopPrank();

        assertTrue(id1 != id2);
        uint64 expectedThreshold = 1;
        _validateWithdrawalData(id1, user, token, amount, false, expectedThreshold, uint64(timestamp));
        _validateWithdrawalData(id2, user, token, amount, false, expectedThreshold, uint64(timestamp));
    }

    function test_revert_callerNotSigner() public {
        vm.prank(user);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
    }

    function test_revert_userAddressZero() public {
        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.InvalidUser.selector);
        custody.startWithdraw(address(0), address(0), 1 ether, 1);
    }

    function test_revert_zeroAmount() public {
        vm.prank(signer1);
        vm.expectRevert(IDeposit.ZeroAmount.selector);
        custody.startWithdraw(user, address(0), 0, 1);
    }

    function test_revert_duplicateNonce() public {
        vm.startPrank(signer1);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyExists.selector);
        custody.startWithdraw(user, address(0), 1 ether, 1);
        vm.stopPrank();
    }
}

// =========================================================================
// finalizeWithdraw
// =========================================================================
contract ThresholdCustodyTest_FinalizeWithdraw is ThresholdCustodyTest_Base {
    uint256 custodyNativeBalance = 1 ether;
    uint256 custodyErc20Balance = 10e18;
    uint256 withdrawalAmount = 1e17;
    uint256 nonce = 1;

    function _setUpTest(address[] memory signers, uint64 threshold, address withdrawalToken, uint256 withdrawalAmount_)
        public
        returns (bytes32)
    {
        custody = new ThresholdCustody(signers, threshold);
        vm.deal(address(custody), custodyNativeBalance);
        token.mint(address(custody), custodyErc20Balance);

        vm.prank(signers[0]);
        bytes32 id = custody.startWithdraw(user, withdrawalToken, withdrawalAmount_, nonce);
        return id;
    }

    function _validateInitialBalances(address user, address token) public view {
        if (token == address(0)) {
            assertEq(address(custody).balance, custodyNativeBalance);
            assertEq(user.balance, 0);
        } else {
            assertEq(ERC20(token).balanceOf(address(custody)), custodyErc20Balance);
            assertEq(ERC20(token).balanceOf(user), 0);
        }
    }

    function _validateBalanceWithdrawn(address user, address token, uint256 withdrawnAmount) public view {
        if (token == address(0)) {
            assertEq(address(custody).balance, custodyNativeBalance - withdrawnAmount);
            assertEq(user.balance, withdrawnAmount);
        } else {
            assertEq(ERC20(token).balanceOf(address(custody)), custodyErc20Balance - withdrawnAmount);
            assertEq(ERC20(token).balanceOf(user), withdrawnAmount);
        }
    }

    function test_success_1of1_eth() public {
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 1;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(oneSigner, expectedThreshold, withdrawalToken, withdrawalAmount);

        _validateInitialBalances(user, withdrawalToken);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        // user, token, amount are cleared; finalized=true, threshold and createdAt remain
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, startedAt);
        _validateBalanceWithdrawn(user, withdrawalToken, withdrawalAmount);
    }

    function test_success_1of1_erc20() public {
        address withdrawalToken = address(token);
        uint64 expectedThreshold = 1;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(oneSigner, expectedThreshold, withdrawalToken, withdrawalAmount);

        _validateInitialBalances(user, withdrawalToken);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        // user, token, amount are cleared; finalized=true, threshold and createdAt remain
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, startedAt);
        _validateBalanceWithdrawn(user, withdrawalToken, withdrawalAmount);
    }

    function test_success_2of2_eth() public {
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 2;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(twoSigners, expectedThreshold, withdrawalToken, withdrawalAmount);

        _validateInitialBalances(user, withdrawalToken);

        vm.warp(block.timestamp + 1 minutes);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, false, expectedThreshold, startedAt);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // user, token, amount are cleared; finalized=true, threshold and createdAt remain
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, startedAt);
        _validateBalanceWithdrawn(user, withdrawalToken, withdrawalAmount);
    }

    function test_success_3of5_eth() public {
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 3;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(fiveSigners, expectedThreshold, withdrawalToken, withdrawalAmount);

        _validateInitialBalances(user, withdrawalToken);
        vm.warp(block.timestamp + 1 minutes);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, false, expectedThreshold, startedAt);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        _validateInitialBalances(user, withdrawalToken);
        vm.warp(block.timestamp + 1 minutes);

        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, false, expectedThreshold, startedAt);

        vm.prank(signer3);
        custody.finalizeWithdraw(id);

        // user, token, amount are cleared; finalized=true, threshold and createdAt remain
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, startedAt);
        _validateBalanceWithdrawn(user, withdrawalToken, withdrawalAmount);
    }

    function test_success_usingSnapshotThreshold_eth() public {
        // Setup: 2 signers, threshold=2
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 2;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(twoSigners, expectedThreshold, withdrawalToken, withdrawalAmount);

        // Lower threshold to 1 AFTER withdrawal was created
        uint64 newThreshold = 1;
        nonce = custody.signerNonce();
        bytes memory sig1 = _signSetThreshold(signer1Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signSetThreshold(signer2Pk, newThreshold, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

        vm.prank(signer1);
        custody.setThreshold(newThreshold, MAX_DEADLINE, encodedSigs);
        assertEq(custody.threshold(), newThreshold);

        vm.warp(block.timestamp + 1 minutes);

        // 1 approval should NOT suffice (snapshot quorum was 2)
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, false, expectedThreshold, startedAt);

        // 2nd approval should finalize the withdrawal
        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // user, token, amount are cleared; finalized=true, threshold and createdAt remain
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, startedAt);
        _validateBalanceWithdrawn(user, withdrawalToken, withdrawalAmount);
    }

    function test_success_eventsEmitted_2of2_eth() public {
        address token = address(0);
        bytes32 id = _setUpTest(twoSigners, 2, token, withdrawalAmount);

        vm.expectEmit(true, true, true, true);
        emit ThresholdCustody.WithdrawalApproved(id, signer1, 1);
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.expectEmit(true, true, true, true);
        emit ThresholdCustody.WithdrawalApproved(id, signer2, 2);
        vm.expectEmit(true, true, true, true);
        emit IWithdraw.WithdrawFinalized(id, true);
        vm.prank(signer2);
        custody.finalizeWithdraw(id);
    }

    function test_revert_notSigner() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.prank(notSigner);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_duplicateApproval() public {
        bytes32 id = _setUpTest(twoSigners, 2, address(0), withdrawalAmount);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.SignerAlreadyApproved.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_nonExistentWithdrawal() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);
        // corrupt the ID to ensure it doesn't match any real withdrawal
        id = bytes32(uint256(id) + 42);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalNotFound.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_alreadyFinalized() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_afterReject() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        // Warp past expiry and reject
        vm.warp(block.timestamp + OPERATION_EXPIRY + 1);
        vm.prank(signer1);
        custody.rejectWithdraw(id);

        // Try to finalize after rejection - should revert
        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.finalizeWithdraw(id);
    }

    function test_success_exactlyAtExpiryBoundary() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);
        uint64 createdAt = uint64(block.timestamp);

        // Warp to exactly createdAt + OPERATION_EXPIRY (should succeed)
        vm.warp(createdAt + OPERATION_EXPIRY);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        // Verify withdrawal was executed successfully
        _validateWithdrawalData(id, address(0), address(0), 0, true, 1, createdAt);
        _validateBalanceWithdrawn(user, address(0), withdrawalAmount);
    }

    function test_revert_oneSecondAfterExpiry() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);
        uint64 createdAt = uint64(block.timestamp);

        // Warp to createdAt + OPERATION_EXPIRY + 1 (should revert)
        vm.warp(createdAt + OPERATION_EXPIRY + 1);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.DeadlineExpired.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_insufficientEth() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), custodyNativeBalance + 1);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_revert_insufficientErc20() public {
        vm.prank(signer1);
        bytes32 id = _setUpTest(oneSigner, 1, address(token), custodyErc20Balance + 1);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        custody.finalizeWithdraw(id);
    }

    function test_multipleConcurrentWithdrawals() public {
        custody = new ThresholdCustody(oneSigner, 1);
        vm.deal(address(custody), custodyNativeBalance);

        address withdrawalToken = address(0);
        uint64 createdAt = uint64(block.timestamp);

        vm.startPrank(signer1);
        bytes32 id1 = custody.startWithdraw(user, withdrawalToken, withdrawalAmount, 1);
        bytes32 id3 = custody.startWithdraw(user, withdrawalToken, withdrawalAmount, 3);
        vm.stopPrank();

        vm.prank(signer1);
        custody.finalizeWithdraw(id1);

        _validateWithdrawalData(id1, address(0), address(0), 0, true, 1, createdAt);
        _validateBalanceWithdrawn(user, address(0), withdrawalAmount);

        // id2 left to expire

        vm.prank(signer1);
        custody.finalizeWithdraw(id3);

        _validateWithdrawalData(id3, address(0), address(0), 0, true, 1, createdAt);
        _validateBalanceWithdrawn(user, address(0), 2 * withdrawalAmount);
    }

    function test_removedSignerApprovalIgnored() public {
        bytes32 id = _setUpTest(threeSigners, 2, address(0), withdrawalAmount);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // Remove signer2 (need 2 sigs since threshold=2)
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;

        nonce = custody.signerNonce();
        bytes memory sig1 = _signRemoveSigners(signer1Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory sig2 = _signRemoveSigners(signer2Pk, toRemove, 2, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig2(signer1, sig1, signer2, sig2);

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
}

// =========================================================================
// rejectWithdraw
// =========================================================================
contract ThresholdCustodyTest_RejectWithdraw is ThresholdCustodyTest_Base {
    uint256 custodyNativeBalance = 1 ether;
    uint256 custodyErc20Balance = 10e18;
    uint256 withdrawalAmount = 1e17;
    uint256 nonce = 1;

    function _setUpTest(address[] memory signers, uint64 threshold, address withdrawalToken, uint256 withdrawalAmount_)
        public
        returns (bytes32)
    {
        custody = new ThresholdCustody(signers, threshold);
        vm.deal(address(custody), custodyNativeBalance);
        token.mint(address(custody), custodyErc20Balance);

        vm.prank(signers[0]);
        bytes32 id = custody.startWithdraw(user, withdrawalToken, withdrawalAmount_, nonce);
        return id;
    }

    function _validateInitialBalances(address user, address token) public view {
        if (token == address(0)) {
            assertEq(address(custody).balance, custodyNativeBalance);
            assertEq(user.balance, 0);
        } else {
            assertEq(ERC20(token).balanceOf(address(custody)), custodyErc20Balance);
            assertEq(ERC20(token).balanceOf(user), 0);
        }
    }

    function test_success() public {
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 1;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(oneSigner, expectedThreshold, withdrawalToken, withdrawalAmount);

        vm.warp(startedAt + OPERATION_EXPIRY + 1);

        vm.prank(signer1);
        custody.rejectWithdraw(id);

        // rejectWithdraw only sets finalized=true, does NOT clear user/token/amount
        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, true, expectedThreshold, startedAt);
    }

    function test_success_emitsEvent() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.warp(block.timestamp + OPERATION_EXPIRY + 1);

        vm.expectEmit(true, false, false, true);
        emit IWithdraw.WithdrawFinalized(id, false);
        vm.prank(signer1);
        custody.rejectWithdraw(id);
    }

    function test_success_afterPartiallyApproved() public {
        address withdrawalToken = address(0);
        uint64 expectedThreshold = 2;
        uint64 startedAt = uint64(block.timestamp);
        bytes32 id = _setUpTest(twoSigners, expectedThreshold, withdrawalToken, withdrawalAmount);

        // Get 1 approval (not enough to finalize)
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.warp(startedAt + OPERATION_EXPIRY + 1);

        (,,, bool finalized,,) = custody.withdrawals(id);
        assertFalse(finalized);

        vm.prank(signer1);
        custody.rejectWithdraw(id);

        // rejectWithdraw only sets finalized=true, does NOT clear user/token/amount
        _validateInitialBalances(user, withdrawalToken);
        _validateWithdrawalData(id, user, withdrawalToken, withdrawalAmount, true, expectedThreshold, startedAt);
    }

    function test_revert_notExpired() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        // do NOT warp past expiry

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.WithdrawalNotExpired.selector);
        custody.rejectWithdraw(id);
    }

    function test_revert_notExpired_lastSecond() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.warp(block.timestamp + OPERATION_EXPIRY);

        vm.prank(signer1);
        vm.expectRevert(ThresholdCustody.WithdrawalNotExpired.selector);
        custody.rejectWithdraw(id);
    }

    function test_revert_nonExistent() public {
        custody = new ThresholdCustody(oneSigner, 1);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalNotFound.selector);
        custody.rejectWithdraw(bytes32(uint256(999)));
    }

    function test_revert_alreadyFinalized() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.warp(block.timestamp + OPERATION_EXPIRY);

        vm.prank(signer1);
        vm.expectRevert(IWithdraw.WithdrawalAlreadyFinalized.selector);
        custody.rejectWithdraw(id);
    }

    function test_revert_notSigner() public {
        bytes32 id = _setUpTest(oneSigner, 1, address(0), withdrawalAmount);

        vm.warp(block.timestamp + OPERATION_EXPIRY);

        vm.prank(user);
        vm.expectRevert(ThresholdCustody.NotSigner.selector);
        custody.rejectWithdraw(id);
    }

    function test_success_rejectWithPartialApprovals() public {
        bytes32 id = _setUpTest(threeSigners, 3, address(0), withdrawalAmount);

        // Get 2 approvals (not enough to finalize)
        vm.prank(signer1);
        custody.finalizeWithdraw(id);

        vm.prank(signer2);
        custody.finalizeWithdraw(id);

        // Verify not finalized yet
        (,,, bool finalized1,,) = custody.withdrawals(id);
        assertFalse(finalized1);

        // Warp past expiry
        vm.warp(block.timestamp + OPERATION_EXPIRY + 1);

        // Reject the withdrawal
        vm.prank(signer1);
        custody.rejectWithdraw(id);

        // Verify finalized with success=false
        (,,, bool finalized2,,) = custody.withdrawals(id);
        assertTrue(finalized2);
    }
}

// =========================================================================
// _executeWithdrawal (internal function via harness)
// =========================================================================
contract ThresholdCustodyTest_ExecuteWithdrawal is ThresholdCustodyTest_Base {
    TestThresholdCustody public testCustody;

    uint256 custodyNativeBalance = 1 ether;
    uint256 custodyErc20Balance = 10e18;
    uint256 withdrawalAmount = 1e17;

    function setUp() public override {
        super.setUp();
        testCustody = new TestThresholdCustody(oneSigner, 1);
        vm.deal(address(testCustody), custodyNativeBalance);
        token.mint(address(testCustody), custodyErc20Balance);
    }

    function _createWithdrawal(address withdrawalToken, uint256 amount) internal returns (bytes32) {
        vm.prank(signer1);
        bytes32 id = testCustody.startWithdraw(user, withdrawalToken, amount, 1);
        return id;
    }

    function _validateWithdrawalData(
        bytes32 id,
        address expectedUser,
        address expectedToken,
        uint256 expectedAmount,
        bool expectedFinalized,
        uint64 expectedThreshold,
        uint64 expectedCreatedAt
    ) internal view override {
        (
            address storedUser,
            address storedToken,
            uint256 storedAmount,
            bool storedFinalized,
            uint64 storedCreatedAt,
            uint64 storedThreshold
        ) = testCustody.withdrawals(id);
        assertEq(storedUser, expectedUser);
        assertEq(storedToken, expectedToken);
        assertEq(storedAmount, expectedAmount);
        assertEq(storedFinalized, expectedFinalized);
        assertEq(storedThreshold, expectedThreshold);
        assertEq(storedCreatedAt, expectedCreatedAt);
    }

    function test_success_eth() public {
        bytes32 id = _createWithdrawal(address(0), withdrawalAmount);
        uint64 expectedThreshold = 1;
        uint64 createdAt = uint64(block.timestamp);

        testCustody.exposed_executeWithdrawal(id);

        // Verify finalized is set to true, user/token/amount are cleared
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, createdAt);

        assertEq(address(testCustody).balance, custodyNativeBalance - withdrawalAmount);
        assertEq(user.balance, withdrawalAmount);
    }

    function test_success_erc20() public {
        bytes32 id = _createWithdrawal(address(token), withdrawalAmount);
        uint64 expectedThreshold = 1;
        uint64 createdAt = uint64(block.timestamp);

        testCustody.exposed_executeWithdrawal(id);

        // Verify finalized is set to true, user/token/amount are cleared
        _validateWithdrawalData(id, address(0), address(0), 0, true, expectedThreshold, createdAt);

        assertEq(token.balanceOf(address(testCustody)), custodyErc20Balance - withdrawalAmount);
        assertEq(token.balanceOf(user), withdrawalAmount);
    }

    function test_revert_eth_insufficientLiquidity() public {
        bytes32 id = _createWithdrawal(address(0), custodyNativeBalance + 1);

        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        testCustody.exposed_executeWithdrawal(id);
    }

    function test_revert_erc20_insufficientLiquidity() public {
        bytes32 id = _createWithdrawal(address(token), custodyErc20Balance + 1);

        vm.expectRevert(IWithdraw.InsufficientLiquidity.selector);
        testCustody.exposed_executeWithdrawal(id);
    }
}

// =========================================================================
// _countValidApprovals (internal function via harness)
// =========================================================================
contract ThresholdCustodyTest_CountValidApprovals is ThresholdCustodyTest_Base {
    using {Utils.hashArrayed} for address[];

    TestThresholdCustody public testCustody;
    bytes32 withdrawalId;

    function setUp() public override {
        super.setUp();
    }

    function _setupCustody(address[] memory signers, uint64 threshold) internal {
        testCustody = new TestThresholdCustody(signers, threshold);
        vm.deal(address(testCustody), 1 ether);

        // Create a withdrawal
        vm.prank(signers[0]);
        withdrawalId = testCustody.startWithdraw(user, address(0), 1e17, 1);
    }

    function _testCustodyDomainSeparator() internal view returns (bytes32) {
        return _domainSeparator(address(testCustody));
    }

    function _signRemoveSignersForTestCustody(
        uint256 pk,
        address[] memory signersToRemove,
        uint256 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(REMOVE_SIGNERS_TYPEHASH, signersToRemove.hashArrayed(), newThreshold, nonce, deadline)
        );
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", _testCustodyDomainSeparator(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    function test_zero_forNoApprovals() public {
        _setupCustody(threeSigners, 2);

        uint256 count = testCustody.exposed_countValidApprovals(withdrawalId);
        assertEq(count, 0);
    }

    function test_1_for1Approval() public {
        _setupCustody(threeSigners, 2);

        testCustody.workaround_setWithdrawalApproval(withdrawalId, signer1, true);

        uint256 count = testCustody.exposed_countValidApprovals(withdrawalId);
        assertEq(count, 1);
    }

    function test_approvalReduces_afterSignerRemoval() public {
        _setupCustody(threeSigners, 1);

        testCustody.workaround_setWithdrawalApproval(withdrawalId, signer1, true);
        testCustody.workaround_setWithdrawalApproval(withdrawalId, signer2, true);

        uint256 countBefore = testCustody.exposed_countValidApprovals(withdrawalId);
        assertEq(countBefore, 2, "Should have 2 approvals before removal");

        // Remove signer2 using threshold=1
        address[] memory toRemove = new address[](1);
        toRemove[0] = signer2;

        uint256 nonce = testCustody.signerNonce();
        bytes memory sig = _signRemoveSignersForTestCustody(signer1Pk, toRemove, 1, nonce, MAX_DEADLINE);
        bytes memory encodedSigs = _encodeMultiSig(signer1, sig);

        testCustody.removeSigners(toRemove, 1, MAX_DEADLINE, encodedSigs);

        // Count should now be 1 (only signer1's approval counts)
        uint256 countAfter = testCustody.exposed_countValidApprovals(withdrawalId);
        assertEq(countAfter, 1, "Should have 1 approval after removal");
    }

    function _signerPk(uint256 i) internal pure returns (uint256) {
        return uint256(keccak256(abi.encode("fuzz_signer", i))) % (type(uint256).max - 1) + 1;
    }

    // restrict to uint8 to avoid memory issues with large arrays in fuzzing
    function testFuzz_countValidApprovals(uint8 x, uint8 y, uint8 z) public {
        // Constrain: x signers total, y approvals, z signers to remove
        vm.assume(x >= 1);
        vm.assume(y <= x);
        vm.assume(z <= y);
        vm.assume(z < x); // Cannot remove all signers

        // Create array of x signers
        address[] memory signers = new address[](x);
        for (uint256 i = 0; i < x; i++) {
            signers[i] = vm.addr(_signerPk(i));
        }

        // Setup custody with threshold=1 for easy operations
        testCustody = new TestThresholdCustody(signers, 1);

        withdrawalId = bytes32("deadbeef");

        // Approve y signers
        for (uint256 i = 0; i < y; i++) {
            testCustody.workaround_setWithdrawalApproval(withdrawalId, signers[i], true);
        }

        uint256 countBefore = testCustody.exposed_countValidApprovals(withdrawalId);
        assertEq(countBefore, y, "Should have y approvals before removal");

        // Remove z signers (from the ones that approved)
        if (z > 0) {
            address[] memory toRemove = new address[](z);
            for (uint256 i = 0; i < z; i++) {
                toRemove[i] = signers[i];
            }

            uint256 nonce = testCustody.signerNonce();
            uint256 signerPk = _signerPk(0);
            bytes memory sig = _signRemoveSignersForTestCustody(signerPk, toRemove, 1, nonce, MAX_DEADLINE);
            bytes memory encodedSigs = _encodeMultiSig(signers[0], sig);

            testCustody.removeSigners(toRemove, 1, MAX_DEADLINE, encodedSigs);

            uint256 countAfter = testCustody.exposed_countValidApprovals(withdrawalId);
            assertEq(countAfter, y - z, "Should have (y - z) approvals after removal");
        }
    }
}
