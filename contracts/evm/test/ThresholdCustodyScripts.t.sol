// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// NOTE: These tests MUST be run with --threads 1 flag to ensure proper test isolation
// due to environment variable usage. Run with:
//   forge test --match-contract ThresholdCustodyScriptsTest --threads 1

import {Test, console} from "forge-std/Test.sol";
import {
    ThresholdCustody,
    NAME,
    VERSION,
    ADD_SIGNERS_TYPEHASH,
    REMOVE_SIGNERS_TYPEHASH,
    SET_THRESHOLD_TYPEHASH
} from "../src/ThresholdCustody.sol";
import {Utils} from "../src/Utils.sol";
import {ThresholdCustodyScriptUtils} from "../script/ThresholdCustody/Utils.sol";

import {SignAddSigners} from "../script/ThresholdCustody/SignAddSigners.s.sol";
import {AddSigners} from "../script/ThresholdCustody/AddSigners.s.sol";
import {SignRemoveSigners} from "../script/ThresholdCustody/SignRemoveSigners.s.sol";
import {RemoveSigners} from "../script/ThresholdCustody/RemoveSigners.s.sol";
import {SignSetThreshold} from "../script/ThresholdCustody/SignSetThreshold.s.sol";
import {SetThreshold} from "../script/ThresholdCustody/SetThreshold.s.sol";

