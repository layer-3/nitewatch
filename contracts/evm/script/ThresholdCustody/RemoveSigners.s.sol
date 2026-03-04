// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ThresholdCustody} from "../../src/ThresholdCustody.sol";
import {ThresholdCustodyScriptUtils as Utils} from "./Utils.sol";

/// @notice Submit transaction to remove signers from ThresholdCustody contract.
///
/// Environment variables:
///   TX_SENDER_PK       – Transaction sender's private key
///   CONTRACT           – ThresholdCustody contract address
///   SIGNERS_TO_REMOVE  – Comma-separated addresses to remove
///   NEW_THRESHOLD      – New threshold value (optional, defaults to current threshold)
///   DEADLINE           – Unix timestamp deadline for the operation
///   SIGNATURES         – Comma-separated signatures from signers (hex, 0x-prefixed)
///
/// Example:
///   TX_SENDER_PK=0xabc... CONTRACT=0x... SIGNERS_TO_REMOVE=0xA NEW_THRESHOLD=2 DEADLINE=1999999999 \
///     SIGNATURES=0x1234...,0x5678... \
///     forge script script/ThresholdCustody/RemoveSigners.s.sol --rpc-url $RPC_URL --broadcast
contract RemoveSigners is Script {
    function run() public {
        uint256 txSenderKey = vm.envUint("TX_SENDER_PK");
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

        // Get signatures and encode them in MultiSignerERC7913 format
        bytes[] memory sigArray = vm.envBytes("SIGNATURES", ",");

        // Build digest to recover signer addresses
        bytes32 domainSeparator = Utils.getDomainSeparator(contractAddr);
        bytes32 structHash =
            Utils.getRemoveSignersStructHash(signersToRemove, newThreshold, custody.signerNonce(), deadline);
        bytes32 digest = Utils.getTypedDataHash(domainSeparator, structHash);

        bytes memory signatures = Utils.encodeMultiSignerSignatures(digest, sigArray);

        console.log("Contract:", contractAddr);
        console.log("Removing %d signer(s), new threshold: %d", signersToRemove.length, newThreshold);
        console.log("Signatures provided:", sigArray.length);
        console.log("Resulting signer count:", resultingSignerCount);

        vm.startBroadcast(txSenderKey);
        custody.removeSigners(signersToRemove, newThreshold, deadline, signatures);
        vm.stopBroadcast();

        console.log("Signers removed successfully.");
    }
}
