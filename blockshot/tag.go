package blockshot

import (
	"database/sql"
	"strings"

	"github.com/ChorusOne/Icarus/blockchain"
	. "github.com/ChorusOne/Icarus/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"log"
	"math/big"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	electionABI          *abi.ABI
	governanceABI        *abi.ABI
	governanceSlasherABI *abi.ABI
	lockedGoldABI        *abi.ABI
	stableTokenABI       *abi.ABI
	goldTokenABI         *abi.ABI
)

// GenerateTransactionTags is a block iterator task to tag transactions and
// store these tags to the database. The data source for this task is the output
// of the ProcessTransactions task. Hence, in the list of tasks passed to the
// BlockIterator, this task should NOT appear in the list before the ProcessTransactions task.
func GenerateTransactionTags(db *sql.DB, blockNumber *big.Int) error {

	hashes, err := GetHashOfTransactions(db, blockNumber)

	if err != nil {
		return err
	}

	for _, txHash := range hashes {
		getAndProcessReceipt(db, blockNumber, txHash)
	}

	return nil

}

func getAndProcessReceipt(db *sql.DB, blockNumber *big.Int, txHash string) error {

	receipt, err := blockchain.GetReceiptLite(rpcClient, txHash)

	if err != nil {
		return err
	}
	tags := ProcessTags(receipt.Logs)

	DumpTagsRow(db, blockNumber, txHash, receipt.From, receipt.To, EventLogs(receipt.Logs), Tags(tags))

	return nil

}

//ProcessTags generates tags for a transaction from the logs provided
func ProcessTags(logs []*types.Log) []*Tag {

	tags := make([]*Tag, 0)

	for _, vLog := range logs {

		var tag *Tag

		//Switch By Contract Name First
		//This is important to avoid clashes with similarly named event signatures across contracts
		// Eg. "Transfer(address,address,uint256)" in GoldToken.sol and StableToken.sol
		switch vLog.Address.Hex() {

		//Downtime Slasher
		case WrapperContractDeploymentAddress[NetActive][DowntimeSlasher]:
			tag = generateDowntimeSlasherTag(vLog)

		// Election
		case WrapperContractDeploymentAddress[NetActive][Election]:
			tag = generateElectionTag(vLog)

		//Governance
		case WrapperContractDeploymentAddress[NetActive][Governance]:
			tag = generateGovernanceTag(vLog)

		//GovernanceSlasher
		case WrapperContractDeploymentAddress[NetActive][GovernanceSlasher]:
			tag = generateGovernanceSlasherTag(vLog)

		//LockedGold
		case WrapperContractDeploymentAddress[NetActive][LockedGold]:
			tag = generateLockedGoldTag(vLog)

		//StableToken
		case WrapperContractDeploymentAddress[NetActive][StableToken]:
			tag = generateStableTokenTag(vLog)

		//GoldToken
		case WrapperContractDeploymentAddress[NetActive][GoldToken]:
			tag = generateGoldTokenTag(vLog)

		}

		if tag != nil {
			tags = append(tags, tag)
		}

	}

	return tags

}

//TopicHashHex returns Keccak256Hash of the topic signature in string format
func TopicHashHex(signature string) string {
	return crypto.Keccak256Hash([]byte(signature)).Hex()
}

