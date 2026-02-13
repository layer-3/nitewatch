// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {QuorumCustody} from "../src/QuorumCustody.sol";

contract DeployQuorumCustody is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address initialSigner = vm.envAddress("INITIAL_SIGNER_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        QuorumCustody quorumCustody = new QuorumCustody(initialSigner);
        console.log("QuorumCustody deployed at:", address(quorumCustody));

        vm.stopBroadcast();
    }
}
