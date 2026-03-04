// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol";

import {MessageHashUtils} from "openzeppelin-contracts/contracts/utils/cryptography/MessageHashUtils.sol";
import {ECDSA} from "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import {
    ADD_SIGNERS_TYPEHASH,
    REMOVE_SIGNERS_TYPEHASH,
    SET_THRESHOLD_TYPEHASH,
    NAME,
    VERSION,
    ThresholdCustody
} from "../../src/ThresholdCustody.sol";
import {Utils as ContractUtils} from "../../src/Utils.sol";

library ThresholdCustodyScriptUtils {
    using {ContractUtils.hashArrayed} for address[];

    error InsufficientThreshold(uint64 threshold, uint256 signerCount);

    // EIP-712 domain fields: name (0x01) | version (0x02) | chainId (0x04) | verifyingContract (0x08)
    bytes1 constant DOMAIN_FIELDS = 0x0f;

    /// @notice Calculate the EIP-712 domain separator for ThresholdCustody
    function getDomainSeparator(address contractAddr) internal view returns (bytes32) {
        return MessageHashUtils.toDomainSeparator(
            DOMAIN_FIELDS,
            NAME,
            VERSION,
            block.chainid,
            contractAddr,
            bytes32(0) // salt not used
        );
    }

    /// @notice Calculate the EIP-712 typed data hash
    function getTypedDataHash(bytes32 domainSeparator, bytes32 structHash) internal pure returns (bytes32) {
        return MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
    }

    /// @notice Build struct hash for addSigners operation
    function getAddSignersStructHash(address[] memory newSigners, uint64 newThreshold, uint256 nonce, uint256 deadline)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(ADD_SIGNERS_TYPEHASH, newSigners.hashArrayed(), newThreshold, nonce, deadline));
    }

    /// @notice Build struct hash for removeSigners operation
    function getRemoveSignersStructHash(
        address[] memory signersToRemove,
        uint64 newThreshold,
        uint256 nonce,
        uint256 deadline
    ) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(REMOVE_SIGNERS_TYPEHASH, signersToRemove.hashArrayed(), newThreshold, nonce, deadline)
        );
    }

    /// @notice Build struct hash for setThreshold operation
    function getSetThresholdStructHash(uint64 newThreshold, uint256 nonce, uint256 deadline)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(SET_THRESHOLD_TYPEHASH, newThreshold, nonce, deadline));
    }

    /// @notice Sign a digest using vm.sign
    function signDigest(Vm vm, uint256 privateKey, bytes32 digest) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    /// @notice Validate that threshold can be reached with the given number of signers
    function validateThreshold(uint64 threshold, uint256 signerCount) internal pure {
        require(threshold <= signerCount, InsufficientThreshold(threshold, signerCount));
    }

    /// @notice Encode signatures in MultiSignerERC7913 format
    /// @dev Recovers signer addresses from signatures and encodes them with the signatures
    function encodeMultiSignerSignatures(bytes32 digest, bytes[] memory sigArray) internal pure returns (bytes memory) {
        // Recover signer addresses from signatures
        address[] memory signerAddrs = new address[](sigArray.length);
        for (uint256 i = 0; i < sigArray.length; i++) {
            signerAddrs[i] = ECDSA.recover(digest, sigArray[i]);
        }

        // Encode signers as bytes[]
        bytes[] memory signers = new bytes[](signerAddrs.length);
        for (uint256 i = 0; i < signerAddrs.length; i++) {
            signers[i] = abi.encodePacked(signerAddrs[i]);
        }

        return abi.encode(signers, sigArray);
    }

    /// @notice Get threshold from env or use current threshold from contract
    function getThreshold(Vm vm, string memory thresholdEnvName, ThresholdCustody custody)
        internal
        view
        returns (uint64)
    {
        // Use envOr with 0 as default - if 0, use current threshold from contract
        uint256 envThreshold = vm.envOr(thresholdEnvName, uint256(0));

        if (envThreshold == 0) {
            uint64 currentThreshold = custody.threshold();
            console.log(thresholdEnvName, " env not specified or zero, using existing threshold:", currentThreshold);
            return currentThreshold;
        }

        return uint64(envThreshold);
    }
}