// List of all relevant Topic Signatures in Keccak256 Hash Format
var (
	//DowntimeSlasher.sol
	DowntimeSlashPerformedHash = TopicHashHex("DowntimeSlashPerformed(address,uint256)")

	//Election.sol

	ValidatorGroupMarkedEligibleHash   = TopicHashHex("ValidatorGroupMarkedEligible(address)")
	ValidatorGroupMarkedIneligibleHash = TopicHashHex("ValidatorGroupMarkedIneligible(address)")

	ValidatorGroupVoteCastHash           = TopicHashHex("ValidatorGroupVoteCast(address,address,uint256)")
	ValidatorGroupPendingVoteRevokedHash = TopicHashHex("ValidatorGroupPendingVoteRevoked(address,address,uint256)")

	ValidatorGroupVoteActivatedHash     = TopicHashHex("ValidatorGroupVoteActivated(address,address,uint256,uint256)")
	ValidatorGroupActiveVoteRevokedHash = TopicHashHex("ValidatorGroupActiveVoteRevoked(address,address,uint256,uint256)")

	//Governance.sol

	ProposalQueuedHash        = TopicHashHex("ProposalQueued(uint256,address,uint256,uint256,uint256)")
	ProposalUpvotedHash       = TopicHashHex("ProposalUpvoted(uint256,address,uint256)")
	ProposalUpvoteRevokedHash = TopicHashHex("ProposalUpvoteRevoked(uint256,address,uint256)")
	ProposalDequeuedHash      = TopicHashHex("ProposalDequeued(uint256,uint256)")
	ProposalApprovedHash      = TopicHashHex("ProposalApproved(uint256)")
	ProposalVotedHash         = TopicHashHex("ProposalVoted(uint256,address,uint256,uint256)")
	ProposalExecutedHash      = TopicHashHex("ProposalExecuted(uint256)")
	ProposalExpiredHash       = TopicHashHex("ProposalExpired(uint256)")

	//GovernanceSlasher.sol
	SlashingApprovedHash         = TopicHashHex("SlashingApproved(address,uint256)")
	GovernanceSlashPerformedHash = TopicHashHex("GovernanceSlashPerformed(address,uint256)")

	//LockedGold.sol

	GoldLockedHash     = TopicHashHex("GoldLocked(address,uint256)")
	GoldUnlockedHash   = TopicHashHex("GoldUnlocked(address,uint256,uint256)")
	GoldRelockedHash   = TopicHashHex("GoldRelocked(address,uint256)")
	GoldWithdrawnHash  = TopicHashHex("GoldWithdrawn(address,uint256)")
	AccountSlashedHash = TopicHashHex("AccountSlashed(address,uint256,address,uint256)")

	//StableToken.sol

	TransferHash        = TopicHashHex("Transfer(address,address,uint256)")
	TransferCommentHash = TopicHashHex("TransferComment(string)")

	//GoldToken.sol

	//Transfer and TransferComment - same signatures as StableToken.sol
	ApprovalHash = TopicHashHex("Approval(address,address,uint256)")
)

//Tag Generators

func generateDowntimeSlasherTag(vLog *types.Log) *Tag {
	var tag *Tag

	//Using if-then, because there is only onçe DowntimeSlasher event of note
	//Use switch-case if there are multiple events of note
	//Should return nil as the default case

	if vLog.Topics[0].Hex() == DowntimeSlashPerformedHash {

		tag = NewTag(
			"DowntimeSlashPerformed",
			"Downtime Slash Performed",
			"DowntimeSlasher")
		tag.Parameters["validator"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["startBlock"] = hashToBigIntString(vLog.Topics[2])
	}

	return tag

}

func generateElectionTag(vLog *types.Log) *Tag {

	//Lazy-load electionABI on first use
	if electionABI == nil {

		abi, err := abi.JSON(strings.NewReader(string(binding.ElectionABI)))
		if err != nil {
			log.Fatal(err)
		}

		electionABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case ValidatorGroupMarkedEligibleHash:
		tag = NewTag(
			"ValidatorGroupMarkedEligible",
			"Validator Group Marked Eligible",
			"Election")
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[1])

	case ValidatorGroupMarkedIneligibleHash:
		tag = NewTag(
			"ValidatorGroupMarkedIneligible",
			"Validator Group Marked Ineligible",
			"Election")
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[1])

	case ValidatorGroupVoteCastHash:

		var logVC LogValidatorGroupVoteCast
		err := electionABI.Unpack(&logVC, "ValidatorGroupVoteCast", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ValidatorGroupVoteCast",
			"Validator Group Vote Cast",
			"Election")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["value"] = logVC.Value.String()

	case ValidatorGroupPendingVoteRevokedHash:

		var logPVR LogValidatorGroupPendingVoteRevoked
		err := electionABI.Unpack(&logPVR, "ValidatorGroupPendingVoteRevoked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ValidatorGroupPendingVoteRevoked",
			"Validator Group Pending Vote Revoked",
			"Election")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["value"] = logPVR.Value.String()

	case ValidatorGroupVoteActivatedHash:

		var logVA LogValidatorGroupVoteActivated
		err := electionABI.Unpack(&logVA, "ValidatorGroupVoteActivated", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ValidatorGroupVoteActivated",
			"Validator Group Vote Activated",
			"Election")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["value"] = logVA.Value.String()
		tag.Parameters["units"] = logVA.Units.String()

	case ValidatorGroupActiveVoteRevokedHash:

		var logAVR LogValidatorGroupActiveVoteRevoked
		err := electionABI.Unpack(&logAVR, "ValidatorGroupActiveVoteRevoked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		tag = NewTag(
			"ValidatorGroupActiveVoteRevoked",
			"Validator Group Active Vote Revoked",
			"Election")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["group"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["value"] = logAVR.Value.String()
		tag.Parameters["units"] = logAVR.Units.String()

	}

	return tag
}

