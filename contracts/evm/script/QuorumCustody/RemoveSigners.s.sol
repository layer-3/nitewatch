// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {QuorumCustody} from "../src/QuorumCustody.sol";

/// @notice Remove signers from an existing QuorumCustody contract.
///
/// Environment variables:
///   PRIVATE_KEY        – Caller's private key (must be an existing signer)
///   CONTRACT           – QuorumCustody contract address
///   SIGNERS_TO_REMOVE  – Comma-separated addresses to remove
///   NEW_QUORUM         – Quorum value after the removal
///   DEADLINE           – Unix timestamp deadline for the operation
///   SIGNATURES         – Comma-separated EIP-712 signatures from other signers (hex, 0x-prefixed).
///                        Omit or leave empty when current quorum is 1 (caller alone suffices).
///
/// Example:
///   PRIVATE_KEY=abc... CONTRACT=0x... SIGNERS_TO_REMOVE=0xA NEW_QUORUM=2 DEADLINE=1999999999 \
///     forge script script/RemoveSigners.s.sol --rpc-url $RPC_URL --broadcast
contract RemoveSigners is Script {
    function run() public {
        uint256 callerKey = vm.envUint("PRIVATE_KEY");
        address contractAddr = vm.envAddress("CONTRACT");
        address[] memory signersToRemove = vm.envAddress("SIGNERS_TO_REMOVE", ",");
        uint64 newQuorum = uint64(vm.envUint("NEW_QUORUM"));
        uint256 deadline = vm.envUint("DEADLINE");

        bytes[] memory sigs;
        try vm.envBytes("SIGNATURES", ",") returns (bytes[] memory s) {
            sigs = s;
        } catch {
            sigs = new bytes[](0);
        }

        QuorumCustody custody = QuorumCustody(contractAddr);

        console.log("Contract:", contractAddr);
        console.log("Removing %d signer(s), new quorum: %d", signersToRemove.length, newQuorum);
        console.log("Co-signatures provided:", sigs.length);

        vm.startBroadcast(callerKey);
        custody.removeSigners(signersToRemove, newQuorum, deadline, sigs);
        vm.stopBroadcast();

        console.log("Signers removed successfully. Total signers:", custody.getSignerCount());
    }
}
