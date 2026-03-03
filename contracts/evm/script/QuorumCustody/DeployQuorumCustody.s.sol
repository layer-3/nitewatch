// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {QuorumCustody} from "../src/QuorumCustody.sol";

contract DeployQuorumCustody is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        uint64 initialQuorum = uint64(vm.envUint("INITIAL_QUORUM"));

        address[] memory initialSigners = vm.envAddress("SIGNERS", ",");

        vm.startBroadcast(deployerPrivateKey);

        QuorumCustody quorumCustody = new QuorumCustody(initialSigners, initialQuorum);
        console.log("QuorumCustody deployed at:", address(quorumCustody));

        vm.stopBroadcast();
    }
}