func generateGovernanceTag(vLog *types.Log) *Tag {

	if governanceABI == nil {
		abi, err := abi.JSON(strings.NewReader(string(binding.GovernanceABI)))
		if err != nil {
			log.Fatal(err)
		}

		governanceABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case ProposalQueuedHash:

		var logPQ LogProposalQueued
		err := governanceABI.Unpack(&logPQ, "ProposalQueued", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ProposalQueued",
			"Proposal Queued",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])
		tag.Parameters["proposer"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[2])

		tag.Parameters["deposit"] = logPQ.Deposit.String()
		tag.Parameters["timestamp"] = logPQ.Timestamp.String()

	case ProposalUpvotedHash:

		var logPU LogProposalUpvoted
		err := governanceABI.Unpack(&logPU, "ProposalUpvoted", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ProposalUpvoted",
			"Proposal Upvoted",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["upvotes"] = logPU.Upvotes.String()

	case ProposalUpvoteRevokedHash:

		var logPUR LogProposalUpvoteRevoked

		err := governanceABI.Unpack(&logPUR, "ProposalUpvoteRevoked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ProposalUpvoteRevoked",
			"Proposal Upvote Revoked",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[2])

		tag.Parameters["revokedUpvotes"] = logPUR.RevokedUpvotes.String()

	case ProposalDequeuedHash:

		var logPD LogProposalDequeued
		err := governanceABI.Unpack(&logPD, "ProposalDequeued", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ProposalDequeued",
			"Proposal Dequeued",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])
		tag.Parameters["timestamp"] = logPD.Timestamp.String()

	case ProposalApprovedHash:
		tag = NewTag(
			"ProposalApproved",
			"Proposal Approved",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])

	case ProposalVotedHash:

		var logPV LogProposalVoted
		err := governanceABI.Unpack(&logPV, "ProposalVoted", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"ProposalVoted",
			"Proposal Voted",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[2])

		tag.Parameters["value"] = logPV.Value.String()
		tag.Parameters["weight"] = logPV.Weight.String()

	case ProposalExecutedHash:
		tag = NewTag(
			"ProposalExecuted",
			"Proposal Executed",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])

	case ProposalExpiredHash:
		tag = NewTag(
			"ProposalExpired",
			"Proposal Expired",
			"Governance")
		tag.Parameters["proposalId"] = hashToBigIntString(vLog.Topics[1])

	}

	return tag
}

func generateGovernanceSlasherTag(vLog *types.Log) *Tag {

	if governanceSlasherABI == nil {

		abi, err := abi.JSON(strings.NewReader(string(binding.GovernanceSlasherABI)))
		if err != nil {
			log.Fatal(err)
		}

		governanceSlasherABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case SlashingApprovedHash:

		var logSA LogSlashingApproved
		err := governanceSlasherABI.Unpack(&logSA, "SlashingApproved", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"SlashingApproved",
			"Slashing Approved",
			"GovernanceSlasher")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["amount"] = logSA.Amount.String()

	case GovernanceSlashPerformedHash:
		var logGSP LogGovernanceSlashPerformed
		err := governanceSlasherABI.Unpack(&logGSP, "GovernanceSlashPerformed", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"GovernanceSlashPerformed",
			"Goernance Slash Performed",
			"GovernanceSlasher")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["amount"] = logGSP.Amount.String()

	}

	return tag
}

