// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ThresholdCustody} from "../../src/ThresholdCustody.sol";
import {ThresholdCustodyScriptUtils as Utils} from "./Utils.sol";

/// @notice Generate a signature for removing signers from ThresholdCustody contract.
///
/// Environment variables:
///   SIGNER_PK          – Signer's private key
///   CONTRACT           – ThresholdCustody contract address
///   SIGNERS_TO_REMOVE  – Comma-separated addresses to remove
///   NEW_THRESHOLD      – New threshold value (optional, defaults to current threshold)
///   DEADLINE           – Unix timestamp deadline for the operation
///
/// Example:
///   SIGNER_PK=0xabc... CONTRACT=0x... SIGNERS_TO_REMOVE=0xA,0xB NEW_THRESHOLD=2 DEADLINE=1999999999 \
///     forge script script/ThresholdCustody/SignRemoveSigners.s.sol --rpc-url $RPC_URL
contract SignRemoveSigners is Script {
    function run() public view {
        uint256 signerKey = vm.envUint("SIGNER_PK");
        address contractAddr = vm.envAddress("CONTRACT");
        address[] memory signersToRemove = vm.envAddress("SIGNERS_TO_REMOVE", ",");
        uint256 deadline = vm.envUint("DEADLINE");

        ThresholdCustody custody = ThresholdCustody(contractAddr);

        // Get threshold: use NEW_THRESHOLD if provided, otherwise use current threshold
        uint64 newThreshold = Utils.getThreshold(vm, "NEW_THRESHOLD", custody);

        // Calculate resulting signer count
        bytes[] memory currentSigners = custody.getSigners(0, type(uint64).max);
        uint256 resultingSignerCount = currentSigners.length - signersToRemove.length;

        // Validate threshold
        Utils.validateThreshold(newThreshold, resultingSignerCount);

        // Build EIP-712 digest
        bytes32 domainSeparator = Utils.getDomainSeparator(contractAddr);
        bytes32 structHash =
            Utils.getRemoveSignersStructHash(signersToRemove, newThreshold, custody.signerNonce(), deadline);
        bytes32 digest = Utils.getTypedDataHash(domainSeparator, structHash);

        // Sign the digest
        bytes memory signature = Utils.signDigest(vm, signerKey, digest);

        console.log("Contract:", contractAddr);
        console.log("Removing %d signer(s), new threshold: %d", signersToRemove.length, newThreshold);
        console.log("Resulting signer count:", resultingSignerCount);
        console.log("Signature:");
        console.logBytes(signature);
    }
}
