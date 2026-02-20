// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {ThresholdCustody} from "../src/ThresholdCustody.sol";

/// @notice Harness contract for testing internal functions of ThresholdCustody
contract TestThresholdCustody is ThresholdCustody {
    constructor(address[] memory initialSigners, uint64 threshold) ThresholdCustody(initialSigners, threshold) {}

    /// @notice Exposes the internal _executeWithdrawal function for testing
    function exposed_executeWithdrawal(bytes32 withdrawalId) external {
        WithdrawalRequest storage request = withdrawals[withdrawalId];
        _executeWithdrawal(request);
    }

    /// @notice Exposes the internal _countValidApprovals function for testing
    function exposed_countValidApprovals(bytes32 withdrawalId) external view returns (uint256) {
        return _countValidApprovals(withdrawalId);
    }

    /// @notice Helper to set withdrawal approvals directly for testing
    function workaround_setWithdrawalApproval(bytes32 withdrawalId, address signer, bool hasApproved) external {
        withdrawalApprovals[withdrawalId][signer] = hasApproved;
    }
}