func generateLockedGoldTag(vLog *types.Log) *Tag {

	if lockedGoldABI == nil {

		abi, err := abi.JSON(strings.NewReader(string(binding.LockedGoldABI)))
		if err != nil {
			log.Fatal(err)
		}

		lockedGoldABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case GoldLockedHash:

		var logGL LogGoldLocked
		err := lockedGoldABI.Unpack(&logGL, "GoldLocked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		tag = NewTag(
			"GoldLocked",
			"Gold Locked",
			"LockedGold")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["value"] = logGL.Value.String()

	case GoldUnlockedHash:

		var logGU LogGoldUnlocked
		err := lockedGoldABI.Unpack(&logGU, "GoldUnlocked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"GoldUnlocked",
			"Gold Unlocked",
			"LockedGold")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["value"] = logGU.Value.String()
		tag.Parameters["available"] = logGU.Available.String()

	case GoldRelockedHash:

		var logGR LogGoldRelocked
		err := lockedGoldABI.Unpack(&logGR, "GoldRelocked", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"GoldRelocked",
			"Gold Relocked",
			"LockedGold")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["value"] = logGR.Value.String()

	case GoldWithdrawnHash:

		var logGW LogGoldWithdrawn
		err := lockedGoldABI.Unpack(&logGW, "GoldWithdrawn", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"GoldWithdrawn",
			"Gold Withdrawn",
			"LockedGold")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["value"] = logGW.Value.String()

	case AccountSlashedHash:
		var logAS LogAccountSlashed
		err := lockedGoldABI.Unpack(&logAS, "AccountSlashed", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"AccountSlashed",
			"Account Slashed",
			"LockedGold")
		tag.Parameters["account"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["reporter"] = hashToAddressString(vLog.Topics[2])

		tag.Parameters["penalty"] = logAS.Penalty.String()
		tag.Parameters["reward"] = logAS.Reward.String()

	}

	return tag
}

func generateStableTokenTag(vLog *types.Log) *Tag {

	if stableTokenABI == nil {

		abi, err := abi.JSON(strings.NewReader(string(binding.StableTokenABI)))
		if err != nil {
			log.Fatal(err)
		}

		stableTokenABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case TransferHash:

		var logT LogTransfer
		err := stableTokenABI.Unpack(&logT, "Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"Transfer",
			"Stable Token Transfer",
			"StableToken")
		tag.Parameters["from"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["to"] = hashToAddressString(vLog.Topics[2])
		tag.Parameters["value"] = logT.Value.String()

	case TransferCommentHash:

		var logTC LogTransferComment
		err := stableTokenABI.Unpack(&logTC, "TransferComment", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"TransferComment",
			"Stable Token Transfer Comment",
			"StableToken")
		tag.Parameters["comment"] = logTC.Comment

	}

	return tag
}

func generateGoldTokenTag(vLog *types.Log) *Tag {

	if goldTokenABI == nil {

		abi, err := abi.JSON(strings.NewReader(string(binding.GoldTokenABI)))
		if err != nil {
			log.Fatal(err)
		}
		goldTokenABI = &abi
	}

	var tag *Tag

	switch vLog.Topics[0].Hex() {

	case TransferHash:

		var logT LogTransfer
		err := goldTokenABI.Unpack(&logT, "Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"Transfer",
			"Gold Token Transfer",
			"GoldToken")
		tag.Parameters["from"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["to"] = hashToAddressString(vLog.Topics[2])

		tag.Parameters["value"] = logT.Value.String()

	case TransferCommentHash:

		var logTC LogTransferComment
		err := goldTokenABI.Unpack(&logTC, "TransferComment", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		tag = NewTag(
			"TransferComment",
			"Gold Token Transfer Comment",
			"GoldToken")

		tag.Parameters["comment"] = logTC.Comment

	case ApprovalHash:

		var logA LogApproval
		err := goldTokenABI.Unpack(&logA, "Approval", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		tag = NewTag(
			"Approval",
			"Spender Approval",
			"GoldToken")
		tag.Parameters["owner"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["spender"] = hashToAddressString(vLog.Topics[1])
		tag.Parameters["value"] = logA.Value.String()

	}

	return tag
}

// Helper functions to remove extraneous zero padding in addresses and integers

func hashToAddressString(hash common.Hash) string {
	return common.HexToAddress(hash.String()).String()
}

func hashToBigIntString(hash common.Hash) string {
	return hash.Big().String()
}
