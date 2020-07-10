// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baklava

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getQueue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"setBaselineQuorumFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHotfixRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stageDurations\",\"outputs\":[{\"name\":\"approval\",\"type\":\"uint256\"},{\"name\":\"referendum\",\"type\":\"uint256\"},{\"name\":\"execution\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"concurrentProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isDequeuedProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participationFloor\",\"type\":\"uint256\"}],\"name\":\"setParticipationFloor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"blsKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isProposalPassing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMostRecentReferendumProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_queueExpiry\",\"type\":\"uint256\"}],\"name\":\"setQueueExpiry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExecutionStageDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"setApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isHotfixPassing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"dequeueProposalsIfReady\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participationBaseline\",\"type\":\"uint256\"}],\"name\":\"setParticipationBaseline\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"whitelister\",\"type\":\"address\"}],\"name\":\"isHotfixWhitelistedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hotfixes\",\"outputs\":[{\"name\":\"executed\",\"type\":\"bool\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"preparedEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"lesser\",\"type\":\"uint256\"},{\"name\":\"greater\",\"type\":\"uint256\"}],\"name\":\"upvote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"}],\"name\":\"setBaselineUpdateFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoteRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVoting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundedDeposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"destinations\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"dataLengths\",\"type\":\"uint256[]\"},{\"name\":\"descriptionUrl\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"executionStageDuration\",\"type\":\"uint256\"}],\"name\":\"setExecutionStageDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDequeue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isDequeuedProposalExpired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dequeueFrequency\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dequeueFrequency\",\"type\":\"uint256\"}],\"name\":\"setDequeueFrequency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"hotfixWhitelistValidatorTally\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"queueExpiry\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"functionId\",\"type\":\"bytes4\"}],\"name\":\"getConstitution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getUpvotes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"approvalStageDuration\",\"type\":\"uint256\"}],\"name\":\"setApprovalStageDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"prepareHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"emptyIndices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReferendumStageDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dequeued\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lesser\",\"type\":\"uint256\"},{\"name\":\"greater\",\"type\":\"uint256\"}],\"name\":\"revokeUpvote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approveHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"whitelistHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDequeue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isQueuedProposalExpired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"_approver\",\"type\":\"address\"},{\"name\":\"_concurrentProposals\",\"type\":\"uint256\"},{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_queueExpiry\",\"type\":\"uint256\"},{\"name\":\"_dequeueFrequency\",\"type\":\"uint256\"},{\"name\":\"approvalStageDuration\",\"type\":\"uint256\"},{\"name\":\"referendumStageDuration\",\"type\":\"uint256\"},{\"name\":\"executionStageDuration\",\"type\":\"uint256\"},{\"name\":\"participationBaseline\",\"type\":\"uint256\"},{\"name\":\"participationFloor\",\"type\":\"uint256\"},{\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"},{\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isQueued\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParticipationParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_concurrentProposals\",\"type\":\"uint256\"}],\"name\":\"setConcurrentProposals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUpvoteRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"referendumStageDuration\",\"type\":\"uint256\"}],\"name\":\"setReferendumStageDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"destinations\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"dataLengths\",\"type\":\"uint256[]\"},{\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getProposalTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoteTotals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"functionId\",\"type\":\"bytes4\"},{\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setConstitution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalStageDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproverSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"concurrentProposals\",\"type\":\"uint256\"}],\"name\":\"ConcurrentProposalsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minDeposit\",\"type\":\"uint256\"}],\"name\":\"MinDepositSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queueExpiry\",\"type\":\"uint256\"}],\"name\":\"QueueExpirySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dequeueFrequency\",\"type\":\"uint256\"}],\"name\":\"DequeueFrequencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"approvalStageDuration\",\"type\":\"uint256\"}],\"name\":\"ApprovalStageDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"referendumStageDuration\",\"type\":\"uint256\"}],\"name\":\"ReferendumStageDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"executionStageDuration\",\"type\":\"uint256\"}],\"name\":\"ExecutionStageDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"functionId\",\"type\":\"bytes4\"},{\"indexed\":false,\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ConstitutionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transactionCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"upvotes\",\"type\":\"uint256\"}],\"name\":\"ProposalUpvoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"revokedUpvotes\",\"type\":\"uint256\"}],\"name\":\"ProposalUpvoteRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProposalDequeued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"participationBaseline\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"participationFloor\",\"type\":\"uint256\"}],\"name\":\"ParticipationFloorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineUpdateFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineQuorumFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"whitelister\",\"type\":\"address\"}],\"name\":\"HotfixWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"HotfixApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"HotfixPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"HotfixExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// GovernanceBin is the compiled bytecode used for deploying new contracts.
var GovernanceBin = "0x60806040526000620000196401000000006200006d810204565b60008054600160a060020a031916600160a060020a0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001805562000071565b3390565b618be280620000816000396000f3fe6080604052600436106104cf5760003560e060020a900480637b1039991161027a578063b8f7700511610153578063cf48eb94116100c5578063e50e652d11610089578063e50e652d14611a31578063ec68307214611a5b578063ed38527414611aa3578063f2fde38b14611aec578063fae8db0a14611b1f578063ffea74c014611b49576104cf565b8063cf48eb94146117a2578063d704f0c514611915578063da35c664146119dd578063df4da461146119f2578063e41db45514611a07576104cf565b8063c73a6d7811610117578063c73a6d78146115e6578063c7f758a814611610578063c805956d146116e0578063c8d8d2b51461171b578063cd845a7614611745578063cea69e7414611778576104cf565b8063b8f77005146114d4578063bbb2eab9146114e9578063c0aee5f414611522578063c134b2fc14611537578063c1939b2014611561576104cf565b80639a6c3d83116101ec578063aa2feb83116101b0578063aa2feb83146113e7578063ad78c10914611411578063add004df14611426578063af108a0e14611450578063b0f9984214611480578063b15f0f58146114aa576104cf565b80639a6c3d83146113215780639a7b3be71461134b5780639b2b592f146113605780639cb02dfc1461138a578063a91ee0dc146113b4576104cf565b80638da5cb5b1161023e5780638da5cb5b1461124b5780638e905ed6146112605780638f32d59b146112755780638fcc9cfb1461128a57806397b9eba6146112b457806398f42702146112f7576104cf565b80637b1039991461111a5780638018556e1461112f57806381d4728d1461115957806387ee8a0f146111835780638a88362614611198576104cf565b80633fa5fed0116103ac5780635f8dd6491161031e5780636de8a63b116102e25780636de8a63b146110225780636f2ab69314611087578063715018a6146110b15780637385e5da146110c657806377d26a2a146110db5780637910867b146110f0576104cf565b80635f8dd64914610d2757806360b4d34d14610d5a57806365bbdaa014610d8d5780636643ac5814610f4557806367960e9114610f6f576104cf565b806357333978116103705780635733397814610bc7578063582ae53b14610bfd5780635c75939414610c4b5780635d180adb14610c755780635d35a3d914610ca55780635f115a8514610cd5576104cf565b80633fa5fed014610a6c57806341b3d18514610aa557806345a7849914610aba5780634b2c2f4414610ae45780635601eaea14610b97576104cf565b806323f0ab65116104455780633156560e116104095780633156560e14610991578063344944cf146109c45780633b1eb4bf146109ee5780633bb0ed2b14610a185780633ccfd60b14610a2d5780633db9dd9a14610a42576104cf565b806323f0ab65146107ab57806327621321146108f5578063283aaefb1461091f5780632c0523551461095257806330a095d01461097c576104cf565b8063123633ea11610497578063123633ea146106a35780631374b22d146106e9578063141a8dd814610727578063152b48341461073c578063158ef93e1461076c5780631c65bc6114610781576104cf565b806301fce27e1461052757806304acaec9146105d55780630e0b78ae146105ff5780630f717e42146106495780631201a0fb1461067c575b3615610525576040805160e560020a62461bcd02815260206004820152600e60248201527f756e6b6e6f776e206d6574686f64000000000000000000000000000000000000604482015290519081900360640190fd5b005b34801561053357600080fd5b5061053c611b5e565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610580578181015183820152602001610568565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156105bf5781810151838201526020016105a7565b5050505090500194505050505060405180910390f35b3480156105e157600080fd5b50610525600480360360208110156105f857600080fd5b5035611c95565b34801561060b57600080fd5b506106296004803603602081101561062257600080fd5b5035611def565b604080519315158452911515602084015282820152519081900360600190f35b34801561065557600080fd5b5061065e611e14565b60408051938452602084019290925282820152519081900360600190f35b34801561068857600080fd5b50610691611e20565b60408051918252519081900360200190f35b3480156106af57600080fd5b506106cd600480360360208110156106c657600080fd5b5035611e26565b60408051600160a060020a039092168252519081900360200190f35b3480156106f557600080fd5b506107136004803603602081101561070c57600080fd5b5035611f3c565b604080519115158252519081900360200190f35b34801561073357600080fd5b506106cd611f59565b34801561074857600080fd5b506107136004803603604081101561075f57600080fd5b5080359060200135611f68565b34801561077857600080fd5b50610713611f88565b34801561078d57600080fd5b50610525600480360360208110156107a457600080fd5b5035611fa9565b3480156107b757600080fd5b50610713600480360360608110156107ce57600080fd5b600160a060020a0382351691908101906040810160208201356401000000008111156107f957600080fd5b82018360208201111561080b57600080fd5b8035906020019184600183028401116401000000008311171561082d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561088057600080fd5b82018360208201111561089257600080fd5b803590602001918460018302840111640100000000831117156108b457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506120ed945050505050565b34801561090157600080fd5b506107136004803603602081101561091857600080fd5b5035612275565b34801561092b57600080fd5b506106916004803603602081101561094257600080fd5b5035600160a060020a031661228c565b34801561095e57600080fd5b506105256004803603602081101561097557600080fd5b50356122aa565b34801561098857600080fd5b506106916123cd565b34801561099d57600080fd5b50610525600480360360208110156109b457600080fd5b5035600160a060020a03166123d4565b3480156109d057600080fd5b50610713600480360360208110156109e757600080fd5b503561253d565b3480156109fa57600080fd5b5061069160048036036020811015610a1157600080fd5b5035612558565b348015610a2457600080fd5b5061052561256b565b348015610a3957600080fd5b50610713612864565b348015610a4e57600080fd5b5061052560048036036020811015610a6557600080fd5b50356129bd565b348015610a7857600080fd5b5061071360048036036040811015610a8f57600080fd5b5080359060200135600160a060020a0316612b17565b348015610ab157600080fd5b50610691612b46565b348015610ac657600080fd5b5061062960048036036020811015610add57600080fd5b5035612b4c565b348015610af057600080fd5b5061069160048036036020811015610b0757600080fd5b810190602081018135640100000000811115610b2257600080fd5b820183602082011115610b3457600080fd5b80359060200191846001830284011164010000000083111715610b5657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612b72945050505050565b348015610ba357600080fd5b5061071360048036036040811015610bba57600080fd5b5080359060200135612ccf565b348015610bd357600080fd5b5061071360048036036060811015610bea57600080fd5b5080359060208101359060400135612e5b565b348015610c0957600080fd5b50610c2760048036036020811015610c2057600080fd5b50356133ca565b60405180826005811115610c3757fe5b60ff16815260200191505060405180910390f35b348015610c5757600080fd5b5061052560048036036020811015610c6e57600080fd5b5035613423565b348015610c8157600080fd5b506106cd60048036036040811015610c9857600080fd5b508035906020013561357d565b348015610cb157600080fd5b5061071360048036036040811015610cc857600080fd5b5080359060200135613693565b348015610ce157600080fd5b50610d0e60048036036040811015610cf857600080fd5b50600160a060020a0381351690602001356138a5565b6040805192835260208301919091528051918290030190f35b348015610d3357600080fd5b5061071360048036036020811015610d4a57600080fd5b5035600160a060020a03166138f1565b348015610d6657600080fd5b5061069160048036036020811015610d7d57600080fd5b5035600160a060020a031661396a565b610691600480360360a0811015610da357600080fd5b810190602081018135640100000000811115610dbe57600080fd5b820183602082011115610dd057600080fd5b80359060200191846020830284011164010000000083111715610df257600080fd5b919390929091602081019035640100000000811115610e1057600080fd5b820183602082011115610e2257600080fd5b80359060200191846020830284011164010000000083111715610e4457600080fd5b919390929091602081019035640100000000811115610e6257600080fd5b820183602082011115610e7457600080fd5b80359060200191846001830284011164010000000083111715610e9657600080fd5b919390929091602081019035640100000000811115610eb457600080fd5b820183602082011115610ec657600080fd5b80359060200191846020830284011164010000000083111715610ee857600080fd5b919390929091602081019035640100000000811115610f0657600080fd5b820183602082011115610f1857600080fd5b80359060200191846001830284011164010000000083111715610f3a57600080fd5b50909250905061397c565b348015610f5157600080fd5b5061052560048036036020811015610f6857600080fd5b5035613c96565b348015610f7b57600080fd5b5061069160048036036020811015610f9257600080fd5b810190602081018135640100000000811115610fad57600080fd5b820183602082011115610fbf57600080fd5b80359060200191846001830284011164010000000083111715610fe157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613dcf945050505050565b34801561102e57600080fd5b50611037613f21565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561107357818101518382015260200161105b565b505050509050019250505060405180910390f35b34801561109357600080fd5b50610713600480360360208110156110aa57600080fd5b5035613f79565b3480156110bd57600080fd5b50610525613fa1565b3480156110d257600080fd5b50610691614044565b3480156110e757600080fd5b50610691614054565b3480156110fc57600080fd5b506107136004803603602081101561111357600080fd5b503561405a565b34801561112657600080fd5b506106cd614071565b34801561113b57600080fd5b506105256004803603602081101561115257600080fd5b5035614080565b34801561116557600080fd5b506106916004803603602081101561117c57600080fd5b50356141a3565b34801561118f57600080fd5b506106916142b6565b3480156111a457600080fd5b50610691600480360360208110156111bb57600080fd5b8101906020810181356401000000008111156111d657600080fd5b8201836020820111156111e857600080fd5b8035906020019184600183028401116401000000008311171561120a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506143d3945050505050565b34801561125757600080fd5b506106cd614525565b34801561126c57600080fd5b50610691614534565b34801561128157600080fd5b5061071361453a565b34801561129657600080fd5b50610525600480360360208110156112ad57600080fd5b503561455e565b3480156112c057600080fd5b50610691600480360360408110156112d757600080fd5b508035600160a060020a03169060200135600160e060020a031916614697565b34801561130357600080fd5b506106916004803603602081101561131a57600080fd5b50356146ab565b34801561132d57600080fd5b506105256004803603602081101561134457600080fd5b50356147b0565b34801561135757600080fd5b506106916148e9565b34801561136c57600080fd5b506106916004803603602081101561138357600080fd5b50356148f4565b34801561139657600080fd5b50610525600480360360208110156113ad57600080fd5b50356149ff565b3480156113c057600080fd5b50610525600480360360208110156113d757600080fd5b5035600160a060020a0316614b41565b3480156113f357600080fd5b506106916004803603602081101561140a57600080fd5b5035614c44565b34801561141d57600080fd5b50610691614c63565b34801561143257600080fd5b506106916004803603602081101561144957600080fd5b5035614c69565b34801561145c57600080fd5b506107136004803603604081101561147357600080fd5b5080359060200135614c77565b34801561148c57600080fd5b50610525600480360360208110156114a357600080fd5b5035615019565b3480156114b657600080fd5b50610525600480360360208110156114cd57600080fd5b5035615118565b3480156114e057600080fd5b506106916151cf565b3480156114f557600080fd5b506107136004803603606081101561150c57600080fd5b508035906020810135906040013560ff166151d5565b34801561152e57600080fd5b5061069161574b565b34801561154357600080fd5b506107136004803603602081101561155a57600080fd5b5035615751565b34801561156d57600080fd5b5061052560048036036101a081101561158557600080fd5b50600160a060020a03813581169160208101359091169060408101359060608101359060808101359060a08101359060c08101359060e081013590610100810135906101208101359061014081013590610160810135906101800135615768565b3480156115f257600080fd5b506107136004803603602081101561160957600080fd5b50356158a1565b34801561161c57600080fd5b5061163a6004803603602081101561163357600080fd5b5035615943565b6040518086600160a060020a0316600160a060020a0316815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156116a1578181015183820152602001611689565b50505050905090810190601f1680156116ce5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156116ec57600080fd5b506116f5615a06565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561172757600080fd5b506105256004803603602081101561173e57600080fd5b5035615a81565b34801561175157600080fd5b50610d0e6004803603602081101561176857600080fd5b5035600160a060020a0316615ba4565b34801561178457600080fd5b506105256004803603602081101561179b57600080fd5b5035615bea565b3480156117ae57600080fd5b50610525600480360360a08110156117c557600080fd5b8101906020810181356401000000008111156117e057600080fd5b8201836020820111156117f257600080fd5b8035906020019184602083028401116401000000008311171561181457600080fd5b91939092909160208101903564010000000081111561183257600080fd5b82018360208201111561184457600080fd5b8035906020019184602083028401116401000000008311171561186657600080fd5b91939092909160208101903564010000000081111561188457600080fd5b82018360208201111561189657600080fd5b803590602001918460018302840111640100000000831117156118b857600080fd5b9193909290916020810190356401000000008111156118d657600080fd5b8201836020820111156118e857600080fd5b8035906020019184602083028401116401000000008311171561190a57600080fd5b919350915035615d23565b34801561192157600080fd5b506119456004803603604081101561193857600080fd5b5080359060200135616067565b6040518084815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156119a0578181015183820152602001611988565b50505050905090810190601f1680156119cd5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156119e957600080fd5b5061069161618e565b3480156119fe57600080fd5b50610691616194565b348015611a1357600080fd5b5061065e60048036036020811015611a2a57600080fd5b5035616284565b348015611a3d57600080fd5b5061069160048036036020811015611a5457600080fd5b50356162ac565b348015611a6757600080fd5b50610d0e600480360360c0811015611a7e57600080fd5b5080359060208101359060408101359060608101359060808101359060a001356162ea565b348015611aaf57600080fd5b5061052560048036036060811015611ac657600080fd5b50600160a060020a0381351690600160e060020a031960208201351690604001356164b1565b348015611af857600080fd5b5061052560048036036020811015611b0f57600080fd5b5035600160a060020a0316616688565b348015611b2b57600080fd5b5061069160048036036020811015611b4257600080fd5b50356166e0565b348015611b5557600080fd5b506106916167eb565b606080601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__6369b317e390916040518263ffffffff1660e060020a0281526004018082815260200191505060006040518083038186803b158015611bb657600080fd5b505af4158015611bca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015611bf357600080fd5b810190808051640100000000811115611c0b57600080fd5b82016020810184811115611c1e57600080fd5b8151856020820283011164010000000082111715611c3b57600080fd5b50509291906020018051640100000000811115611c5757600080fd5b82016020810184811115611c6a57600080fd5b8151856020820283011164010000000082111715611c8757600080fd5b509496509450505050509091565b611c9d61453a565b1515611ce1576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b611ce96183f2565b611cf2826167f1565b9050611cfd8161680b565b1515611d3d5760405160e560020a62461bcd0281526004018080602001828103825260278152602001806187506027913960400191505060405180910390fd5b6040805160208101909152601c548152611d5e90829063ffffffff61681e16565b15611db3576040805160e560020a62461bcd02815260206004820181905260248201527f426173656c696e652071756f72756d20666163746f7220756e6368616e676564604482015290519081900360640190fd5b8051601c556040805183815290517fddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc9181900360200190a15050565b6000908152601160205260409020805460019091015460ff6101008304811693921691565b60035460045460055483565b600a5481565b60408051602080820184905243828401528251808303840181526060928301938490528051600094859360fa939282918401908083835b60208310611e7c5780518252601f199092019160209182019101611e5d565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114611edc576040519150601f19603f3d011682016040523d82523d6000602084013e611ee1565b606091505b5092509050801515611f275760405160e560020a62461bcd02815260040180806020018281038252603d8152602001806188a3603d913960400191505060405180910390fd5b611f32826000616825565b925050505b919050565b6000818152600f60205260408120611f5390616831565b92915050565b600854600160a060020a031681565b6000828152600f60205260408120611f8190848461683b565b9392505050565b60005474010000000000000000000000000000000000000000900460ff1681565b611fb161453a565b1515611ff5576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b611ffd6183f2565b612006826167f1565b90506120118161680b565b15156120515760405160e560020a62461bcd028152600401808060200182810382526024815260200180618abd6024913960400191505060405180910390fd5b6040805160208101909152601a54815261207290829063ffffffff61681e16565b156120b15760405160e560020a62461bcd0281526004018080602001828103825260268152602001806187bf6026913960400191505060405180910390fd5b8051601a556040805183815290517f122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b59181900360200190a15050565b60008060fb600160a060020a03168585856040516020018084600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140183805190602001908083835b602083106121575780518252601f199092019160209182019101612138565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061219f5780518252601f199092019160209182019101612180565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040526040518082805190602001908083835b602083106122045780518252601f1990920191602091820191016121e5565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114612264576040519150601f19603f3d011682016040523d82523d6000602084013e612269565b606091505b50909695505050505050565b6000818152600f60205260408120611f53906168b6565b600160a060020a031660009081526010602052604090206002015490565b6122b261453a565b15156122f6576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b600081116123385760405160e560020a62461bcd02815260040180806020018281038252602181526020018061863f6021913960400191505060405180910390fd5b600654811415612392576040805160e560020a62461bcd02815260206004820152601560248201527f517565756545787069727920756e6368616e6765640000000000000000000000604482015290519081900360640190fd5b60068190556040805182815290517f4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d79181900360200190a150565b6005545b90565b6123dc61453a565b1515612420576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b600160a060020a0381161515612480576040805160e560020a62461bcd02815260206004820152601460248201527f417070726f7665722063616e6e6f742062652030000000000000000000000000604482015290519081900360640190fd5b600854600160a060020a03828116911614156124e6576040805160e560020a62461bcd02815260206004820152601260248201527f417070726f76657220756e6368616e6765640000000000000000000000000000604482015290519081900360640190fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383169081179091556040517fa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e190600090a250565b6000612547614044565b612550836141a3565b101592915050565b6000611f5382612566616194565b616a50565b6007546009546125809163ffffffff616a8416565b421061286257600a5460145460009161259891616ae1565b90506060601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__6377b024799091846040518363ffffffff1660e060020a028152600401808381526020018281526020019250505060006040518083038186803b1580156125f957600080fd5b505af415801561260d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561263657600080fd5b81019080805164010000000081111561264e57600080fd5b8201602081018481111561266157600080fd5b815185602082028301116401000000008211171561267e57600080fd5b50909450600093505050505b8281101561285a57600082828151811015156126a257fe5b60209081029091018101516000818152600f90925260409091209091506126c881616af7565b156126ff5760405182907f88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a264390600090a25050612842565b60018101548154600160a060020a03166000908152600d602052604090205461272d9163ffffffff616a8416565b8154600160a060020a03166000908152600d602052604081209190915542600283015560185411156127d35760185460009061277090600163ffffffff616b1b16565b905082601760188381548110151561278457fe5b906000526020600020015481548110151561279b57fe5b60009182526020909120015560188054829081106127b557fe5b6000918252602082200155806127cc601882618404565b5050612809565b601780546001810182556000919091527fc624b66cc0138b8fabc209247f72d758e1cf3343756d543badbf24212bed8c15018290555b60408051428152905183917f3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543919081900360200190a250505b61285381600163ffffffff616a8416565b905061268a565b505042600955505b565b60018054810190819055336000908152600d60205260408120549091908281116128d8576040805160e560020a62461bcd02815260206004820152601360248201527f4e6f7468696e6720746f20776974686472617700000000000000000000000000604482015290519081900360640190fd5b3031811115612931576040805160e560020a62461bcd02815260206004820152601460248201527f496e636f6e73697374656e742062616c616e6365000000000000000000000000604482015290519081900360640190fd5b336000818152600d60205260408082208290555183156108fc0291849190818181858888f1935050505015801561296c573d6000803e3d6000fd5b50600192505060015481146129b9576040805160e560020a62461bcd02815260206004820152600e602482015260008051602061868e833981519152604482015290519081900360640190fd5b5090565b6129c561453a565b1515612a09576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b612a116183f2565b612a1a826167f1565b9050612a258161680b565b1515612a655760405160e560020a62461bcd0281526004018080602001828103825260278152602001806186ae6027913960400191505060405180910390fd5b60408051602081019091526019548152612a8690829063ffffffff61681e16565b15612adb576040805160e560020a62461bcd02815260206004820181905260248201527f50617274696369706174696f6e20626173656c696e6520756e6368616e676564604482015290519081900360640190fd5b80516019556040805183815290517f51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c59181900360200190a15050565b6000828152601160209081526040808320600160a060020a038516845260020190915290205460ff1692915050565b600c5481565b6011602052600090815260409020805460019091015460ff808316926101009004169083565b60006060600060f4600160a060020a0316846040516020018082805190602001908083835b60208310612bb65780518252601f199092019160209182019101612b97565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310612c195780518252601f199092019160209182019101612bfa565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114612c79576040519150601f19603f3d011682016040523d82523d6000602084013e612c7e565b606091505b5092509050801515612cc45760405160e560020a62461bcd02815260040180806020018281038252603881526020018061880c6038913960400191505060405180910390fd5b611f32826000616b5d565b60018054810190819055600090612ce461256b565b600080612cf18686616bc6565b915091506000612d0083616831565b90508015612e09576004826005811115612d1657fe5b148015612d275750612d27836168b6565b1515612d675760405160e560020a62461bcd02815260040180806020018281038252602e81526020018061896b602e913960400191505060405180910390fd5b8273__$930539a6eede91684063b1b5ba01037e9c$__63c67e7b4b90916040518263ffffffff1660e060020a0281526004018082815260200191505060006040518083038186803b158015612dbb57600080fd5b505af4158015612dcf573d6000803e3d6000fd5b50506040518992507f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f9150600090a2612e09838888616c71565b935050506001548114612e54576040805160e560020a62461bcd02815260206004820152600e602482015260008051602061868e833981519152604482015290519081900360640190fd5b5092915050565b60018054810190819055600090612e7061256b565b612e7985616d6b565b15612e87576000915061337b565b6000612e91616eca565b600160a060020a0316636642d594336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015612ee957600080fd5b505afa158015612efd573d6000803e3d6000fd5b505050506040513d6020811015612f1357600080fd5b5051600160a060020a0381166000908152601060205260409020805491925090612f3c90616d6b565b506000612f47616f9f565b600160a060020a03166330ec70f5846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015612f9f57600080fd5b505afa158015612fb3573d6000803e3d6000fd5b505050506040513d6020811015612fc957600080fd5b505190506000811161300f5760405160e560020a62461bcd028152600401808060200182810382526022815260200180618a296022913960400191505060405180910390fd5b604080517fbfc5163800000000000000000000000000000000000000000000000000000000815260126004820152602481018a9052905173__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163bfc51638916044808301926020929190829003018186803b15801561308157600080fd5b505af4158015613095573d6000803e3d6000fd5b505050506040513d60208110156130ab57600080fd5b505115156130ed5760405160e560020a62461bcd02815260040180806020018281038252602981526020018061887a6029913960400191505060405180910390fd5b8154158061319857508154604080517fbfc516380000000000000000000000000000000000000000000000000000000081526012600482015260248101929092525173__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163bfc51638916044808301926020929190829003018186803b15801561316a57600080fd5b505af415801561317e573d6000803e3d6000fd5b505050506040513d602081101561319457600080fd5b5051155b15156131d85760405160e560020a62461bcd02815260040180806020018281038252602b815260200180618b3e602b913960400191505060405180910390fd5b600061327382601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__637577759990918d6040518363ffffffff1660e060020a028152600401808381526020018281526020019250505060206040518083038186803b15801561323b57600080fd5b505af415801561324f573d6000803e3d6000fd5b505050506040513d602081101561326557600080fd5b50519063ffffffff616a8416565b604080517f239491ba00000000000000000000000000000000000000000000000000000000815260126004820152602481018c905260448101839052606481018b9052608481018a9052905191925073__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163239491ba9160a480820192600092909190829003018186803b1580156132fe57600080fd5b505af4158015613312573d6000803e3d6000fd5b50506040805180820182528c815260209081018690528c87556001870186905581518681529151600160a060020a03891694508d93507fd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d629281900390910190a360019550505050505b60015481146133c2576040805160e560020a62461bcd02815260206004820152600e602482015260008051602061868e833981519152604482015290519081900360640190fd5b509392505050565b60008115806133da5750600b5482115b156133e757506000611f37565b6133f0826158a1565b156133fd57506001611f37565b6000828152600f6020526040902061341c90600363ffffffff61704316565b9050611f37565b61342b61453a565b151561346f576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b6134776183f2565b613480826167f1565b905061348b8161680b565b15156134cb5760405160e560020a62461bcd0281526004018080602001828103825260278152602001806187e56027913960400191505060405180910390fd5b6040805160208101909152601b5481526134ec90829063ffffffff61681e16565b15613541576040805160e560020a62461bcd02815260206004820181905260248201527f426173656c696e652075706461746520666163746f7220756e6368616e676564604482015290519081900360640190fd5b8051601b556040805183815290517f8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f9181900360200190a15050565b6040805160208082018590528183018490528251808303840181526060928301938490528051600094859360fa939282918401908083835b602083106135d45780518252601f1990920191602091820191016135b5565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114613634576040519150601f19603f3d011682016040523d82523d6000602084013e613639565b606091505b509250905080151561367f5760405160e560020a62461bcd0281526004018080602001828103825260368152602001806189156036913960400191505060405180910390fd5b61368a826000616825565b95945050505050565b600854600090600160a060020a031633146136f8576040805160e560020a62461bcd02815260206004820152601760248201527f6d73672e73656e646572206e6f7420617070726f766572000000000000000000604482015290519081900360640190fd5b61370061256b565b60008061370d8585616bc6565b9150915061371a82616831565b151561372b57600092505050611f53565b613734826170e0565b15613789576040805160e560020a62461bcd02815260206004820152601960248201527f50726f706f73616c20616c726561647920617070726f76656400000000000000604482015290519081900360640190fd5b600281600581111561379757fe5b146137ec576040805160e560020a62461bcd02815260206004820152601e60248201527f50726f706f73616c206e6f7420696e20617070726f76616c2073746167650000604482015290519081900360640190fd5b60078201805460ff19166001179055613803616f9f565b600160a060020a03166330a61d596040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561383e57600080fd5b505afa158015613852573d6000803e3d6000fd5b505050506040513d602081101561386857600080fd5b5051600883015560405185907f28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d92990600090a2506001949350505050565b600160a060020a03821660009081526010602090815260408083208484526003908101909252822060018101548154849360ff909116908111156138e557fe5b92509250509250929050565b600160a060020a0381166000908152601060205260408120805482811580159061391f575061391f826158a1565b60028401546000908152600f60205260408120919250600361394883600363ffffffff61704316565b600581111561395357fe5b149050828061395f5750805b979650505050505050565b600d6020526000908152604090205481565b600061398661256b565b600c543410156139e0576040805160e560020a62461bcd02815260206004820152601160248201527f546f6f20736d616c6c206465706f736974000000000000000000000000000000604482015290519081900360640190fd5b600b546139f490600163ffffffff616a8416565b600b819055506000600f6000600b54815260200190815260200160002090508073__$930539a6eede91684063b1b5ba01037e9c$__633053123f90918e8e8e8e8e8e8e8e33346040518c63ffffffff1660e060020a028152600401808c81526020018060200180602001806020018060200187600160a060020a0316600160a060020a0316815260200186815260200185810385528f8f82818152602001925060200280828437600083820152601f01601f191690910186810385528d8152602090810191508e908e0280828437600083820152601f01601f191690910186810384528b815260200190508b8b80828437600083820152601f01601f19169091018681038352898152602090810191508a908a0280828437600081840152601f19601f8201169050808301925050509f5050505050505050505050505050505060006040518083038186803b158015613b4c57600080fd5b505af4158015613b60573d6000803e3d6000fd5b50505050613bac84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250859392505063ffffffff6170ea169050565b600b54604080517fd7a8acc10000000000000000000000000000000000000000000000000000000081526012600482015260248101929092525173__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163d7a8acc1916044808301926000929190829003018186803b158015613c2157600080fd5b505af4158015613c35573d6000803e3d6000fd5b5050600b546006840154604080519182523460208301524282820152513394509192507f1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe919081900360600190a35050600b549a9950505050505050505050565b613c9e61453a565b1515613ce2576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008111613d3a576040805160e560020a62461bcd02815260206004820152601e60248201527f4475726174696f6e206d757374206265206c6172676572207468616e20300000604482015290519081900360640190fd5b600554811415613d94576040805160e560020a62461bcd02815260206004820152601260248201527f4475726174696f6e20756e6368616e6765640000000000000000000000000000604482015290519081900360640190fd5b60058190556040805182815290517f7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d89181900360200190a150565b60006060600060f6600160a060020a0316846040516020018082805190602001908083835b60208310613e135780518252601f199092019160209182019101613df4565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310613e765780518252601f199092019160209182019101613e57565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114613ed6576040519150601f19603f3d011682016040523d82523d6000602084013e613edb565b606091505b5092509050801515612cc45760405160e560020a62461bcd028152600401808060200182810382526023815260200180618b696023913960400191505060405180910390fd5b60606017805480602002602001604051908101604052809291908181526020018280548015613f6f57602002820191906000526020600020905b815481526020019060010190808311613f5b575b5050505050905090565b6000818152600f60205260408120611f8181613f9c81600363ffffffff61704316565b617146565b613fa961453a565b1515613fed576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600061404f436162ac565b905090565b60075481565b6000818152600f60205260408120611f53906170e0565b600254600160a060020a031681565b61408861453a565b15156140cc576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b6000811161410e5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806189996026913960400191505060405180910390fd5b600754811415614168576040805160e560020a62461bcd02815260206004820152601a60248201527f646571756575654672657175656e637920756e6368616e676564000000000000604482015290519081900360640190fd5b60078190556040805182815290517f391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c349181900360200190a150565b600080806141af6142b6565b905060006141bb616eca565b905060005b828110156142ac5760006141d382611e26565b9050600083600160a060020a03166393c5c487836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561423057600080fd5b505afa158015614244573d6000803e3d6000fd5b505050506040513d602081101561425a57600080fd5b505190506142688883612b17565b8061427857506142788882612b17565b156142915761428e86600163ffffffff616a8416565b95505b506142a5905081600163ffffffff616a8416565b90506141c0565b5091949350505050565b60006060600060f9600160a060020a031643604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106143165780518252601f1990920191602091820191016142f7565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114614376576040519150601f19603f3d011682016040523d82523d6000602084013e61437b565b606091505b50925090508015156143c15760405160e560020a62461bcd0281526004018080602001828103825260358152602001806188e06035913960400191505060405180910390fd5b6143cc826000616825565b9250505090565b60006060600060f7600160a060020a0316846040516020018082805190602001908083835b602083106144175780518252601f1990920191602091820191016143f8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831061447a5780518252601f19909201916020918201910161445b565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146144da576040519150601f19603f3d011682016040523d82523d6000602084013e6144df565b606091505b5092509050801515611f275760405160e560020a62461bcd028152600401808060200182810382526031815260200180618b0d6031913960400191505060405180910390fd5b600054600160a060020a031690565b60065481565b60008054600160a060020a031661454f6171ab565b600160a060020a031614905090565b61456661453a565b15156145aa576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008111614602576040805160e560020a62461bcd02815260206004820181905260248201527f6d696e4465706f736974206d757374206265206c6172676572207468616e2030604482015290519081900360640190fd5b600c5481141561465c576040805160e560020a62461bcd02815260206004820152601960248201527f4d696e696d756d206465706f73697420756e6368616e67656400000000000000604482015290519081900360640190fd5b600c8190556040805182815290517fc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba819181900360200190a150565b6000611f816146a684846171af565b6172c6565b60006146b6826158a1565b151561470c576040805160e560020a62461bcd02815260206004820152601360248201527f50726f706f73616c206e6f742071756575656400000000000000000000000000604482015290519081900360640190fd5b604080517f757775990000000000000000000000000000000000000000000000000000000081526012600482015260248101849052905173__$23a392d0baa2c696c1f8abcbfee88c03a9$__916375777599916044808301926020929190829003018186803b15801561477e57600080fd5b505af4158015614792573d6000803e3d6000fd5b505050506040513d60208110156147a857600080fd5b505192915050565b6147b861453a565b15156147fc576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008111614854576040805160e560020a62461bcd02815260206004820152601e60248201527f4475726174696f6e206d757374206265206c6172676572207468616e20300000604482015290519081900360640190fd5b6003548114156148ae576040805160e560020a62461bcd02815260206004820152601260248201527f4475726174696f6e20756e6368616e6765640000000000000000000000000000604482015290519081900360640190fd5b60038190556040805182815290517fbc44c003a66b841b48c220869efb738b265af305649ac8bafe8efe0c8130e3139181900360200190a150565b600061404f43612558565b60006060600060f9600160a060020a031684604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106149545780518252601f199092019160209182019101614935565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146149b4576040519150601f19603f3d011682016040523d82523d6000602084013e6149b9565b606091505b5092509050801515611f275760405160e560020a62461bcd02815260040180806020018281038252602e815260200180618660602e913960400191505060405180910390fd5b600081815260116020526040902054819060ff1615614a56576040805160e560020a62461bcd028152602060048201526017602482015260008051602061894b833981519152604482015290519081900360640190fd5b614a5f8261253d565b1515614a9f5760405160e560020a62461bcd0281526004018080602001828103825260298152602001806189e06029913960400191505060405180910390fd5b6000614aa96148e9565b6000848152601160205260409020600101549091508111614afe5760405160e560020a62461bcd028152600401808060200182810382526026815260200180618a726026913960400191505060405180910390fd5b60008381526011602052604080822060010183905551829185917f6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e09190a3505050565b614b4961453a565b1515614b8d576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b600160a060020a0381161515614bed576040805160e560020a62461bcd02815260206004820181905260248201527f43616e6e6f7420726567697374657220746865206e756c6c2061646472657373604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383169081179091556040517f27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b90600090a250565b6018805482908110614c5257fe5b600091825260209091200154905081565b60045490565b6017805482908110614c5257fe5b60018054810190819055600090614c8c61256b565b6000614c96616eca565b600160a060020a0316636642d594336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015614cee57600080fd5b505afa158015614d02573d6000803e3d6000fd5b505050506040513d6020811015614d1857600080fd5b5051600160a060020a0381166000908152601060205260409020805491925090801515614d8f576040805160e560020a62461bcd02815260206004820181905260248201527f4163636f756e7420686173206e6f20686973746f726963616c207570766f7465604482015290519081900360640190fd5b614d9881616d6b565b50604080517fbfc516380000000000000000000000000000000000000000000000000000000081526012600482015260248101839052905173__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163bfc51638916044808301926020929190829003018186803b158015614e0b57600080fd5b505af4158015614e1f573d6000803e3d6000fd5b505050506040513d6020811015614e3557600080fd5b505115614fae57601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__63239491ba909183614efb8660000160010154601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__63757775999091896040518363ffffffff1660e060020a028152600401808381526020018281526020019250505060206040518083038186803b158015614ec357600080fd5b505af4158015614ed7573d6000803e3d6000fd5b505050506040513d6020811015614eed57600080fd5b50519063ffffffff616b1b16565b8b8b6040518663ffffffff1660e060020a028152600401808681526020018581526020018481526020018381526020018281526020019550505050505060006040518083038186803b158015614f5057600080fd5b505af4158015614f64573d6000803e3d6000fd5b50505060018301546040805191825251600160a060020a038616925083917f7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114919081900360200190a35b5060408051808201909152600080825260209091018190528082556001918201559250506001548114612e54576040805160e560020a62461bcd02815260206004820152600e602482015260008051602061868e833981519152604482015290519081900360640190fd5b600081815260116020526040902054819060ff1615615070576040805160e560020a62461bcd028152602060048201526017602482015260008051602061894b833981519152604482015290519081900360640190fd5b600854600160a060020a031633146150d2576040805160e560020a62461bcd02815260206004820152601760248201527f6d73672e73656e646572206e6f7420617070726f766572000000000000000000604482015290519081900360640190fd5b600082815260116020526040808220805461ff0019166101001790555183917f36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f91a25050565b600081815260116020526040902054819060ff161561516f576040805160e560020a62461bcd028152602060048201526017602482015260008051602061894b833981519152604482015290519081900360640190fd5b6000828152601160209081526040808320338085526002909101835292819020805460ff1916600117905580519283525184927ff6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f92908290030190a25050565b60145490565b600180548101908190556000906151ea61256b565b6000806151f78787616bc6565b9150915061520482616831565b15156152155760009350505061337b565b600061521f616eca565b600160a060020a0316636642d594336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561527757600080fd5b505afa15801561528b573d6000803e3d6000fd5b505050506040513d60208110156152a157600080fd5b5051600160a060020a03811660009081526010602052604081209192506152c6616f9f565b600160a060020a03166330ec70f5846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561531e57600080fd5b505afa158015615332573d6000803e3d6000fd5b505050506040513d602081101561534857600080fd5b50519050615355856170e0565b15156153ab576040805160e560020a62461bcd02815260206004820152601560248201527f50726f706f73616c206e6f7420617070726f7665640000000000000000000000604482015290519081900360640190fd5b60038460058111156153b957fe5b1461540e576040805160e560020a62461bcd02815260206004820152601860248201527f496e636f72726563742070726f706f73616c2073746174650000000000000000604482015290519081900360640190fd5b600088600381111561541c57fe5b1415615472576040805160e560020a62461bcd02815260206004820152601060248201527f566f74652076616c756520756e73657400000000000000000000000000000000604482015290519081900360640190fd5b600081116154ca576040805160e560020a62461bcd02815260206004820152601160248201527f566f74657220776569676874207a65726f000000000000000000000000000000604482015290519081900360640190fd5b60008260030160008b815260200190815260200160002090508573__$930539a6eede91684063b1b5ba01037e9c$__63b05cd27f90918360020154858f86600101541461551857600061551e565b855460ff165b8e6040518663ffffffff1660e060020a0281526004018086815260200185815260200184815260200183600381111561555357fe5b60ff16815260200182600381111561556757fe5b60ff1681526020019550505050505060006040518083038186803b15801561558e57600080fd5b505af41580156155a2573d6000803e3d6000fd5b505050506155ae616f9f565b600160a060020a03166330a61d596040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156155e957600080fd5b505afa1580156155fd573d6000803e3d6000fd5b505050506040513d602081101561561357600080fd5b505160088701556040805160608101909152808a600381111561563257fe5b81526020018c8152602001838152508360030160008c815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600381111561567a57fe5b021790555060208201516001820155604090910151600291820155838101549087015411156156ab57600283018b90555b83600160a060020a03168b7ff3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df8b60038111156156e357fe5b60408051918252602082018790528051918290030190a36001975050505050505060015481146133c2576040805160e560020a62461bcd02815260206004820152600e602482015260008051602061868e833981519152604482015290519081900360640190fd5b60095481565b6000818152600f60205260408120611f5390616af7565b60005474010000000000000000000000000000000000000000900460ff16156157db576040805160e560020a62461bcd02815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a656400000000604482015290519081900360640190fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055615819336172ca565b6158228d614b41565b61582b8c6123d4565b6158348b615a81565b61583d8a61455e565b615846896122aa565b61584f88614080565b615858876147b0565b61586186615bea565b61586a85613c96565b615873846129bd565b61587c83611fa9565b61588582613423565b61588e81611c95565b5050426009555050505050505050505050565b6000601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__63bfc516389091846040518363ffffffff1660e060020a028152600401808381526020018281526020019250505060206040518083038186803b15801561590057600080fd5b505af4158015615914573d6000803e3d6000fd5b505050506040513d602081101561592a57600080fd5b50518015611f53575061593c82615751565b1592915050565b6000806000806060615966600f600088815260200190815260200160002061737c565b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281529183918301828280156159ec5780601f106159c1576101008083540402835291602001916159ec565b820191906000526020600020905b8154815290600101906020018083116159cf57829003601f168201915b505050505090509450945094509450945091939590929450565b60408051602081019091526019548152600090819081908190615a28906172c6565b6040805160208101909152601a548152615a41906172c6565b6040805160208101909152601b548152615a5a906172c6565b6040805160208101909152601c548152615a73906172c6565b935093509350935090919293565b615a8961453a565b1515615acd576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008111615b0f5760405160e560020a62461bcd02815260040180806020018281038252602c8152602001806186d5602c913960400191505060405180910390fd5b600a54811415615b69576040805160e560020a62461bcd02815260206004820152601d60248201527f4e756d626572206f662070726f706f73616c7320756e6368616e676564000000604482015290519081900360640190fd5b600a8190556040805182815290517f85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a69181900360200190a150565b600080615baf618428565b505050600160a060020a0316600090815260106020908152604091829020825180840190935280548084526001909101549290910182905291565b615bf261453a565b1515615c36576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b60008111615c8e576040805160e560020a62461bcd02815260206004820152601e60248201527f4475726174696f6e206d757374206265206c6172676572207468616e20300000604482015290519081900360640190fd5b600454811415615ce8576040805160e560020a62461bcd02815260206004820152601260248201527f4475726174696f6e20756e6368616e6765640000000000000000000000000000604482015290519081900360640190fd5b60048190556040805182815290517f90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed69181900360200190a150565b6000898989898989898989604051602001808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600083820152601f01601f191690910186810385528c8152602090810191508d908d0280828437600083820152601f01601f191690910186810384528a815260200190508a8a80828437600083820152601f01601f19169091018681038352888152602090810191508990890280828437600081840152601f19601f8201169050808301925050509d50505050505050505050505050506040516020818303038152906040528051906020012090506000806000615e1f84611def565b919450925090508115615e6a576040805160e560020a62461bcd028152602060048201526017602482015260008051602061894b833981519152604482015290519081900360640190fd5b821515615ec1576040805160e560020a62461bcd02815260206004820152601360248201527f686f74666978206e6f7420617070726f76656400000000000000000000000000604482015290519081900360640190fd5b615ec96148e9565b8114615f095760405160e560020a62461bcd0281526004018080602001828103825260268152602001806186196026913960400191505060405180910390fd5b6160186160138e8e80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092018290525033935091506173a69050565b617570565b600084815260116020526040808220805460ff191660011790555185917f708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b6391a250505050505050505050505050565b6000828152600f602052604080822081517fe6a5192f0000000000000000000000000000000000000000000000000000000081526004810191909152602481018490529051829160609173__$930539a6eede91684063b1b5ba01037e9c$__9163e6a5192f9160448083019287929190829003018186803b1580156160eb57600080fd5b505af41580156160ff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052606081101561612857600080fd5b815160208301516040840180519294919382019264010000000081111561614e57600080fd5b8201602081018481111561616157600080fd5b815164010000000081118282018710171561617b57600080fd5b50959b949a509850929650505050505050565b600b5481565b604080516000808252602082019283905281519092606092849260f89290819081908082805b602083106161d95780518252601f1990920191602091820191016161ba565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114616239576040519150601f19603f3d011682016040523d82523d6000602084013e61623e565b606091505b50925090508015156143c15760405160e560020a62461bcd028152600401808060200182810382526025815260200180618a986025913960400191505060405180910390fd5b6000818152600f602052604081208190819061629f9061757d565b9250925092509193909250565b6000611f5360036162de60026162d260026162c6886148f4565b9063ffffffff61759116565b9063ffffffff616a8416565b9063ffffffff6175f116565b60008086158015906162fb57508415155b1515616351576040805160e560020a62461bcd02815260206004820152601560248201527f612064656e6f6d696e61746f72206973207a65726f0000000000000000000000604482015290519081900360640190fd5b6000806000606060fc600160a060020a03168c8c8c8c8c8c6040516020018087815260200186815260200185815260200184815260200183815260200182815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b602083106163da5780518252601f1990920191602091820191016163bb565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d806000811461643a576040519150601f19603f3d011682016040523d82523d6000602084013e61643f565b606091505b5090925090508115156164865760405160e560020a62461bcd028152600401808060200182810382526027815260200180618a4b6027913960400191505060405180910390fd5b616491816000616825565b935061649e816020616825565b939c939b50929950505050505050505050565b6164b961453a565b15156164fd576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b600160a060020a038316151561655d576040805160e560020a62461bcd02815260206004820152601a60248201527f44657374696e6174696f6e2063616e6e6f74206265207a65726f000000000000604482015290519081900360640190fd5b6969e10de76676d080000081118015616580575061657c6146a6617633565b8111155b15156165c05760405160e560020a62461bcd0281526004018080602001828103825260488152602001806187776048913960600191505060405180910390fd5b600160e060020a0319821615156165fb576165da816167f1565b600160a060020a0384166000908152600e6020526040902090519055616639565b616604816167f1565b600160a060020a0384166000908152600e60209081526040808320600160e060020a0319871684526001019091529020905190555b604080518281529051600160e060020a0319841691600160a060020a038616917f60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef9181900360200190a3505050565b61669061453a565b15156166d4576040805160e560020a62461bcd0281526020600482018190526024820152600080516020618a09833981519152604482015290519081900360640190fd5b6166dd816172ca565b50565b60006060600060f5600160a060020a031684604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106167405780518252601f199092019160209182019101616721565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146167a0576040519150601f19603f3d011682016040523d82523d6000602084013e6167a5565b606091505b5092509050801515612cc45760405160e560020a62461bcd02815260040180806020018281038252602c815260200180618ae1602c913960400191505060405180910390fd5b60035490565b6167f96183f2565b50604080516020810190915290815290565b6000611f5382616819617633565b617657565b5190511490565b6000611f818383616b5d565b6002015460001090565b60175460009082106168815760405160e560020a62461bcd02815260040180806020018281038252602b815260200180618b8c602b913960400191505060405180910390fd5b61688a84616831565b80156168ae5750826017838154811015156168a157fe5b9060005260206000200154145b949350505050565b60006168c06183f2565b6040805160208181018352601c548252825190810190925260195482526168fd916168f09163ffffffff61765f16565b849063ffffffff6179f116565b905060005b6006840154811015616a465760006169c6856006018381548110151561692457fe5b600091825260209182902060026003909202018101805460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835291929091908301828280156169bc5780601f10616991576101008083540402835291602001916169bc565b820191906000526020600020905b81548152906001019060200180831161699f57829003601f168201915b5050505050617aa7565b90506169d06183f2565b616a0786600601848154811015156169e457fe5b6000918252602090912060016003909202010154600160a060020a0316836171af565b9050616a19848263ffffffff61765716565b15616a2b576000945050505050611f37565b50616a3f905081600163ffffffff616a8416565b9050616902565b5060019392505050565b6000808284811515616a5e57fe5b0490508284811515616a6c57fe5b061515616a7a579050611f53565b6001019050611f53565b600082820183811015611f81576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000818310616af05781611f81565b5090919050565b6000616b126006548360020154616a8490919063ffffffff16565b42101592915050565b6000611f8183836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250617c22565b600081602001835110151515616bbd576040805160e560020a62461bcd02815260206004820152601460248201527f736c6963696e67206f7574206f662072616e6765000000000000000000000000604482015290519081900360640190fd5b50016020015190565b6000828152600f602052604081208190616be181868661683b565b1515616c37576040805160e560020a62461bcd02815260206004820152601560248201527f50726f706f73616c206e6f742064657175657565640000000000000000000000604482015290519081900360640190fd5b6000616c4a82600363ffffffff61704316565b9050616c568282617146565b15616c6657616c66828787616c71565b909590945092505050565b616c7a836170e0565b8015616c8a575060008360080154115b15616c9857616c9883617cbc565b6000601782815481101515616ca957fe5b60009182526020808320909101929092556018805460018181019092557fb13d2d76d1f4b7be834882e410b3e3a8afaf69f83600ae24db354391d2378d2e01849055848252600f90925260408120805473ffffffffffffffffffffffffffffffffffffffff1916815591820181905560028201819055600382018190556004820181905560058201819055616d41600683018261843f565b60078201805460ff19169055600060088301819055616d64906009840190618460565b5050505050565b600080601273__$23a392d0baa2c696c1f8abcbfee88c03a9$__63bfc516389091856040518363ffffffff1660e060020a028152600401808381526020018281526020019250505060206040518083038186803b158015616dcb57600080fd5b505af4158015616ddf573d6000803e3d6000fd5b505050506040513d6020811015616df557600080fd5b50518015616e075750616e0783615751565b90508015611f5357604080517feed5f7be0000000000000000000000000000000000000000000000000000000081526012600482015260248101859052905173__$23a392d0baa2c696c1f8abcbfee88c03a9$__9163eed5f7be916044808301926000929190829003018186803b158015616e8157600080fd5b505af4158015616e95573d6000803e3d6000fd5b50506040518592507f88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a26439150600090a292915050565b600254604080517f4163636f756e747300000000000000000000000000000000000000000000000060208083019190915282518083036008018152602883018085528151918301919091207fdcf0aaed00000000000000000000000000000000000000000000000000000000909152602c8301529151600093600160a060020a03169263dcf0aaed92604c8082019391829003018186803b158015616f6e57600080fd5b505afa158015616f82573d6000803e3d6000fd5b505050506040513d6020811015616f9857600080fd5b5051905090565b600254604080517f4c6f636b6564476f6c64000000000000000000000000000000000000000000006020808301919091528251808303600a018152602a83018085528151918301919091207fdcf0aaed00000000000000000000000000000000000000000000000000000000909152602e8301529151600093600160a060020a03169263dcf0aaed92604e8082019391829003018186803b158015616f6e57600080fd5b60008061707183600201546162d285600101546162d287600001548960020154616a8490919063ffffffff16565b9050428111617084576005915050611f53565b600283015461709a90829063ffffffff616b1b16565b90504281116170ad576004915050611f53565b60018301546170c390829063ffffffff616b1b16565b90504281116170d6576003915050611f53565b5060029392505050565b6007015460ff1690565b8051151561712c5760405160e560020a62461bcd0281526004018080602001828103825260298152602001806187276029913960400191505060405180910390fd5b805161714190600984019060208401906184a4565b505050565b6000600482600581111561715657fe5b118061717e5750600382600581111561716b57fe5b11801561717e575061717c836168b6565b155b80611f815750600282600581111561719257fe5b118015611f8157506171a3836170e0565b159392505050565b3390565b6171b76183f2565b6171bf6183f2565b6171d26969e10de76676d08000006167f1565b600160a060020a0385166000908152600e60209081526040808320600160e060020a03198816845260010182529182902082519182019092529054815290915061721b906172c6565b156172635750600160a060020a0383166000908152600e60209081526040808320600160e060020a031986168452600101825291829020825191820190925290548152611f81565b600160a060020a0384166000908152600e6020908152604091829020825191820190925290548152617294906172c6565b15611f8157505050600160a060020a03166000908152600e602090815260409182902082519182019092529054815290565b5190565b600160a060020a03811615156173145760405160e560020a62461bcd0281526004018080602001828103825260268152602001806187016026913960400191505060405180910390fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b8054600182015460028301546006840154600160a060020a03909316926009850191939590929450565b6173ae61851e565b855187511480156173c0575083518651145b1515617416576040805160e560020a62461bcd02815260206004820152601560248201527f4172726179206c656e677468206d69736d617463680000000000000000000000604482015290519081900360640190fd5b865161742061851e565b600160a060020a038516815260208082018590524260408084019190915280518481528483028101909201905260009083801561747757816020015b617464618575565b81526020019060019003908161745c5790505b50608083015260005b83811015617561576060604051908101604052808c838151811015156174a257fe5b9060200190602002015181526020018b838151811015156174bf57fe5b90602001906020020151600160a060020a03168152602001617501848b858151811015156174e957fe5b602090810290910101518d919063ffffffff617dea16565b9052608084015180518390811061751457fe5b6020908102909101015287516175479089908390811061753057fe5b60209081029091010151839063ffffffff616a8416565b915061755a81600163ffffffff616a8416565b9050617480565b50909998505050505050505050565b6166dd8160800151617e6c565b600381015460048201546005909201549092565b60008215156175a257506000611f53565b8282028284828115156175b157fe5b0414611f815760405160e560020a62461bcd0281526004018080602001828103825260218152602001806189bf6021913960400191505060405180910390fd5b6000611f8183836040805190810160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250617f60565b61763b6183f2565b50604080516020810190915269d3c21bcecceda1000000815290565b519051111590565b6176676183f2565b8251158061767457508151155b1561768e5750604080516020810190915260008152611f53565b815169d3c21bcecceda100000014156176a8575081611f53565b825169d3c21bcecceda100000014156176c2575080611f53565b600069d3c21bcecceda10000006176d885617fcd565b518115156176e257fe5b04905060006176f085617ff8565b519050600069d3c21bcecceda100000061770986617fcd565b5181151561771357fe5b049050600061772186617ff8565b519050838202841561778f5782858281151561773957fe5b041461778f576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207831793120646574656374656400000000000000000000604482015290519081900360640190fd5b69d3c21bcecceda10000008102811561780e5769d3c21bcecceda100000082828115156177b857fe5b041461780e576040805160e560020a62461bcd02815260206004820152601f60248201527f6f766572666c6f772078317931202a2066697865643120646574656374656400604482015290519081900360640190fd5b905080848402851561787c5784868281151561782657fe5b041461787c576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207832793120646574656374656400000000000000000000604482015290519081900360640190fd5b86840287156178e75784888281151561789157fe5b04146178e7576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207831793220646574656374656400000000000000000000604482015290519081900360640190fd5b6178ef618024565b878115156178f957fe5b049650617904618024565b8581151561790e57fe5b049450868502871561797c5785888281151561792657fe5b041461797c576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207832793220646574656374656400000000000000000000604482015290519081900360640190fd5b6179846183f2565b5060408051602081810183528782528251908101909252848252906179aa90829061802d565b90506179c5816020604051908101604052808681525061802d565b90506179e0816020604051908101604052808581525061802d565b9d9c50505050505050505050505050565b6179f96183f2565b6003830154801515617a1757617a0f60006180a9565b915050611f53565b60048401546005850154600090617a38906162d2858563ffffffff616a8416565b90506000617a61617a5c617a4f89600801546180a9565b889063ffffffff61765f16565b61811a565b905081811115617a8e57617a8b617a7e828463ffffffff616b1b16565b849063ffffffff616a8416565b92505b61395f84617aa2818663ffffffff616a8416565b61812b565b60006018826003815181101515617aba57fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908060020a82049150506010836002815181101515617b2957fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908060020a82049150506008846001815181101515617b9857fe5b602091010151855160029290920a60f860020a918290049091027fff0000000000000000000000000000000000000000000000000000000000000016049085906000908110617be357fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161717179050919050565b60008184841115617cb45760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015617c79578181015183820152602001617c61565b50505050905090810190601f168015617ca65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b617cc46183f2565b617ccd82618163565b9050617cd76183f2565b6040805160208101909152601b548152617cf890839063ffffffff61765f16565b9050617d026183f2565b6040805160208101909152601b548152617d4a90617d2e90617d22617633565b9063ffffffff6181a216565b604080516020810190915260195481529063ffffffff61765f16565b9050617d5c828263ffffffff61802d16565b5160198190556040805160208181018352601a5482528251908101909252918152617d8c9163ffffffff61821d16565b15617d9857601a546019555b604080516020810190915260195481527f51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c590617dd3906172c6565b60408051918252519081900360200190a150505050565b6060818301845110151515617dfe57600080fd5b606082158015617e1957604051915060208201604052617e63565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015617e52578051835260209283019201617e3a565b5050858452601f01601f1916604052505b50949350505050565b60005b8151811015617f5c57617eee8282815181101515617e8957fe5b90602001906020020151602001518383815181101515617ea557fe5b60209081029091010151518451859085908110617ebe57fe5b9060200190602002015160400151518585815181101515617edb57fe5b9060200190602002015160400151618224565b1515617f44576040805160e560020a62461bcd02815260206004820152601960248201527f50726f706f73616c20657865637574696f6e206661696c656400000000000000604482015290519081900360640190fd5b617f5581600163ffffffff616a8416565b9050617e6f565b5050565b600081818411617fb55760405160e560020a62461bcd02815260040180806020018281038252838181518152602001915080519060200190808383600083811015617c79578181015183820152602001617c61565b5060008385811515617fc357fe5b0495945050505050565b617fd56183f2565b506040805160208101909152905169d3c21bcecceda10000009081900402815290565b6180006183f2565b506040805160208101909152905169d3c21bcecceda1000000808204029003815290565b64e8d4a5100090565b6180356183f2565b8151835190810190811015618094576040805160e560020a62461bcd02815260206004820152601560248201527f616464206f766572666c6f772064657465637465640000000000000000000000604482015290519081900360640190fd5b60408051602081019091529081529392505050565b6180b16183f2565b6180b96182af565b8211156180fa5760405160e560020a62461bcd0281526004018080602001828103825260368152602001806188446036913960400191505060405180910390fd5b50604080516020810190915269d3c21bcecceda100000082028152919050565b5169d3c21bcecceda1000000900490565b6181336183f2565b61813b6183f2565b618144846180a9565b905061814e6183f2565b618157846180a9565b905061368a82826182ca565b61816b6183f2565b6005820154600483015460038401546000926181929290916162d29163ffffffff616a8416565b9050611f8181846008015461812b565b6181aa6183f2565b815183511015618204576040805160e560020a62461bcd02815260206004820152601f60248201527f737562737472616374696f6e20756e646572666c6f7720646574656374656400604482015290519081900360640190fd5b5060408051602081019091528151835103815292915050565b5190511090565b600080600084111561828f57618239866183b9565b151561828f576040805160e560020a62461bcd02815260206004820152601860248201527f496e76616c696420636f6e747261637420616464726573730000000000000000604482015290519081900360640190fd5b6040516020840160008287838a8c6187965a03f198975050505050505050565b7601357c299a88ea76a58924d52ce4f26a85af186c2b9e7490565b6182d26183f2565b8151151561832a576040805160e560020a62461bcd02815260206004820152601160248201527f63616e2774206469766964652062792030000000000000000000000000000000604482015290519081900360640190fd5b825169d3c21bcecceda10000008181029190820414618393576040805160e560020a62461bcd02815260206004820152601260248201527f6f766572666c6f77206174206469766964650000000000000000000000000000604482015290519081900360640190fd5b6020604051908101604052808460000151838115156183ae57fe5b049052949350505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590611f32575050151592915050565b60408051602081019091526000815290565b81548183558181111561714157600083815260209020617141918101908301618594565b604080518082019091526000808252602082015290565b50805460008255600302906000526020600020908101906166dd91906185ae565b50805460018160011615610100020316600290046000825580601f1061848657506166dd565b601f0160209004906000526020600020908101906166dd9190618594565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106184e557805160ff1916838001178555618512565b82800160010185558215618512579182015b828111156185125782518255916020019190600101906184f7565b506129b9929150618594565b610140604051908101604052806000600160a060020a0316815260200160008152602001600081526020016185516185f6565b81526020016060815260200160001515815260200160008152602001606081525090565b6040805160608181018352600080835260208301529181019190915290565b6123d191905b808211156129b9576000815560010161859a565b6123d191905b808211156129b957600080825560018201805473ffffffffffffffffffffffffffffffffffffffff191690556185ed6002830182618460565b506003016185b4565b606060405190810160405280600081526020016000815260200160008152509056fe686f74666978206d75737420626520707265706172656420666f7220746869732065706f63685175657565457870697279206d757374206265206c6172676572207468616e20306572726f722063616c6c696e67206e756d62657256616c696461746f7273496e53657420707265636f6d70696c657265656e7472616e742063616c6c00000000000000000000000000000000000050617274696369706174696f6e20626173656c696e652067726561746572207468616e206f6e654e756d626572206f662070726f706f73616c73206d757374206265206c6172676572207468616e207a65726f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734465736372697074696f6e2075726c206d7573742068617665206e6f6e2d7a65726f206c656e677468426173656c696e652071756f72756d20666163746f722067726561746572207468616e206f6e655468726573686f6c642068617320746f2062652067726561746572207468616e206d616a6f7269747920616e64206e6f742067726561746572207468616e20756e616e696d69747950617274696369706174696f6e20626173656c696e6520666c6f6f7220756e6368616e676564426173656c696e652075706461746520666163746f722067726561746572207468616e206f6e656572726f722063616c6c696e672067657456657269666965645365616c4269746d617046726f6d48656164657220707265636f6d70696c6563616e277420637265617465206669786964697479206e756d626572206c6172676572207468616e206d61784e65774669786564282963616e6e6f74207570766f746520612070726f706f73616c206e6f7420696e207468652071756575656572726f722063616c6c696e672076616c696461746f725369676e65724164647265737346726f6d43757272656e7453657420707265636f6d70696c656572726f722063616c6c696e67206e756d62657256616c696461746f7273496e43757272656e7453657420707265636f6d70696c656572726f722063616c6c696e672076616c696461746f725369676e65724164647265737346726f6d53657420707265636f6d70696c65686f7466697820616c726561647920657865637574656400000000000000000050726f706f73616c206e6f7420696e20657865637574696f6e207374616765206f72206e6f742070617373696e67646571756575654672657175656e6379206d757374206265206c6172676572207468616e2030536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77686f74666978206e6f742077686974656c69737465642062792032662b312076616c696461746f72734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657263616e6e6f74207570766f746520776974686f7574206c6f636b696e6720676f6c646572726f722063616c6c696e67206672616374696f6e4d756c45787020707265636f6d70696c65686f7466697820616c726561647920707265706172656420666f7220746869732065706f63686572726f722063616c6c696e672067657445706f636853697a6520707265636f6d70696c6550617274696369706174696f6e20666c6f6f722067726561746572207468616e206f6e656572726f722063616c6c696e6720676574506172656e745365616c4269746d617020707265636f6d70696c656572726f722063616c6c696e6720676574426c6f636b4e756d62657246726f6d48656164657220707265636f6d70696c6563616e6e6f74207570766f7465206d6f7265207468616e206f6e65207175657565642070726f706f73616c6572726f722063616c6c696e67206861736848656164657220707265636f6d70696c6550726f766964656420696e6465782067726561746572207468616e2064657175657565206c656e6774682ea165627a7a7230582067794f847a83ba05997d0a1380d56833dd44971b991040c5bb88696d9978ee9e0029"

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() constant returns(address)
func (_Governance *GovernanceCaller) Approver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "approver")
	return *ret0, err
}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() constant returns(address)
func (_Governance *GovernanceSession) Approver() (common.Address, error) {
	return _Governance.Contract.Approver(&_Governance.CallOpts)
}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() constant returns(address)
func (_Governance *GovernanceCallerSession) Approver() (common.Address, error) {
	return _Governance.Contract.Approver(&_Governance.CallOpts)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Governance *GovernanceCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "checkProofOfPossession", sender, blsKey, blsPop)
	return *ret0, err
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Governance *GovernanceSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Governance.Contract.CheckProofOfPossession(&_Governance.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Governance *GovernanceCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Governance.Contract.CheckProofOfPossession(&_Governance.CallOpts, sender, blsKey, blsPop)
}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() constant returns(uint256)
func (_Governance *GovernanceCaller) ConcurrentProposals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "concurrentProposals")
	return *ret0, err
}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() constant returns(uint256)
func (_Governance *GovernanceSession) ConcurrentProposals() (*big.Int, error) {
	return _Governance.Contract.ConcurrentProposals(&_Governance.CallOpts)
}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() constant returns(uint256)
func (_Governance *GovernanceCallerSession) ConcurrentProposals() (*big.Int, error) {
	return _Governance.Contract.ConcurrentProposals(&_Governance.CallOpts)
}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() constant returns(uint256)
func (_Governance *GovernanceCaller) DequeueFrequency(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dequeueFrequency")
	return *ret0, err
}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() constant returns(uint256)
func (_Governance *GovernanceSession) DequeueFrequency() (*big.Int, error) {
	return _Governance.Contract.DequeueFrequency(&_Governance.CallOpts)
}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DequeueFrequency() (*big.Int, error) {
	return _Governance.Contract.DequeueFrequency(&_Governance.CallOpts)
}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) Dequeued(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dequeued", arg0)
	return *ret0, err
}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) Dequeued(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Dequeued(&_Governance.CallOpts, arg0)
}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) Dequeued(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Dequeued(&_Governance.CallOpts, arg0)
}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) EmptyIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "emptyIndices", arg0)
	return *ret0, err
}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) EmptyIndices(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.EmptyIndices(&_Governance.CallOpts, arg0)
}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) EmptyIndices(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.EmptyIndices(&_Governance.CallOpts, arg0)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Governance *GovernanceCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Governance.contract.Call(opts, out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
	return *ret0, *ret1, err
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Governance *GovernanceSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.FractionMulExp(&_Governance.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Governance *GovernanceCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.FractionMulExp(&_Governance.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetApprovalStageDuration is a free data retrieval call binding the contract method 0xffea74c0.
//
// Solidity: function getApprovalStageDuration() constant returns(uint256)
func (_Governance *GovernanceCaller) GetApprovalStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getApprovalStageDuration")
	return *ret0, err
}

// GetApprovalStageDuration is a free data retrieval call binding the contract method 0xffea74c0.
//
// Solidity: function getApprovalStageDuration() constant returns(uint256)
func (_Governance *GovernanceSession) GetApprovalStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetApprovalStageDuration(&_Governance.CallOpts)
}

// GetApprovalStageDuration is a free data retrieval call binding the contract method 0xffea74c0.
//
// Solidity: function getApprovalStageDuration() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetApprovalStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetApprovalStageDuration(&_Governance.CallOpts)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Governance *GovernanceCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getBlockNumberFromHeader", header)
	return *ret0, err
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Governance *GovernanceSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Governance.Contract.GetBlockNumberFromHeader(&_Governance.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Governance.Contract.GetBlockNumberFromHeader(&_Governance.CallOpts, header)
}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) constant returns(uint256)
func (_Governance *GovernanceCaller) GetConstitution(opts *bind.CallOpts, destination common.Address, functionId [4]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getConstitution", destination, functionId)
	return *ret0, err
}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) constant returns(uint256)
func (_Governance *GovernanceSession) GetConstitution(destination common.Address, functionId [4]byte) (*big.Int, error) {
	return _Governance.Contract.GetConstitution(&_Governance.CallOpts, destination, functionId)
}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetConstitution(destination common.Address, functionId [4]byte) (*big.Int, error) {
	return _Governance.Contract.GetConstitution(&_Governance.CallOpts, destination, functionId)
}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() constant returns(uint256[])
func (_Governance *GovernanceCaller) GetDequeue(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getDequeue")
	return *ret0, err
}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() constant returns(uint256[])
func (_Governance *GovernanceSession) GetDequeue() ([]*big.Int, error) {
	return _Governance.Contract.GetDequeue(&_Governance.CallOpts)
}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() constant returns(uint256[])
func (_Governance *GovernanceCallerSession) GetDequeue() ([]*big.Int, error) {
	return _Governance.Contract.GetDequeue(&_Governance.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Governance *GovernanceCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getEpochNumber")
	return *ret0, err
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Governance *GovernanceSession) GetEpochNumber() (*big.Int, error) {
	return _Governance.Contract.GetEpochNumber(&_Governance.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Governance.Contract.GetEpochNumber(&_Governance.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getEpochNumberOfBlock", blockNumber)
	return *ret0, err
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetEpochNumberOfBlock(&_Governance.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetEpochNumberOfBlock(&_Governance.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Governance *GovernanceCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getEpochSize")
	return *ret0, err
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Governance *GovernanceSession) GetEpochSize() (*big.Int, error) {
	return _Governance.Contract.GetEpochSize(&_Governance.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochSize() (*big.Int, error) {
	return _Governance.Contract.GetEpochSize(&_Governance.CallOpts)
}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() constant returns(uint256)
func (_Governance *GovernanceCaller) GetExecutionStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getExecutionStageDuration")
	return *ret0, err
}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() constant returns(uint256)
func (_Governance *GovernanceSession) GetExecutionStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetExecutionStageDuration(&_Governance.CallOpts)
}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetExecutionStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetExecutionStageDuration(&_Governance.CallOpts)
}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) constant returns(bool, bool, uint256)
func (_Governance *GovernanceCaller) GetHotfixRecord(opts *bind.CallOpts, hash [32]byte) (bool, bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(bool)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Governance.contract.Call(opts, out, "getHotfixRecord", hash)
	return *ret0, *ret1, *ret2, err
}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) constant returns(bool, bool, uint256)
func (_Governance *GovernanceSession) GetHotfixRecord(hash [32]byte) (bool, bool, *big.Int, error) {
	return _Governance.Contract.GetHotfixRecord(&_Governance.CallOpts, hash)
}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) constant returns(bool, bool, uint256)
func (_Governance *GovernanceCallerSession) GetHotfixRecord(hash [32]byte) (bool, bool, *big.Int, error) {
	return _Governance.Contract.GetHotfixRecord(&_Governance.CallOpts, hash)
}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) constant returns(uint256)
func (_Governance *GovernanceCaller) GetMostRecentReferendumProposal(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getMostRecentReferendumProposal", account)
	return *ret0, err
}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) constant returns(uint256)
func (_Governance *GovernanceSession) GetMostRecentReferendumProposal(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetMostRecentReferendumProposal(&_Governance.CallOpts, account)
}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetMostRecentReferendumProposal(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetMostRecentReferendumProposal(&_Governance.CallOpts, account)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Governance *GovernanceCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getParentSealBitmap", blockNumber)
	return *ret0, err
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Governance *GovernanceSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Governance.Contract.GetParentSealBitmap(&_Governance.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Governance *GovernanceCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Governance.Contract.GetParentSealBitmap(&_Governance.CallOpts, blockNumber)
}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetParticipationParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Governance.contract.Call(opts, out, "getParticipationParameters")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceSession) GetParticipationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetParticipationParameters(&_Governance.CallOpts)
}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetParticipationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetParticipationParameters(&_Governance.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) constant returns(address, uint256, uint256, uint256, string)
func (_Governance *GovernanceCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Governance.contract.Call(opts, out, "getProposal", proposalId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) constant returns(address, uint256, uint256, uint256, string)
func (_Governance *GovernanceSession) GetProposal(proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, error) {
	return _Governance.Contract.GetProposal(&_Governance.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) constant returns(address, uint256, uint256, uint256, string)
func (_Governance *GovernanceCallerSession) GetProposal(proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, error) {
	return _Governance.Contract.GetProposal(&_Governance.CallOpts, proposalId)
}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) constant returns(uint8)
func (_Governance *GovernanceCaller) GetProposalStage(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getProposalStage", proposalId)
	return *ret0, err
}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) constant returns(uint8)
func (_Governance *GovernanceSession) GetProposalStage(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.GetProposalStage(&_Governance.CallOpts, proposalId)
}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) constant returns(uint8)
func (_Governance *GovernanceCallerSession) GetProposalStage(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.GetProposalStage(&_Governance.CallOpts, proposalId)
}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) constant returns(uint256, address, bytes)
func (_Governance *GovernanceCaller) GetProposalTransaction(opts *bind.CallOpts, proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Governance.contract.Call(opts, out, "getProposalTransaction", proposalId, index)
	return *ret0, *ret1, *ret2, err
}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) constant returns(uint256, address, bytes)
func (_Governance *GovernanceSession) GetProposalTransaction(proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	return _Governance.Contract.GetProposalTransaction(&_Governance.CallOpts, proposalId, index)
}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) constant returns(uint256, address, bytes)
func (_Governance *GovernanceCallerSession) GetProposalTransaction(proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	return _Governance.Contract.GetProposalTransaction(&_Governance.CallOpts, proposalId, index)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() constant returns(uint256[], uint256[])
func (_Governance *GovernanceCaller) GetQueue(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Governance.contract.Call(opts, out, "getQueue")
	return *ret0, *ret1, err
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() constant returns(uint256[], uint256[])
func (_Governance *GovernanceSession) GetQueue() ([]*big.Int, []*big.Int, error) {
	return _Governance.Contract.GetQueue(&_Governance.CallOpts)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() constant returns(uint256[], uint256[])
func (_Governance *GovernanceCallerSession) GetQueue() ([]*big.Int, []*big.Int, error) {
	return _Governance.Contract.GetQueue(&_Governance.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() constant returns(uint256)
func (_Governance *GovernanceCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getQueueLength")
	return *ret0, err
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() constant returns(uint256)
func (_Governance *GovernanceSession) GetQueueLength() (*big.Int, error) {
	return _Governance.Contract.GetQueueLength(&_Governance.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetQueueLength() (*big.Int, error) {
	return _Governance.Contract.GetQueueLength(&_Governance.CallOpts)
}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() constant returns(uint256)
func (_Governance *GovernanceCaller) GetReferendumStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getReferendumStageDuration")
	return *ret0, err
}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() constant returns(uint256)
func (_Governance *GovernanceSession) GetReferendumStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetReferendumStageDuration(&_Governance.CallOpts)
}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetReferendumStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetReferendumStageDuration(&_Governance.CallOpts)
}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) constant returns(uint256, uint256)
func (_Governance *GovernanceCaller) GetUpvoteRecord(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Governance.contract.Call(opts, out, "getUpvoteRecord", account)
	return *ret0, *ret1, err
}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) constant returns(uint256, uint256)
func (_Governance *GovernanceSession) GetUpvoteRecord(account common.Address) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetUpvoteRecord(&_Governance.CallOpts, account)
}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) constant returns(uint256, uint256)
func (_Governance *GovernanceCallerSession) GetUpvoteRecord(account common.Address) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetUpvoteRecord(&_Governance.CallOpts, account)
}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) constant returns(uint256)
func (_Governance *GovernanceCaller) GetUpvotes(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getUpvotes", proposalId)
	return *ret0, err
}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) constant returns(uint256)
func (_Governance *GovernanceSession) GetUpvotes(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetUpvotes(&_Governance.CallOpts, proposalId)
}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetUpvotes(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetUpvotes(&_Governance.CallOpts, proposalId)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getVerifiedSealBitmapFromHeader", header)
	return *ret0, err
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.GetVerifiedSealBitmapFromHeader(&_Governance.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.GetVerifiedSealBitmapFromHeader(&_Governance.CallOpts, header)
}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) constant returns(uint256, uint256)
func (_Governance *GovernanceCaller) GetVoteRecord(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Governance.contract.Call(opts, out, "getVoteRecord", account, index)
	return *ret0, *ret1, err
}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) constant returns(uint256, uint256)
func (_Governance *GovernanceSession) GetVoteRecord(account common.Address, index *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteRecord(&_Governance.CallOpts, account, index)
}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) constant returns(uint256, uint256)
func (_Governance *GovernanceCallerSession) GetVoteRecord(account common.Address, index *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteRecord(&_Governance.CallOpts, account, index)
}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) constant returns(uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetVoteTotals(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Governance.contract.Call(opts, out, "getVoteTotals", proposalId)
	return *ret0, *ret1, *ret2, err
}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) constant returns(uint256, uint256, uint256)
func (_Governance *GovernanceSession) GetVoteTotals(proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteTotals(&_Governance.CallOpts, proposalId)
}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) constant returns(uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetVoteTotals(proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteTotals(&_Governance.CallOpts, proposalId)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "hashHeader", header)
	return *ret0, err
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceSession) HashHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.HashHeader(&_Governance.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Governance *GovernanceCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.HashHeader(&_Governance.CallOpts, header)
}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) constant returns(uint256)
func (_Governance *GovernanceCaller) HotfixWhitelistValidatorTally(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "hotfixWhitelistValidatorTally", hash)
	return *ret0, err
}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) constant returns(uint256)
func (_Governance *GovernanceSession) HotfixWhitelistValidatorTally(hash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HotfixWhitelistValidatorTally(&_Governance.CallOpts, hash)
}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) constant returns(uint256)
func (_Governance *GovernanceCallerSession) HotfixWhitelistValidatorTally(hash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HotfixWhitelistValidatorTally(&_Governance.CallOpts, hash)
}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) constant returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceCaller) Hotfixes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	ret := new(struct {
		Executed      bool
		Approved      bool
		PreparedEpoch *big.Int
	})
	out := ret
	err := _Governance.contract.Call(opts, out, "hotfixes", arg0)
	return *ret, err
}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) constant returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceSession) Hotfixes(arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	return _Governance.Contract.Hotfixes(&_Governance.CallOpts, arg0)
}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) constant returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceCallerSession) Hotfixes(arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	return _Governance.Contract.Hotfixes(&_Governance.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Governance *GovernanceCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Governance *GovernanceSession) Initialized() (bool, error) {
	return _Governance.Contract.Initialized(&_Governance.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Governance *GovernanceCallerSession) Initialized() (bool, error) {
	return _Governance.Contract.Initialized(&_Governance.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) IsApproved(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isApproved", proposalId)
	return *ret0, err
}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) IsApproved(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsApproved(&_Governance.CallOpts, proposalId)
}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsApproved(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsApproved(&_Governance.CallOpts, proposalId)
}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) constant returns(bool)
func (_Governance *GovernanceCaller) IsDequeuedProposal(opts *bind.CallOpts, proposalId *big.Int, index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isDequeuedProposal", proposalId, index)
	return *ret0, err
}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) constant returns(bool)
func (_Governance *GovernanceSession) IsDequeuedProposal(proposalId *big.Int, index *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposal(&_Governance.CallOpts, proposalId, index)
}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsDequeuedProposal(proposalId *big.Int, index *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposal(&_Governance.CallOpts, proposalId, index)
}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) IsDequeuedProposalExpired(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isDequeuedProposalExpired", proposalId)
	return *ret0, err
}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) IsDequeuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsDequeuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) constant returns(bool)
func (_Governance *GovernanceCaller) IsHotfixPassing(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isHotfixPassing", hash)
	return *ret0, err
}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) constant returns(bool)
func (_Governance *GovernanceSession) IsHotfixPassing(hash [32]byte) (bool, error) {
	return _Governance.Contract.IsHotfixPassing(&_Governance.CallOpts, hash)
}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsHotfixPassing(hash [32]byte) (bool, error) {
	return _Governance.Contract.IsHotfixPassing(&_Governance.CallOpts, hash)
}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) constant returns(bool)
func (_Governance *GovernanceCaller) IsHotfixWhitelistedBy(opts *bind.CallOpts, hash [32]byte, whitelister common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isHotfixWhitelistedBy", hash, whitelister)
	return *ret0, err
}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) constant returns(bool)
func (_Governance *GovernanceSession) IsHotfixWhitelistedBy(hash [32]byte, whitelister common.Address) (bool, error) {
	return _Governance.Contract.IsHotfixWhitelistedBy(&_Governance.CallOpts, hash, whitelister)
}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsHotfixWhitelistedBy(hash [32]byte, whitelister common.Address) (bool, error) {
	return _Governance.Contract.IsHotfixWhitelistedBy(&_Governance.CallOpts, hash, whitelister)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Governance *GovernanceCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Governance *GovernanceSession) IsOwner() (bool, error) {
	return _Governance.Contract.IsOwner(&_Governance.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Governance *GovernanceCallerSession) IsOwner() (bool, error) {
	return _Governance.Contract.IsOwner(&_Governance.CallOpts)
}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) IsProposalPassing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isProposalPassing", proposalId)
	return *ret0, err
}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) IsProposalPassing(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsProposalPassing(&_Governance.CallOpts, proposalId)
}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsProposalPassing(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsProposalPassing(&_Governance.CallOpts, proposalId)
}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) IsQueued(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isQueued", proposalId)
	return *ret0, err
}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) IsQueued(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueued(&_Governance.CallOpts, proposalId)
}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsQueued(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueued(&_Governance.CallOpts, proposalId)
}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) IsQueuedProposalExpired(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isQueuedProposalExpired", proposalId)
	return *ret0, err
}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) IsQueuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsQueuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) constant returns(bool)
func (_Governance *GovernanceCaller) IsVoting(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "isVoting", account)
	return *ret0, err
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) constant returns(bool)
func (_Governance *GovernanceSession) IsVoting(account common.Address) (bool, error) {
	return _Governance.Contract.IsVoting(&_Governance.CallOpts, account)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) constant returns(bool)
