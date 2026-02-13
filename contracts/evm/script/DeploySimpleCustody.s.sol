// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {SimpleCustody} from "../src/SimpleCustody.sol";

contract DeploySimpleCustody is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address admin = vm.envAddress("ADMIN_ADDRESS");
        address neodax = vm.envAddress("NEODAX_ADDRESS");
        address nitewatch = vm.envAddress("NITEWATCH_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        SimpleCustody simpleCustody = new SimpleCustody(admin, neodax, nitewatch);
        console.log("SimpleCustody deployed at:", address(simpleCustody));

        vm.stopBroadcast();
    }
}
