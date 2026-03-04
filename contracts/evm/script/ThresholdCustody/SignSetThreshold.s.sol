// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ThresholdCustody} from "../../src/ThresholdCustody.sol";
import {ThresholdCustodyScriptUtils as Utils} from "./Utils.sol";

/// @notice Generate a signature for setting threshold in ThresholdCustody contract.
///
/// Environment variables:
///   SIGNER_PK      – Signer's private key
///   CONTRACT       – ThresholdCustody contract address
///   NEW_THRESHOLD  – New threshold value
///   DEADLINE       – Unix timestamp deadline for the operation
///
/// Example:
///   SIGNER_PK=0xabc... CONTRACT=0x... NEW_THRESHOLD=3 DEADLINE=1999999999 \
///     forge script script/ThresholdCustody/SignSetThreshold.s.sol --rpc-url $RPC_URL
contract SignSetThreshold is Script {
    function run() public view returns (bytes memory) {
        uint256 signerKey = vm.envUint("SIGNER_PK");
        address contractAddr = vm.envAddress("CONTRACT");
        uint64 newThreshold = uint64(vm.envUint("NEW_THRESHOLD"));
        uint256 deadline = vm.envUint("DEADLINE");

        ThresholdCustody custody = ThresholdCustody(contractAddr);

        // Get current signer count
        bytes[] memory currentSigners = custody.getSigners(0, type(uint64).max);
        uint256 signerCount = currentSigners.length;

        // Validate threshold
        Utils.validateThreshold(newThreshold, signerCount);

        // Build EIP-712 digest
        bytes32 domainSeparator = Utils.getDomainSeparator(contractAddr);
        bytes32 structHash = Utils.getSetThresholdStructHash(newThreshold, custody.signerNonce(), deadline);
        bytes32 digest = Utils.getTypedDataHash(domainSeparator, structHash);

        // Sign the digest
        bytes memory signature = Utils.signDigest(vm, signerKey, digest);

        console.log("Contract:", contractAddr);
        console.log("New threshold:", newThreshold);
        console.log("Current signer count:", signerCount);
        console.log("Signature:");
        console.logBytes(signature);

        return signature;
    }
}
