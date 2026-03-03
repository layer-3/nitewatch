// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ThresholdCustody} from "../../src/ThresholdCustody.sol";
import {ThresholdCustodyScriptUtils as Utils} from "./Utils.sol";

/// @notice Submit transaction to set threshold in ThresholdCustody contract.
///
/// Environment variables:
///   TX_SENDER_PK   – Transaction sender's private key
///   CONTRACT       – ThresholdCustody contract address
///   NEW_THRESHOLD  – New threshold value
///   DEADLINE       – Unix timestamp deadline for the operation
///   SIGNATURES     – Comma-separated signatures from signers (hex, 0x-prefixed)
///
/// Example:
///   TX_SENDER_PK=0xabc... CONTRACT=0x... NEW_THRESHOLD=3 DEADLINE=1999999999 \
///     SIGNATURES=0x1234...,0x5678... \
///     forge script script/ThresholdCustody/SetThreshold.s.sol --rpc-url $RPC_URL --broadcast
contract SetThreshold is Script {
    function run() public {
        uint256 txSenderKey = vm.envUint("TX_SENDER_PK");
        address contractAddr = vm.envAddress("CONTRACT");
        uint64 newThreshold = uint64(vm.envUint("NEW_THRESHOLD"));
        uint256 deadline = vm.envUint("DEADLINE");

        ThresholdCustody custody = ThresholdCustody(contractAddr);

        // Get current signer count
        bytes[] memory currentSigners = custody.getSigners(0, type(uint64).max);
        uint256 signerCount = currentSigners.length;

        // Validate threshold
        Utils.validateThreshold(newThreshold, signerCount);

        // Get signatures and concatenate them
        bytes[] memory sigArray = vm.envBytes("SIGNATURES", ",");
        bytes memory signatures = Utils.concatenateSignatures(sigArray);

        console.log("Contract:", contractAddr);
        console.log("New threshold:", newThreshold);
        console.log("Signatures provided:", sigArray.length);
        console.log("Current signer count:", signerCount);

        vm.startBroadcast(txSenderKey);
        custody.setThreshold(newThreshold, deadline, signatures);
        vm.stopBroadcast();

        console.log("Threshold set successfully.");
    }
}