func (_Governance *GovernanceCallerSession) IsVoting(account common.Address) (bool, error) {
	return _Governance.Contract.IsVoting(&_Governance.CallOpts, account)
}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() constant returns(uint256)
func (_Governance *GovernanceCaller) LastDequeue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lastDequeue")
	return *ret0, err
}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() constant returns(uint256)
func (_Governance *GovernanceSession) LastDequeue() (*big.Int, error) {
	return _Governance.Contract.LastDequeue(&_Governance.CallOpts)
}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LastDequeue() (*big.Int, error) {
	return _Governance.Contract.LastDequeue(&_Governance.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Governance *GovernanceCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minDeposit")
	return *ret0, err
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Governance *GovernanceSession) MinDeposit() (*big.Int, error) {
	return _Governance.Contract.MinDeposit(&_Governance.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinDeposit() (*big.Int, error) {
	return _Governance.Contract.MinDeposit(&_Governance.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minQuorumSize", blockNumber)
	return *ret0, err
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinQuorumSize(&_Governance.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinQuorumSize(&_Governance.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minQuorumSizeInCurrentSet")
	return *ret0, err
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.MinQuorumSizeInCurrentSet(&_Governance.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.MinQuorumSizeInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "numberValidatorsInCurrentSet")
	return *ret0, err
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "numberValidatorsInSet", blockNumber)
	return *ret0, err
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInSet(&_Governance.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Governance *GovernanceCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInSet(&_Governance.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() constant returns(uint256)
func (_Governance *GovernanceCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "proposalCount")
	return *ret0, err
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() constant returns(uint256)
func (_Governance *GovernanceSession) ProposalCount() (*big.Int, error) {
	return _Governance.Contract.ProposalCount(&_Governance.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalCount() (*big.Int, error) {
	return _Governance.Contract.ProposalCount(&_Governance.CallOpts)
}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCaller) ProposalExists(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "proposalExists", proposalId)
	return *ret0, err
}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceSession) ProposalExists(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.ProposalExists(&_Governance.CallOpts, proposalId)
}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) constant returns(bool)
func (_Governance *GovernanceCallerSession) ProposalExists(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.ProposalExists(&_Governance.CallOpts, proposalId)
}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() constant returns(uint256)
func (_Governance *GovernanceCaller) QueueExpiry(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "queueExpiry")
	return *ret0, err
}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() constant returns(uint256)
func (_Governance *GovernanceSession) QueueExpiry() (*big.Int, error) {
	return _Governance.Contract.QueueExpiry(&_Governance.CallOpts)
}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() constant returns(uint256)
func (_Governance *GovernanceCallerSession) QueueExpiry() (*big.Int, error) {
	return _Governance.Contract.QueueExpiry(&_Governance.CallOpts)
}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) constant returns(uint256)
func (_Governance *GovernanceCaller) RefundedDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "refundedDeposits", arg0)
	return *ret0, err
}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) constant returns(uint256)
func (_Governance *GovernanceSession) RefundedDeposits(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.RefundedDeposits(&_Governance.CallOpts, arg0)
}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) RefundedDeposits(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.RefundedDeposits(&_Governance.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Governance *GovernanceCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Governance *GovernanceSession) Registry() (common.Address, error) {
	return _Governance.Contract.Registry(&_Governance.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Governance *GovernanceCallerSession) Registry() (common.Address, error) {
	return _Governance.Contract.Registry(&_Governance.CallOpts)
}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() constant returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceCaller) StageDurations(opts *bind.CallOpts) (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	ret := new(struct {
		Approval   *big.Int
		Referendum *big.Int
		Execution  *big.Int
	})
	out := ret
	err := _Governance.contract.Call(opts, out, "stageDurations")
	return *ret, err
}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() constant returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceSession) StageDurations() (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	return _Governance.Contract.StageDurations(&_Governance.CallOpts)
}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() constant returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceCallerSession) StageDurations() (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	return _Governance.Contract.StageDurations(&_Governance.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Governance *GovernanceCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "validatorSignerAddressFromCurrentSet", index)
	return *ret0, err
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Governance *GovernanceSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromCurrentSet(&_Governance.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Governance *GovernanceCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromCurrentSet(&_Governance.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Governance *GovernanceCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "validatorSignerAddressFromSet", index, blockNumber)
	return *ret0, err
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Governance *GovernanceSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromSet(&_Governance.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Governance *GovernanceCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromSet(&_Governance.CallOpts, index, blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactor) Approve(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "approve", proposalId, index)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceSession) Approve(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Approve(&_Governance.TransactOpts, proposalId, index)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactorSession) Approve(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Approve(&_Governance.TransactOpts, proposalId, index)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) ApproveHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "approveHotfix", hash)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) ApproveHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ApproveHotfix(&_Governance.TransactOpts, hash)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) ApproveHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ApproveHotfix(&_Governance.TransactOpts, hash)
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceTransactor) DequeueProposalsIfReady(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "dequeueProposalsIfReady")
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceSession) DequeueProposalsIfReady() (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalsIfReady(&_Governance.TransactOpts)
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceTransactorSession) DequeueProposalsIfReady() (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalsIfReady(&_Governance.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "execute", proposalId, index)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceSession) Execute(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId, index)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactorSession) Execute(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId, index)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceTransactor) ExecuteHotfix(opts *bind.TransactOpts, values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "executeHotfix", values, destinations, data, dataLengths, salt)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceSession) ExecuteHotfix(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ExecuteHotfix(&_Governance.TransactOpts, values, destinations, data, dataLengths, salt)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceTransactorSession) ExecuteHotfix(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ExecuteHotfix(&_Governance.TransactOpts, values, destinations, data, dataLengths, salt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc1939b20.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 approvalStageDuration, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, approvalStageDuration *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, approvalStageDuration, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc1939b20.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 approvalStageDuration, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceSession) Initialize(registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, approvalStageDuration *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, approvalStageDuration, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc1939b20.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 approvalStageDuration, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactorSession) Initialize(registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, approvalStageDuration *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, approvalStageDuration, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) PrepareHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "prepareHotfix", hash)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) PrepareHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.PrepareHotfix(&_Governance.TransactOpts, hash)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) PrepareHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.PrepareHotfix(&_Governance.TransactOpts, hash)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) returns(uint256)
func (_Governance *GovernanceTransactor) Propose(opts *bind.TransactOpts, values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose", values, destinations, data, dataLengths, descriptionUrl)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) returns(uint256)
func (_Governance *GovernanceSession) Propose(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, values, destinations, data, dataLengths, descriptionUrl)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) returns(uint256)
func (_Governance *GovernanceTransactorSession) Propose(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, values, destinations, data, dataLengths, descriptionUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactor) RevokeUpvote(opts *bind.TransactOpts, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "revokeUpvote", lesser, greater)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceSession) RevokeUpvote(lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RevokeUpvote(&_Governance.TransactOpts, lesser, greater)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactorSession) RevokeUpvote(lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RevokeUpvote(&_Governance.TransactOpts, lesser, greater)
}