contract ThresholdCustodyScriptsTest is Test {
    using {Utils.toAddressBytesArray} for address[];

    address internal signer1;
    uint256 internal signer1Pk;
    address internal signer2;
    uint256 internal signer2Pk;
    address internal signer3;
    uint256 internal signer3Pk;
    address internal signer4;
    uint256 internal signer4Pk;
    address internal txSender;
    uint256 internal txSenderPk;

    address[] twoSigners;
    address[] threeSigners;
    address[] fourSigners;

    SignAddSigners internal signAddSignersScript;
    AddSigners internal addSignersScript;
    SignRemoveSigners internal signRemoveSignersScript;
    RemoveSigners internal removeSignersScript;
    SignSetThreshold internal signSetThresholdScript;
    SetThreshold internal setThresholdScript;

    function setUp() public {
        (signer1, signer1Pk) = makeAddrAndKey("signer1");
        (signer2, signer2Pk) = makeAddrAndKey("signer2");
        (signer3, signer3Pk) = makeAddrAndKey("signer3");
        (signer4, signer4Pk) = makeAddrAndKey("signer4");
        (txSender, txSenderPk) = makeAddrAndKey("txSender");

        twoSigners = new address[](2);
        twoSigners[0] = signer1;
        twoSigners[1] = signer2;

        threeSigners = new address[](3);
        threeSigners[0] = signer1;
        threeSigners[1] = signer2;
        threeSigners[2] = signer3;

        fourSigners = new address[](4);
        fourSigners[0] = signer1;
        fourSigners[1] = signer2;
        fourSigners[2] = signer3;
        fourSigners[3] = signer4;

        vm.deal(txSender, 100 ether);

        signAddSignersScript = new SignAddSigners();
        addSignersScript = new AddSigners();
        signRemoveSignersScript = new SignRemoveSigners();
        removeSignersScript = new RemoveSigners();
        signSetThresholdScript = new SignSetThreshold();
        setThresholdScript = new SetThreshold();
    }

    function _sign(uint256 privateKey, bytes32 digest) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    function test_AddSigners_WithNewThreshold() public {
        ThresholdCustody custody = new ThresholdCustody(twoSigners, 2);

        // Add signer3 and signer4 as new signers
        string memory newSigners = string.concat(vm.toString(signer3), ",", vm.toString(signer4));
        uint256 deadline = block.timestamp + 1 hours;
        uint64 newThreshold = 4;

        // Step 1: Generate signature from signer1 using SignAddSigners script
        vm.setEnv("SIGNER_PK", vm.toString(signer1Pk));
        vm.setEnv("CONTRACT", vm.toString(address(custody)));
        vm.setEnv("NEW_SIGNERS", newSigners);
        vm.setEnv("NEW_THRESHOLD", vm.toString(newThreshold));
        vm.setEnv("DEADLINE", vm.toString(deadline));

        bytes memory sig1 = signAddSignersScript.run();

        // Step 2: Generate signature from signer2
        vm.setEnv("SIGNER_PK", vm.toString(signer2Pk));
        bytes memory sig2 = signAddSignersScript.run();

        string memory signaturesEnv = string.concat(vm.toString(sig1), ",", vm.toString(sig2));

        vm.setEnv("TX_SENDER_PK", vm.toString(txSenderPk));
        vm.setEnv("SIGNATURES", signaturesEnv);

        // Verify initial state
        assertEq(custody.getSigners(0, type(uint64).max).length, 2, "Initial signer count should be 2");
        assertEq(custody.threshold(), 2, "Initial threshold should be 2");

        // Execute the add signers transaction
        addSignersScript.run();

        // Verify final state
        assertEq(custody.getSigners(0, type(uint64).max).length, 4, "Final signer count should be 4");
        assertEq(custody.threshold(), 4, "Final threshold should be 4");
        assertTrue(custody.isSigner(signer3), "signer3 should be a signer");
        assertTrue(custody.isSigner(signer4), "signer4 should be a signer");
    }

    function test_AddSigners_WithoutNewThreshold() public {
        ThresholdCustody custody = new ThresholdCustody(twoSigners, 2);

        uint256 deadline = block.timestamp + 1 hours;

        // Unset NEW_THRESHOLD to indicate "use existing threshold"
        vm.setEnv("NEW_THRESHOLD", "");
        // Add signer3 and signer4 as new signers
        string memory newSigners = string.concat(vm.toString(signer3), ",", vm.toString(signer4));
        vm.setEnv("SIGNER_PK", vm.toString(signer1Pk));
        vm.setEnv("CONTRACT", vm.toString(address(custody)));
        vm.setEnv("NEW_SIGNERS", newSigners);
        vm.setEnv("DEADLINE", vm.toString(deadline));

        // Step 1: Generate signature from signer1 using SignAddSigners script
        bytes memory sig1 = signAddSignersScript.run();

        // Step 2: Generate signature from signer2
        vm.setEnv("SIGNER_PK", vm.toString(signer2Pk));
        bytes memory sig2 = signAddSignersScript.run();

        string memory signaturesEnv = string.concat(vm.toString(sig1), ",", vm.toString(sig2));

        vm.setEnv("TX_SENDER_PK", vm.toString(txSenderPk));
        vm.setEnv("SIGNATURES", signaturesEnv);

        // Verify initial state
        assertEq(custody.getSigners(0, type(uint64).max).length, 2, "Initial signer count should be 2");
        assertEq(custody.threshold(), 2, "Initial threshold should be 2");

        // Execute the add signers transaction
        addSignersScript.run();

        // Verify final state
        assertEq(custody.getSigners(0, type(uint64).max).length, 4, "Final signer count should be 4");
        assertEq(custody.threshold(), 2, "Threshold should remain 2");
        assertTrue(custody.isSigner(signer3), "signer3 should be a signer");
        assertTrue(custody.isSigner(signer4), "signer4 should be a signer");
    }

    function test_RemoveSigners_WithNewThreshold() public {
        ThresholdCustody custody = new ThresholdCustody(fourSigners, 3);

        uint256 deadline = block.timestamp + 1 hours;
        uint64 newThreshold = 1;

        // Remove signer3 and signer4, leaving signer1 and signer2
        string memory signersToRemove = string.concat(vm.toString(signer3), ",", vm.toString(signer4));
        vm.setEnv("SIGNER_PK", vm.toString(signer1Pk));
        vm.setEnv("CONTRACT", vm.toString(address(custody)));
        vm.setEnv("SIGNERS_TO_REMOVE", signersToRemove);
        vm.setEnv("NEW_THRESHOLD", vm.toString(newThreshold));
        vm.setEnv("DEADLINE", vm.toString(deadline));

        // Step 1: Generate signature from signer1 using SignRemoveSigners script
        bytes memory sig1 = signRemoveSignersScript.run();

        // Step 2: Generate signature from signer2
        vm.setEnv("SIGNER_PK", vm.toString(signer2Pk));
        bytes memory sig2 = signRemoveSignersScript.run();

        // Step 3: Generate signature from signer3
        vm.setEnv("SIGNER_PK", vm.toString(signer3Pk));
        bytes memory sig3 = signRemoveSignersScript.run();

        string memory signaturesEnv = string.concat(vm.toString(sig1), ",", vm.toString(sig2), ",", vm.toString(sig3));

        vm.setEnv("TX_SENDER_PK", vm.toString(txSenderPk));
        vm.setEnv("SIGNATURES", signaturesEnv);

        // Verify initial state
        assertEq(custody.getSigners(0, type(uint64).max).length, 4, "Initial signer count should be 4");
        assertEq(custody.threshold(), 3, "Initial threshold should be 3");
        assertTrue(custody.isSigner(signer3), "signer3 should initially be a signer");
        assertTrue(custody.isSigner(signer4), "signer4 should initially be a signer");

        // Execute the remove signers transaction
        removeSignersScript.run();

        // Verify final state
        assertEq(custody.getSigners(0, type(uint64).max).length, 2, "Final signer count should be 2");
        assertEq(custody.threshold(), 1, "Final threshold should be 1");
        assertFalse(custody.isSigner(signer3), "signer3 should no longer be a signer");
        assertFalse(custody.isSigner(signer4), "signer4 should no longer be a signer");
        assertTrue(custody.isSigner(signer1), "signer1 should still be a signer");
        assertTrue(custody.isSigner(signer2), "signer2 should still be a signer");
    }

    function test_RemoveSigners_WithoutNewThreshold() public {
        ThresholdCustody custody = new ThresholdCustody(fourSigners, 2);

        uint256 deadline = block.timestamp + 1 hours;

        // Unset NEW_THRESHOLD to indicate "use existing threshold"
        vm.setEnv("NEW_THRESHOLD", "");
        // Remove signer3 and signer4, keeping threshold at 2
        string memory signersToRemove = string.concat(vm.toString(signer3), ",", vm.toString(signer4));
        vm.setEnv("SIGNER_PK", vm.toString(signer1Pk));
        vm.setEnv("CONTRACT", vm.toString(address(custody)));
        vm.setEnv("SIGNERS_TO_REMOVE", signersToRemove);
        vm.setEnv("DEADLINE", vm.toString(deadline));

        // Step 1: Generate signature from signer1 using SignRemoveSigners script
        bytes memory sig1 = signRemoveSignersScript.run();

        // Step 2: Generate signature from signer2
        vm.setEnv("SIGNER_PK", vm.toString(signer2Pk));
        bytes memory sig2 = signRemoveSignersScript.run();

        string memory signaturesEnv = string.concat(vm.toString(sig1), ",", vm.toString(sig2));

        vm.setEnv("TX_SENDER_PK", vm.toString(txSenderPk));
        vm.setEnv("SIGNATURES", signaturesEnv);

        // Verify initial state
        assertEq(custody.getSigners(0, type(uint64).max).length, 4, "Initial signer count should be 4");
        assertEq(custody.threshold(), 2, "Initial threshold should be 2");

        // Execute the remove signers transaction
        removeSignersScript.run();

        // Verify final state
        assertEq(custody.getSigners(0, type(uint64).max).length, 2, "Final signer count should be 2");
        assertEq(custody.threshold(), 2, "Threshold should remain 2");
        assertFalse(custody.isSigner(signer3), "signer3 should no longer be a signer");
        assertFalse(custody.isSigner(signer4), "signer4 should no longer be a signer");
    }

    function test_SetThreshold() public {
        ThresholdCustody custody = new ThresholdCustody(threeSigners, 2);

        uint256 deadline = block.timestamp + 1 hours;
        uint64 newThreshold = 3;

        vm.setEnv("SIGNER_PK", vm.toString(signer1Pk));
        vm.setEnv("CONTRACT", vm.toString(address(custody)));
        vm.setEnv("NEW_THRESHOLD", vm.toString(newThreshold));
        vm.setEnv("DEADLINE", vm.toString(deadline));

        // Step 1: Generate signature from signer1 using SignSetThreshold script
        bytes memory sig1 = signSetThresholdScript.run();

        // Step 2: Generate signature from signer2
        vm.setEnv("SIGNER_PK", vm.toString(signer2Pk));
        bytes memory sig2 = signSetThresholdScript.run();

        string memory signaturesEnv = string.concat(vm.toString(sig1), ",", vm.toString(sig2));

        vm.setEnv("TX_SENDER_PK", vm.toString(txSenderPk));
        vm.setEnv("SIGNATURES", signaturesEnv);

        // Verify initial state
        assertEq(custody.threshold(), 2, "Initial threshold should be 2");

        // Execute the set threshold transaction
        setThresholdScript.run();

        // Verify final state
        assertEq(custody.threshold(), 3, "Final threshold should be 3");
        assertEq(custody.getSigners(0, type(uint64).max).length, 3, "Signer count should remain 3");
    }
}
