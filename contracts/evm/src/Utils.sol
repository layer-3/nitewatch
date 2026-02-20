// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

library Utils {
    function getWithdrawalId(address custodyAddress, address user, address token, uint256 amount, uint256 nonce)
        internal
        view
        returns (bytes32)
    {
        return keccak256(abi.encode(block.chainid, custodyAddress, user, token, amount, nonce));
    }

    // Helpers for conversion
    function toBytes(address a) internal pure returns (bytes memory) {
        return abi.encodePacked(a);
    }

    function toAddressBytesArray(address[] memory addrs) internal pure returns (bytes[] memory) {
        bytes[] memory b = new bytes[](addrs.length);
        for (uint256 i = 0; i < addrs.length; i++) {
            b[i] = toBytes(addrs[i]);
        }
        return b;
    }

    function hashArrayed(address[] memory arr) internal pure returns (bytes32) {
        bytes32[] memory encoded = new bytes32[](arr.length);
        for (uint256 i = 0; i < arr.length; i++) {
            encoded[i] = bytes32(uint256(uint160(arr[i])));
        }
        return keccak256(abi.encodePacked(encoded));
    }

    function toAddress(bytes memory b) internal pure returns (address) {
        return address(bytes20(b));
    }
}