// SetApprovalStageDuration is a paid mutator transaction binding the contract method 0x9a6c3d83.
//
// Solidity: function setApprovalStageDuration(uint256 approvalStageDuration) returns()
func (_Governance *GovernanceTransactor) SetApprovalStageDuration(opts *bind.TransactOpts, approvalStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setApprovalStageDuration", approvalStageDuration)
}

// SetApprovalStageDuration is a paid mutator transaction binding the contract method 0x9a6c3d83.
//
// Solidity: function setApprovalStageDuration(uint256 approvalStageDuration) returns()
func (_Governance *GovernanceSession) SetApprovalStageDuration(approvalStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetApprovalStageDuration(&_Governance.TransactOpts, approvalStageDuration)
}

// SetApprovalStageDuration is a paid mutator transaction binding the contract method 0x9a6c3d83.
//
// Solidity: function setApprovalStageDuration(uint256 approvalStageDuration) returns()
func (_Governance *GovernanceTransactorSession) SetApprovalStageDuration(approvalStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetApprovalStageDuration(&_Governance.TransactOpts, approvalStageDuration)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceTransactor) SetApprover(opts *bind.TransactOpts, _approver common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setApprover", _approver)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceSession) SetApprover(_approver common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetApprover(&_Governance.TransactOpts, _approver)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceTransactorSession) SetApprover(_approver common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetApprover(&_Governance.TransactOpts, _approver)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactor) SetBaselineQuorumFactor(opts *bind.TransactOpts, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setBaselineQuorumFactor", baselineQuorumFactor)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceSession) SetBaselineQuorumFactor(baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineQuorumFactor(&_Governance.TransactOpts, baselineQuorumFactor)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactorSession) SetBaselineQuorumFactor(baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineQuorumFactor(&_Governance.TransactOpts, baselineQuorumFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceTransactor) SetBaselineUpdateFactor(opts *bind.TransactOpts, baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setBaselineUpdateFactor", baselineUpdateFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceSession) SetBaselineUpdateFactor(baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineUpdateFactor(&_Governance.TransactOpts, baselineUpdateFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceTransactorSession) SetBaselineUpdateFactor(baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineUpdateFactor(&_Governance.TransactOpts, baselineUpdateFactor)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceTransactor) SetConcurrentProposals(opts *bind.TransactOpts, _concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setConcurrentProposals", _concurrentProposals)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceSession) SetConcurrentProposals(_concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConcurrentProposals(&_Governance.TransactOpts, _concurrentProposals)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceTransactorSession) SetConcurrentProposals(_concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConcurrentProposals(&_Governance.TransactOpts, _concurrentProposals)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceTransactor) SetConstitution(opts *bind.TransactOpts, destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setConstitution", destination, functionId, threshold)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceSession) SetConstitution(destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConstitution(&_Governance.TransactOpts, destination, functionId, threshold)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceTransactorSession) SetConstitution(destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConstitution(&_Governance.TransactOpts, destination, functionId, threshold)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceTransactor) SetDequeueFrequency(opts *bind.TransactOpts, _dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setDequeueFrequency", _dequeueFrequency)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceSession) SetDequeueFrequency(_dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetDequeueFrequency(&_Governance.TransactOpts, _dequeueFrequency)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceTransactorSession) SetDequeueFrequency(_dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetDequeueFrequency(&_Governance.TransactOpts, _dequeueFrequency)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceTransactor) SetExecutionStageDuration(opts *bind.TransactOpts, executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setExecutionStageDuration", executionStageDuration)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceSession) SetExecutionStageDuration(executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetExecutionStageDuration(&_Governance.TransactOpts, executionStageDuration)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceTransactorSession) SetExecutionStageDuration(executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetExecutionStageDuration(&_Governance.TransactOpts, executionStageDuration)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceTransactor) SetMinDeposit(opts *bind.TransactOpts, _minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setMinDeposit", _minDeposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceSession) SetMinDeposit(_minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinDeposit(&_Governance.TransactOpts, _minDeposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceTransactorSession) SetMinDeposit(_minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinDeposit(&_Governance.TransactOpts, _minDeposit)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceTransactor) SetParticipationBaseline(opts *bind.TransactOpts, participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setParticipationBaseline", participationBaseline)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceSession) SetParticipationBaseline(participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationBaseline(&_Governance.TransactOpts, participationBaseline)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceTransactorSession) SetParticipationBaseline(participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationBaseline(&_Governance.TransactOpts, participationBaseline)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceTransactor) SetParticipationFloor(opts *bind.TransactOpts, participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setParticipationFloor", participationFloor)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceSession) SetParticipationFloor(participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationFloor(&_Governance.TransactOpts, participationFloor)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceTransactorSession) SetParticipationFloor(participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationFloor(&_Governance.TransactOpts, participationFloor)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceTransactor) SetQueueExpiry(opts *bind.TransactOpts, _queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setQueueExpiry", _queueExpiry)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceSession) SetQueueExpiry(_queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueExpiry(&_Governance.TransactOpts, _queueExpiry)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceTransactorSession) SetQueueExpiry(_queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueExpiry(&_Governance.TransactOpts, _queueExpiry)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceTransactor) SetReferendumStageDuration(opts *bind.TransactOpts, referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setReferendumStageDuration", referendumStageDuration)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceSession) SetReferendumStageDuration(referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetReferendumStageDuration(&_Governance.TransactOpts, referendumStageDuration)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceTransactorSession) SetReferendumStageDuration(referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetReferendumStageDuration(&_Governance.TransactOpts, referendumStageDuration)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRegistry(&_Governance.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRegistry(&_Governance.TransactOpts, registryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactor) Upvote(opts *bind.TransactOpts, proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "upvote", proposalId, lesser, greater)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceSession) Upvote(proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Upvote(&_Governance.TransactOpts, proposalId, lesser, greater)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactorSession) Upvote(proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Upvote(&_Governance.TransactOpts, proposalId, lesser, greater)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceTransactor) Vote(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "vote", proposalId, index, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceSession) Vote(proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, proposalId, index, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceTransactorSession) Vote(proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, proposalId, index, value)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) WhitelistHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "whitelistHotfix", hash)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) WhitelistHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.WhitelistHotfix(&_Governance.TransactOpts, hash)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) WhitelistHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.WhitelistHotfix(&_Governance.TransactOpts, hash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// GovernanceApprovalStageDurationSetIterator is returned from FilterApprovalStageDurationSet and is used to iterate over the raw logs and unpacked data for ApprovalStageDurationSet events raised by the Governance contract.
type GovernanceApprovalStageDurationSetIterator struct {
	Event *GovernanceApprovalStageDurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceApprovalStageDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApprovalStageDurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceApprovalStageDurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceApprovalStageDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApprovalStageDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApprovalStageDurationSet represents a ApprovalStageDurationSet event raised by the Governance contract.
type GovernanceApprovalStageDurationSet struct {
	ApprovalStageDuration *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterApprovalStageDurationSet is a free log retrieval operation binding the contract event 0xbc44c003a66b841b48c220869efb738b265af305649ac8bafe8efe0c8130e313.
//
// Solidity: event ApprovalStageDurationSet(uint256 approvalStageDuration)
func (_Governance *GovernanceFilterer) FilterApprovalStageDurationSet(opts *bind.FilterOpts) (*GovernanceApprovalStageDurationSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ApprovalStageDurationSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceApprovalStageDurationSetIterator{contract: _Governance.contract, event: "ApprovalStageDurationSet", logs: logs, sub: sub}, nil
}

// WatchApprovalStageDurationSet is a free log subscription operation binding the contract event 0xbc44c003a66b841b48c220869efb738b265af305649ac8bafe8efe0c8130e313.
//
// Solidity: event ApprovalStageDurationSet(uint256 approvalStageDuration)
func (_Governance *GovernanceFilterer) WatchApprovalStageDurationSet(opts *bind.WatchOpts, sink chan<- *GovernanceApprovalStageDurationSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ApprovalStageDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApprovalStageDurationSet)
				if err := _Governance.contract.UnpackLog(event, "ApprovalStageDurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalStageDurationSet is a log parse operation binding the contract event 0xbc44c003a66b841b48c220869efb738b265af305649ac8bafe8efe0c8130e313.
//
// Solidity: event ApprovalStageDurationSet(uint256 approvalStageDuration)
func (_Governance *GovernanceFilterer) ParseApprovalStageDurationSet(log types.Log) (*GovernanceApprovalStageDurationSet, error) {
	event := new(GovernanceApprovalStageDurationSet)
	if err := _Governance.contract.UnpackLog(event, "ApprovalStageDurationSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceApproverSetIterator is returned from FilterApproverSet and is used to iterate over the raw logs and unpacked data for ApproverSet events raised by the Governance contract.
type GovernanceApproverSetIterator struct {
	Event *GovernanceApproverSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceApproverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceApproverSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceApproverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverSet represents a ApproverSet event raised by the Governance contract.
type GovernanceApproverSet struct {
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproverSet is a free log retrieval operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) FilterApproverSet(opts *bind.FilterOpts, approver []common.Address) (*GovernanceApproverSetIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ApproverSet", approverRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverSetIterator{contract: _Governance.contract, event: "ApproverSet", logs: logs, sub: sub}, nil
}

// WatchApproverSet is a free log subscription operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) WatchApproverSet(opts *bind.WatchOpts, sink chan<- *GovernanceApproverSet, approver []common.Address) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ApproverSet", approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverSet)
				if err := _Governance.contract.UnpackLog(event, "ApproverSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproverSet is a log parse operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) ParseApproverSet(log types.Log) (*GovernanceApproverSet, error) {
	event := new(GovernanceApproverSet)
	if err := _Governance.contract.UnpackLog(event, "ApproverSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceConcurrentProposalsSetIterator is returned from FilterConcurrentProposalsSet and is used to iterate over the raw logs and unpacked data for ConcurrentProposalsSet events raised by the Governance contract.
type GovernanceConcurrentProposalsSetIterator struct {
	Event *GovernanceConcurrentProposalsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceConcurrentProposalsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConcurrentProposalsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceConcurrentProposalsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceConcurrentProposalsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConcurrentProposalsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConcurrentProposalsSet represents a ConcurrentProposalsSet event raised by the Governance contract.
type GovernanceConcurrentProposalsSet struct {
	ConcurrentProposals *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConcurrentProposalsSet is a free log retrieval operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) FilterConcurrentProposalsSet(opts *bind.FilterOpts) (*GovernanceConcurrentProposalsSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConcurrentProposalsSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceConcurrentProposalsSetIterator{contract: _Governance.contract, event: "ConcurrentProposalsSet", logs: logs, sub: sub}, nil
}

// WatchConcurrentProposalsSet is a free log subscription operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) WatchConcurrentProposalsSet(opts *bind.WatchOpts, sink chan<- *GovernanceConcurrentProposalsSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConcurrentProposalsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConcurrentProposalsSet)
				if err := _Governance.contract.UnpackLog(event, "ConcurrentProposalsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConcurrentProposalsSet is a log parse operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) ParseConcurrentProposalsSet(log types.Log) (*GovernanceConcurrentProposalsSet, error) {
	event := new(GovernanceConcurrentProposalsSet)
	if err := _Governance.contract.UnpackLog(event, "ConcurrentProposalsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceConstitutionSetIterator is returned from FilterConstitutionSet and is used to iterate over the raw logs and unpacked data for ConstitutionSet events raised by the Governance contract.
type GovernanceConstitutionSetIterator struct {
	Event *GovernanceConstitutionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceConstitutionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConstitutionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceConstitutionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceConstitutionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConstitutionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConstitutionSet represents a ConstitutionSet event raised by the Governance contract.
type GovernanceConstitutionSet struct {
	Destination common.Address
	FunctionId  [4]byte
	Threshold   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConstitutionSet is a free log retrieval operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) FilterConstitutionSet(opts *bind.FilterOpts, destination []common.Address, functionId [][4]byte) (*GovernanceConstitutionSetIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConstitutionSet", destinationRule, functionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceConstitutionSetIterator{contract: _Governance.contract, event: "ConstitutionSet", logs: logs, sub: sub}, nil
}

// WatchConstitutionSet is a free log subscription operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) WatchConstitutionSet(opts *bind.WatchOpts, sink chan<- *GovernanceConstitutionSet, destination []common.Address, functionId [][4]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConstitutionSet", destinationRule, functionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConstitutionSet)
				if err := _Governance.contract.UnpackLog(event, "ConstitutionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConstitutionSet is a log parse operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) ParseConstitutionSet(log types.Log) (*GovernanceConstitutionSet, error) {
	event := new(GovernanceConstitutionSet)
	if err := _Governance.contract.UnpackLog(event, "ConstitutionSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceDequeueFrequencySetIterator is returned from FilterDequeueFrequencySet and is used to iterate over the raw logs and unpacked data for DequeueFrequencySet events raised by the Governance contract.
type GovernanceDequeueFrequencySetIterator struct {
	Event *GovernanceDequeueFrequencySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceDequeueFrequencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDequeueFrequencySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceDequeueFrequencySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceDequeueFrequencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceDequeueFrequencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceDequeueFrequencySet represents a DequeueFrequencySet event raised by the Governance contract.
type GovernanceDequeueFrequencySet struct {
	DequeueFrequency *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDequeueFrequencySet is a free log retrieval operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) FilterDequeueFrequencySet(opts *bind.FilterOpts) (*GovernanceDequeueFrequencySetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "DequeueFrequencySet")
	if err != nil {
		return nil, err
	}
	return &GovernanceDequeueFrequencySetIterator{contract: _Governance.contract, event: "DequeueFrequencySet", logs: logs, sub: sub}, nil
}

// WatchDequeueFrequencySet is a free log subscription operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) WatchDequeueFrequencySet(opts *bind.WatchOpts, sink chan<- *GovernanceDequeueFrequencySet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "DequeueFrequencySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceDequeueFrequencySet)
				if err := _Governance.contract.UnpackLog(event, "DequeueFrequencySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDequeueFrequencySet is a log parse operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) ParseDequeueFrequencySet(log types.Log) (*GovernanceDequeueFrequencySet, error) {
	event := new(GovernanceDequeueFrequencySet)
	if err := _Governance.contract.UnpackLog(event, "DequeueFrequencySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceExecutionStageDurationSetIterator is returned from FilterExecutionStageDurationSet and is used to iterate over the raw logs and unpacked data for ExecutionStageDurationSet events raised by the Governance contract.
type GovernanceExecutionStageDurationSetIterator struct {
	Event *GovernanceExecutionStageDurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceExecutionStageDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceExecutionStageDurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceExecutionStageDurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceExecutionStageDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceExecutionStageDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceExecutionStageDurationSet represents a ExecutionStageDurationSet event raised by the Governance contract.
type GovernanceExecutionStageDurationSet struct {
	ExecutionStageDuration *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterExecutionStageDurationSet is a free log retrieval operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) FilterExecutionStageDurationSet(opts *bind.FilterOpts) (*GovernanceExecutionStageDurationSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ExecutionStageDurationSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceExecutionStageDurationSetIterator{contract: _Governance.contract, event: "ExecutionStageDurationSet", logs: logs, sub: sub}, nil
}

// WatchExecutionStageDurationSet is a free log subscription operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) WatchExecutionStageDurationSet(opts *bind.WatchOpts, sink chan<- *GovernanceExecutionStageDurationSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ExecutionStageDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceExecutionStageDurationSet)
				if err := _Governance.contract.UnpackLog(event, "ExecutionStageDurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionStageDurationSet is a log parse operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) ParseExecutionStageDurationSet(log types.Log) (*GovernanceExecutionStageDurationSet, error) {
	event := new(GovernanceExecutionStageDurationSet)
	if err := _Governance.contract.UnpackLog(event, "ExecutionStageDurationSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceHotfixApprovedIterator is returned from FilterHotfixApproved and is used to iterate over the raw logs and unpacked data for HotfixApproved events raised by the Governance contract.
type GovernanceHotfixApprovedIterator struct {
	Event *GovernanceHotfixApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceHotfixApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceHotfixApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceHotfixApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixApproved represents a HotfixApproved event raised by the Governance contract.
type GovernanceHotfixApproved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHotfixApproved is a free log retrieval operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) FilterHotfixApproved(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixApprovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixApproved", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixApprovedIterator{contract: _Governance.contract, event: "HotfixApproved", logs: logs, sub: sub}, nil
}

// WatchHotfixApproved is a free log subscription operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) WatchHotfixApproved(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixApproved, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixApproved", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixApproved)
				if err := _Governance.contract.UnpackLog(event, "HotfixApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHotfixApproved is a log parse operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) ParseHotfixApproved(log types.Log) (*GovernanceHotfixApproved, error) {
	event := new(GovernanceHotfixApproved)
	if err := _Governance.contract.UnpackLog(event, "HotfixApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceHotfixExecutedIterator is returned from FilterHotfixExecuted and is used to iterate over the raw logs and unpacked data for HotfixExecuted events raised by the Governance contract.
type GovernanceHotfixExecutedIterator struct {
	Event *GovernanceHotfixExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceHotfixExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceHotfixExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceHotfixExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixExecuted represents a HotfixExecuted event raised by the Governance contract.
type GovernanceHotfixExecuted struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHotfixExecuted is a free log retrieval operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) FilterHotfixExecuted(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixExecutedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixExecuted", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixExecutedIterator{contract: _Governance.contract, event: "HotfixExecuted", logs: logs, sub: sub}, nil
}

// WatchHotfixExecuted is a free log subscription operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) WatchHotfixExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixExecuted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixExecuted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixExecuted)
				if err := _Governance.contract.UnpackLog(event, "HotfixExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHotfixExecuted is a log parse operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) ParseHotfixExecuted(log types.Log) (*GovernanceHotfixExecuted, error) {
	event := new(GovernanceHotfixExecuted)
	if err := _Governance.contract.UnpackLog(event, "HotfixExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceHotfixPreparedIterator is returned from FilterHotfixPrepared and is used to iterate over the raw logs and unpacked data for HotfixPrepared events raised by the Governance contract.
type GovernanceHotfixPreparedIterator struct {
	Event *GovernanceHotfixPrepared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceHotfixPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixPrepared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceHotfixPrepared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceHotfixPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixPrepared represents a HotfixPrepared event raised by the Governance contract.
type GovernanceHotfixPrepared struct {
	Hash  [32]byte
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHotfixPrepared is a free log retrieval operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) FilterHotfixPrepared(opts *bind.FilterOpts, hash [][32]byte, epoch []*big.Int) (*GovernanceHotfixPreparedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixPrepared", hashRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixPreparedIterator{contract: _Governance.contract, event: "HotfixPrepared", logs: logs, sub: sub}, nil
}

// WatchHotfixPrepared is a free log subscription operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) WatchHotfixPrepared(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixPrepared, hash [][32]byte, epoch []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixPrepared", hashRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixPrepared)
				if err := _Governance.contract.UnpackLog(event, "HotfixPrepared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHotfixPrepared is a log parse operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) ParseHotfixPrepared(log types.Log) (*GovernanceHotfixPrepared, error) {
	event := new(GovernanceHotfixPrepared)
	if err := _Governance.contract.UnpackLog(event, "HotfixPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceHotfixWhitelistedIterator is returned from FilterHotfixWhitelisted and is used to iterate over the raw logs and unpacked data for HotfixWhitelisted events raised by the Governance contract.
type GovernanceHotfixWhitelistedIterator struct {
	Event *GovernanceHotfixWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceHotfixWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceHotfixWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceHotfixWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixWhitelisted represents a HotfixWhitelisted event raised by the Governance contract.
type GovernanceHotfixWhitelisted struct {
	Hash        [32]byte
	Whitelister common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHotfixWhitelisted is a free log retrieval operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) FilterHotfixWhitelisted(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixWhitelistedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixWhitelisted", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixWhitelistedIterator{contract: _Governance.contract, event: "HotfixWhitelisted", logs: logs, sub: sub}, nil
}

// WatchHotfixWhitelisted is a free log subscription operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) WatchHotfixWhitelisted(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixWhitelisted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixWhitelisted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixWhitelisted)
				if err := _Governance.contract.UnpackLog(event, "HotfixWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHotfixWhitelisted is a log parse operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) ParseHotfixWhitelisted(log types.Log) (*GovernanceHotfixWhitelisted, error) {
	event := new(GovernanceHotfixWhitelisted)
	if err := _Governance.contract.UnpackLog(event, "HotfixWhitelisted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceMinDepositSetIterator is returned from FilterMinDepositSet and is used to iterate over the raw logs and unpacked data for MinDepositSet events raised by the Governance contract.
type GovernanceMinDepositSetIterator struct {
	Event *GovernanceMinDepositSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceMinDepositSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceMinDepositSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceMinDepositSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceMinDepositSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceMinDepositSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceMinDepositSet represents a MinDepositSet event raised by the Governance contract.
type GovernanceMinDepositSet struct {
	MinDeposit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinDepositSet is a free log retrieval operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) FilterMinDepositSet(opts *bind.FilterOpts) (*GovernanceMinDepositSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "MinDepositSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceMinDepositSetIterator{contract: _Governance.contract, event: "MinDepositSet", logs: logs, sub: sub}, nil
}

// WatchMinDepositSet is a free log subscription operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) WatchMinDepositSet(opts *bind.WatchOpts, sink chan<- *GovernanceMinDepositSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "MinDepositSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceMinDepositSet)
				if err := _Governance.contract.UnpackLog(event, "MinDepositSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDepositSet is a log parse operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) ParseMinDepositSet(log types.Log) (*GovernanceMinDepositSet, error) {
	event := new(GovernanceMinDepositSet)
	if err := _Governance.contract.UnpackLog(event, "MinDepositSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Governance contract.
type GovernanceOwnershipTransferredIterator struct {
	Event *GovernanceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceOwnershipTransferred represents a OwnershipTransferred event raised by the Governance contract.
type GovernanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernanceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceOwnershipTransferredIterator{contract: _Governance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceOwnershipTransferred)
				if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) ParseOwnershipTransferred(log types.Log) (*GovernanceOwnershipTransferred, error) {
	event := new(GovernanceOwnershipTransferred)
	if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceParticipationBaselineQuorumFactorSetIterator is returned from FilterParticipationBaselineQuorumFactorSet and is used to iterate over the raw logs and unpacked data for ParticipationBaselineQuorumFactorSet events raised by the Governance contract.
type GovernanceParticipationBaselineQuorumFactorSetIterator struct {
	Event *GovernanceParticipationBaselineQuorumFactorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineQuorumFactorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceParticipationBaselineQuorumFactorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineQuorumFactorSet represents a ParticipationBaselineQuorumFactorSet event raised by the Governance contract.
type GovernanceParticipationBaselineQuorumFactorSet struct {
	BaselineQuorumFactor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineQuorumFactorSet is a free log retrieval operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineQuorumFactorSet(opts *bind.FilterOpts) (*GovernanceParticipationBaselineQuorumFactorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineQuorumFactorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineQuorumFactorSetIterator{contract: _Governance.contract, event: "ParticipationBaselineQuorumFactorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineQuorumFactorSet is a free log subscription operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineQuorumFactorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineQuorumFactorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineQuorumFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineQuorumFactorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineQuorumFactorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipationBaselineQuorumFactorSet is a log parse operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineQuorumFactorSet(log types.Log) (*GovernanceParticipationBaselineQuorumFactorSet, error) {
	event := new(GovernanceParticipationBaselineQuorumFactorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineQuorumFactorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceParticipationBaselineUpdateFactorSetIterator is returned from FilterParticipationBaselineUpdateFactorSet and is used to iterate over the raw logs and unpacked data for ParticipationBaselineUpdateFactorSet events raised by the Governance contract.
type GovernanceParticipationBaselineUpdateFactorSetIterator struct {
	Event *GovernanceParticipationBaselineUpdateFactorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineUpdateFactorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceParticipationBaselineUpdateFactorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineUpdateFactorSet represents a ParticipationBaselineUpdateFactorSet event raised by the Governance contract.
type GovernanceParticipationBaselineUpdateFactorSet struct {
	BaselineUpdateFactor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineUpdateFactorSet is a free log retrieval operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineUpdateFactorSet(opts *bind.FilterOpts) (*GovernanceParticipationBaselineUpdateFactorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineUpdateFactorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineUpdateFactorSetIterator{contract: _Governance.contract, event: "ParticipationBaselineUpdateFactorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineUpdateFactorSet is a free log subscription operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineUpdateFactorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineUpdateFactorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineUpdateFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineUpdateFactorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdateFactorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipationBaselineUpdateFactorSet is a log parse operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineUpdateFactorSet(log types.Log) (*GovernanceParticipationBaselineUpdateFactorSet, error) {
	event := new(GovernanceParticipationBaselineUpdateFactorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdateFactorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceParticipationBaselineUpdatedIterator is returned from FilterParticipationBaselineUpdated and is used to iterate over the raw logs and unpacked data for ParticipationBaselineUpdated events raised by the Governance contract.
type GovernanceParticipationBaselineUpdatedIterator struct {
	Event *GovernanceParticipationBaselineUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceParticipationBaselineUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceParticipationBaselineUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceParticipationBaselineUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineUpdated represents a ParticipationBaselineUpdated event raised by the Governance contract.
type GovernanceParticipationBaselineUpdated struct {
	ParticipationBaseline *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineUpdated is a free log retrieval operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineUpdated(opts *bind.FilterOpts) (*GovernanceParticipationBaselineUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineUpdatedIterator{contract: _Governance.contract, event: "ParticipationBaselineUpdated", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineUpdated is a free log subscription operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineUpdated)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipationBaselineUpdated is a log parse operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineUpdated(log types.Log) (*GovernanceParticipationBaselineUpdated, error) {
	event := new(GovernanceParticipationBaselineUpdated)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceParticipationFloorSetIterator is returned from FilterParticipationFloorSet and is used to iterate over the raw logs and unpacked data for ParticipationFloorSet events raised by the Governance contract.
type GovernanceParticipationFloorSetIterator struct {
	Event *GovernanceParticipationFloorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceParticipationFloorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationFloorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceParticipationFloorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceParticipationFloorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationFloorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationFloorSet represents a ParticipationFloorSet event raised by the Governance contract.
type GovernanceParticipationFloorSet struct {
	ParticipationFloor *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterParticipationFloorSet is a free log retrieval operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) FilterParticipationFloorSet(opts *bind.FilterOpts) (*GovernanceParticipationFloorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationFloorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationFloorSetIterator{contract: _Governance.contract, event: "ParticipationFloorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationFloorSet is a free log subscription operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) WatchParticipationFloorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationFloorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationFloorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationFloorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationFloorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipationFloorSet is a log parse operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) ParseParticipationFloorSet(log types.Log) (*GovernanceParticipationFloorSet, error) {
	event := new(GovernanceParticipationFloorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationFloorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the Governance contract.
type GovernanceProposalApprovedIterator struct {
	Event *GovernanceProposalApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalApproved represents a ProposalApproved event raised by the Governance contract.
type GovernanceProposalApproved struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalApproved(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalApprovedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalApproved", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalApprovedIterator{contract: _Governance.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *GovernanceProposalApproved, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalApproved", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalApproved)
				if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalApproved is a log parse operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalApproved(log types.Log) (*GovernanceProposalApproved, error) {
	event := new(GovernanceProposalApproved)
	if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalDequeuedIterator is returned from FilterProposalDequeued and is used to iterate over the raw logs and unpacked data for ProposalDequeued events raised by the Governance contract.
type GovernanceProposalDequeuedIterator struct {
	Event *GovernanceProposalDequeued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalDequeuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalDequeued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalDequeued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalDequeuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalDequeuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalDequeued represents a ProposalDequeued event raised by the Governance contract.
type GovernanceProposalDequeued struct {
	ProposalId *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalDequeued is a free log retrieval operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) FilterProposalDequeued(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalDequeuedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalDequeued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalDequeuedIterator{contract: _Governance.contract, event: "ProposalDequeued", logs: logs, sub: sub}, nil
}

// WatchProposalDequeued is a free log subscription operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) WatchProposalDequeued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalDequeued, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalDequeued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalDequeued)
				if err := _Governance.contract.UnpackLog(event, "ProposalDequeued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalDequeued is a log parse operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) ParseProposalDequeued(log types.Log) (*GovernanceProposalDequeued, error) {
	event := new(GovernanceProposalDequeued)
	if err := _Governance.contract.UnpackLog(event, "ProposalDequeued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governance contract.
type GovernanceProposalExecutedIterator struct {
	Event *GovernanceProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecuted represents a ProposalExecuted event raised by the Governance contract.
type GovernanceProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutedIterator{contract: _Governance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalExecuted(log types.Log) (*GovernanceProposalExecuted, error) {
	event := new(GovernanceProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalExpiredIterator is returned from FilterProposalExpired and is used to iterate over the raw logs and unpacked data for ProposalExpired events raised by the Governance contract.
type GovernanceProposalExpiredIterator struct {
	Event *GovernanceProposalExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExpired represents a ProposalExpired event raised by the Governance contract.
type GovernanceProposalExpired struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExpired is a free log retrieval operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalExpired(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalExpiredIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExpired", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExpiredIterator{contract: _Governance.contract, event: "ProposalExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExpired is a free log subscription operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalExpired(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExpired, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExpired", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExpired)
				if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExpired is a log parse operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalExpired(log types.Log) (*GovernanceProposalExpired, error) {
	event := new(GovernanceProposalExpired)
	if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Governance contract.
type GovernanceProposalQueuedIterator struct {
	Event *GovernanceProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalQueued represents a ProposalQueued event raised by the Governance contract.
type GovernanceProposalQueued struct {
	ProposalId       *big.Int
	Proposer         common.Address
	TransactionCount *big.Int
	Deposit          *big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) FilterProposalQueued(opts *bind.FilterOpts, proposalId []*big.Int, proposer []common.Address) (*GovernanceProposalQueuedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalQueued", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalQueuedIterator{contract: _Governance.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalQueued, proposalId []*big.Int, proposer []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalQueued", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalQueued)
				if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) ParseProposalQueued(log types.Log) (*GovernanceProposalQueued, error) {
	event := new(GovernanceProposalQueued)
	if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalUpvoteRevokedIterator is returned from FilterProposalUpvoteRevoked and is used to iterate over the raw logs and unpacked data for ProposalUpvoteRevoked events raised by the Governance contract.
type GovernanceProposalUpvoteRevokedIterator struct {
	Event *GovernanceProposalUpvoteRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalUpvoteRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalUpvoteRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalUpvoteRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalUpvoteRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalUpvoteRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalUpvoteRevoked represents a ProposalUpvoteRevoked event raised by the Governance contract.
type GovernanceProposalUpvoteRevoked struct {
	ProposalId     *big.Int
	Account        common.Address
	RevokedUpvotes *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalUpvoteRevoked is a free log retrieval operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) FilterProposalUpvoteRevoked(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalUpvoteRevokedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalUpvoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalUpvoteRevokedIterator{contract: _Governance.contract, event: "ProposalUpvoteRevoked", logs: logs, sub: sub}, nil
}

// WatchProposalUpvoteRevoked is a free log subscription operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) WatchProposalUpvoteRevoked(opts *bind.WatchOpts, sink chan<- *GovernanceProposalUpvoteRevoked, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalUpvoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalUpvoteRevoked)
				if err := _Governance.contract.UnpackLog(event, "ProposalUpvoteRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalUpvoteRevoked is a log parse operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) ParseProposalUpvoteRevoked(log types.Log) (*GovernanceProposalUpvoteRevoked, error) {
	event := new(GovernanceProposalUpvoteRevoked)
	if err := _Governance.contract.UnpackLog(event, "ProposalUpvoteRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalUpvotedIterator is returned from FilterProposalUpvoted and is used to iterate over the raw logs and unpacked data for ProposalUpvoted events raised by the Governance contract.
type GovernanceProposalUpvotedIterator struct {
	Event *GovernanceProposalUpvoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalUpvotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalUpvoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalUpvoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalUpvotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalUpvotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalUpvoted represents a ProposalUpvoted event raised by the Governance contract.
type GovernanceProposalUpvoted struct {
	ProposalId *big.Int
	Account    common.Address
	Upvotes    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalUpvoted is a free log retrieval operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) FilterProposalUpvoted(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalUpvotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalUpvoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalUpvotedIterator{contract: _Governance.contract, event: "ProposalUpvoted", logs: logs, sub: sub}, nil
}

// WatchProposalUpvoted is a free log subscription operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) WatchProposalUpvoted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalUpvoted, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalUpvoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalUpvoted)
				if err := _Governance.contract.UnpackLog(event, "ProposalUpvoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalUpvoted is a log parse operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) ParseProposalUpvoted(log types.Log) (*GovernanceProposalUpvoted, error) {
	event := new(GovernanceProposalUpvoted)
	if err := _Governance.contract.UnpackLog(event, "ProposalUpvoted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Governance contract.
type GovernanceProposalVotedIterator struct {
	Event *GovernanceProposalVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVoted represents a ProposalVoted event raised by the Governance contract.
type GovernanceProposalVoted struct {
	ProposalId *big.Int
	Account    common.Address
	Value      *big.Int
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalVotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVotedIterator{contract: _Governance.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVoted, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVoted)
				if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVoted is a log parse operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) ParseProposalVoted(log types.Log) (*GovernanceProposalVoted, error) {
	event := new(GovernanceProposalVoted)
	if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceQueueExpirySetIterator is returned from FilterQueueExpirySet and is used to iterate over the raw logs and unpacked data for QueueExpirySet events raised by the Governance contract.
type GovernanceQueueExpirySetIterator struct {
	Event *GovernanceQueueExpirySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceQueueExpirySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceQueueExpirySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceQueueExpirySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceQueueExpirySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceQueueExpirySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceQueueExpirySet represents a QueueExpirySet event raised by the Governance contract.
type GovernanceQueueExpirySet struct {
	QueueExpiry *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQueueExpirySet is a free log retrieval operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) FilterQueueExpirySet(opts *bind.FilterOpts) (*GovernanceQueueExpirySetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "QueueExpirySet")
	if err != nil {
		return nil, err
	}
	return &GovernanceQueueExpirySetIterator{contract: _Governance.contract, event: "QueueExpirySet", logs: logs, sub: sub}, nil
}

// WatchQueueExpirySet is a free log subscription operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) WatchQueueExpirySet(opts *bind.WatchOpts, sink chan<- *GovernanceQueueExpirySet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "QueueExpirySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceQueueExpirySet)
				if err := _Governance.contract.UnpackLog(event, "QueueExpirySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQueueExpirySet is a log parse operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) ParseQueueExpirySet(log types.Log) (*GovernanceQueueExpirySet, error) {
	event := new(GovernanceQueueExpirySet)
	if err := _Governance.contract.UnpackLog(event, "QueueExpirySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceReferendumStageDurationSetIterator is returned from FilterReferendumStageDurationSet and is used to iterate over the raw logs and unpacked data for ReferendumStageDurationSet events raised by the Governance contract.
type GovernanceReferendumStageDurationSetIterator struct {
	Event *GovernanceReferendumStageDurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceReferendumStageDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceReferendumStageDurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceReferendumStageDurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceReferendumStageDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceReferendumStageDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceReferendumStageDurationSet represents a ReferendumStageDurationSet event raised by the Governance contract.
type GovernanceReferendumStageDurationSet struct {
	ReferendumStageDuration *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReferendumStageDurationSet is a free log retrieval operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) FilterReferendumStageDurationSet(opts *bind.FilterOpts) (*GovernanceReferendumStageDurationSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ReferendumStageDurationSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceReferendumStageDurationSetIterator{contract: _Governance.contract, event: "ReferendumStageDurationSet", logs: logs, sub: sub}, nil
}

// WatchReferendumStageDurationSet is a free log subscription operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) WatchReferendumStageDurationSet(opts *bind.WatchOpts, sink chan<- *GovernanceReferendumStageDurationSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ReferendumStageDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceReferendumStageDurationSet)
				if err := _Governance.contract.UnpackLog(event, "ReferendumStageDurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferendumStageDurationSet is a log parse operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) ParseReferendumStageDurationSet(log types.Log) (*GovernanceReferendumStageDurationSet, error) {
	event := new(GovernanceReferendumStageDurationSet)
	if err := _Governance.contract.UnpackLog(event, "ReferendumStageDurationSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Governance contract.
type GovernanceRegistrySetIterator struct {
	Event *GovernanceRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRegistrySet represents a RegistrySet event raised by the Governance contract.
type GovernanceRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Governance *GovernanceFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*GovernanceRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRegistrySetIterator{contract: _Governance.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Governance *GovernanceFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *GovernanceRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRegistrySet)
				if err := _Governance.contract.UnpackLog(event, "RegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Governance *GovernanceFilterer) ParseRegistrySet(log types.Log) (*GovernanceRegistrySet, error) {
	event := new(GovernanceRegistrySet)
	if err := _Governance.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}
