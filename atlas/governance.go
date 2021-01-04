package atlas

import (
	"math/big"

	"context"

	"log"
	"strings"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ProposalApprovedEvent corresponds to event ProposalApproved(uint256 indexed proposalId); */
type ProposalApprovedEvent struct {
	ProposalId  common.Hash
	BlockNumber *big.Int
}

// GetProposalApprovedEvents fetches data for all Proposal Approved events between (and including) the two specified block numbers
func GetProposalApprovedEvents(fromBlock *big.Int, toBlock *big.Int) []*ProposalApprovedEvent {

	var logProposalApprovedSig = []byte("ProposalApproved(uint256)")
	var logProposalApprovedSigHash = crypto.Keccak256Hash(logProposalApprovedSig)
	var TopicsFilter = [][]common.Hash{{logProposalApprovedSigHash}}

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Governance])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    TopicsFilter,

		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	approvedProposals := make([]*ProposalApprovedEvent, 0, len(logs))

	for _, vLog := range logs {

		var approvedProposal ProposalApprovedEvent
		approvedProposal.ProposalId = vLog.Topics[1]
		approvedProposal.BlockNumber = new(big.Int).SetUint64(vLog.BlockNumber)

		approvedProposals = append(approvedProposals, &approvedProposal)
	}

	return approvedProposals
}

//ProposalQueuedEvent corresponds to event
/*
 event ProposalQueued(
    uint256 indexed proposalId,
    address indexed proposer,
    uint256 transactionCount,
    uint256 deposit,
    uint256 timestamp
  );
*/
type ProposalQueuedEvent struct {
	ProposalId       common.Hash
	Proposer         common.Address
	TransactionCount *big.Int
	Deposit          *big.Int
	Timestamp        *big.Int
	BlockNumber      *big.Int
}

// GetProposalQueuedEvents fetches data of all Proposal Queued events between (and including) the two specified block numbers
func GetProposalQueuedEvents(fromBlock *big.Int, toBlock *big.Int) []*ProposalQueuedEvent {

	var logProposalQueuedSig = []byte("ProposalQueued(uint256,address,uint256,uint256,uint256)")
	var logProposalQueuedSigHash = crypto.Keccak256Hash(logProposalQueuedSig)
	var TopicsFilter = [][]common.Hash{{logProposalQueuedSigHash}}

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Governance])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    TopicsFilter,

		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	queuedProposals := make([]*ProposalQueuedEvent, 0, len(logs))

	contractAbi, err := abi.JSON(strings.NewReader(string(binding.GovernanceABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		var queuedProposal ProposalQueuedEvent
		err := contractAbi.Unpack(&queuedProposal, "ProposalQueued", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}

		queuedProposal.ProposalId = vLog.Topics[1]
		queuedProposal.Proposer = common.HexToAddress(vLog.Topics[2].Hex())
		queuedProposal.BlockNumber = new(big.Int).SetUint64(vLog.BlockNumber)

		queuedProposals = append(queuedProposals, &queuedProposal)
	}

	return queuedProposals
}

//GetExecutionStatusOfProposal checks if a proposal has been executed upto a block number
func GetExecutionStatusOfProposal(proposal *ProposalQueuedEvent, uptoBlockNumber *big.Int) bool {

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Governance])
	//Search through all ProposalExecuted logs in the block number range (when proposal was queued, uptoBlockNumber)
	//for this proposalID
	var logProposalExecutedSig = []byte("ProposalExecuted(uint256)")
	var logProposalExecutedSigHash = crypto.Keccak256Hash(logProposalExecutedSig)

	// Will filter through topics for which
	// Topics[0]  == logProposalExecutedSigHash (ProposalExecuted)
	// AND (Note : Confirmed that the following syntax for TopicsFilter represents an "AND" operator between Topics[0] and Topics[1] filter)
	// Topics[1] == proposalID
	var TopicsFilter = [][]common.Hash{{logProposalExecutedSigHash}, {proposal.ProposalId}}

	query := ethereum.FilterQuery{
		FromBlock: proposal.BlockNumber,
		ToBlock:   uptoBlockNumber,
		Topics:    TopicsFilter,

		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)

	if err != nil {
		log.Printf("Error retrieving ProposalExecuted Logs for Proposal ID #%s at Block Number %s \n", proposal.ProposalId.Big(), uptoBlockNumber)
		log.Printf("Hence reporting executed status for this proposal as false")
		log.Print(err)
		return false
	}

	if len(logs) > 1 {
		log.Println("Something went wrong in Governance Atlas")
		log.Printf(" - as it is observing more than 1 ProposalExecuted log for Proposal #%s\n", proposal.ProposalId.Big())
	}

	return len(logs) > 0

}
