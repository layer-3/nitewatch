// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ThresholdCustody} from "../../src/ThresholdCustody.sol";
import {ThresholdCustodyScriptUtils as Utils} from "./Utils.sol";

/// @notice Submit transaction to add signers to ThresholdCustody contract.
///
/// Environment variables:
///   TX_SENDER_PK   – Transaction sender's private key
///   CONTRACT       – ThresholdCustody contract address
///   NEW_SIGNERS    – Comma-separated addresses to add
///   NEW_THRESHOLD  – New threshold value (optional, defaults to current threshold)
///   DEADLINE       – Unix timestamp deadline for the operation
///   SIGNATURES     – Comma-separated signatures from signers (hex, 0x-prefixed)
///
/// Example:
///   TX_SENDER_PK=0xabc... CONTRACT=0x... NEW_SIGNERS=0xA,0xB NEW_THRESHOLD=3 DEADLINE=1999999999 \
///     SIGNATURES=0x1234...,0x5678... \
///     forge script script/ThresholdCustody/AddSigners.s.sol --rpc-url $RPC_URL --broadcast
contract AddSigners is Script {
    function run() public {
        uint256 txSenderKey = vm.envUint("TX_SENDER_PK");
        address contractAddr = vm.envAddress("CONTRACT");
        address[] memory newSigners = vm.envAddress("NEW_SIGNERS", ",");
        uint256 deadline = vm.envUint("DEADLINE");

        ThresholdCustody custody = ThresholdCustody(contractAddr);

        // Get threshold: use NEW_THRESHOLD if provided, otherwise use current threshold
        uint64 newThreshold = Utils.getThreshold(vm, "NEW_THRESHOLD", custody);

        // Calculate resulting signer count
        bytes[] memory currentSigners = custody.getSigners(0, type(uint64).max);
        uint256 resultingSignerCount = currentSigners.length + newSigners.length;

        // Validate threshold
        Utils.validateThreshold(newThreshold, resultingSignerCount);

        // Get signatures and concatenate them
        bytes[] memory sigArray = vm.envBytes("SIGNATURES", ",");
        bytes memory signatures = Utils.concatenateSignatures(sigArray);

        console.log("Contract:", contractAddr);
        console.log("Adding %d signer(s), new threshold: %d", newSigners.length, newThreshold);
        console.log("Signatures provided:", sigArray.length);
        console.log("Resulting signer count:", resultingSignerCount);

        vm.startBroadcast(txSenderKey);
        custody.addSigners(newSigners, newThreshold, deadline, signatures);
        vm.stopBroadcast();

        console.log("Signers added successfully.");
    }
}
