package contract

import (
	"github.com/ChorusOne/Icarus/atlas"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var governanceContractInstance *binding.Governance

// enum Stage { None, Queued, Approval, Referendum, Execution, Expiration }
const (
	noneStageEnumValue       uint8 = 0
	queuedStageEnumValue     uint8 = 1
	approvalStageEnumValue   uint8 = 2
	referendumStageEnumValue uint8 = 3
	executionStageEnumValue  uint8 = 4
	expirationStageEnumValue uint8 = 5
)

var StageMapping = map[uint8]string{
	noneStageEnumValue:       "None",
	queuedStageEnumValue:     "Queued",
	approvalStageEnumValue:   "Approval",
	referendumStageEnumValue: "Referendum",
	executionStageEnumValue:  "Execution",
	expirationStageEnumValue: "Expiration",
}

func InitGovernanceContractInstance() {

	instance, err := binding.NewGovernance(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Governance]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Governance Contract is loaded")
	governanceContractInstance = instance

}

func GetQueue(atBlockNumber *big.Int) ([]*big.Int, []*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	//Return zero value if polled before contract deployment
	if !check {
		blankSlice := []*big.Int{}
		return blankSlice, blankSlice, nil
	}

	ids, upVotes, err := governanceContractInstance.GetQueue(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		return nil, nil, err
	}

	return ids, upVotes, nil

}

func GetDequeue(atBlockNumber *big.Int) ([]*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	//Return zero value if polled before contract deployment
	if !check {
		blankSlice := []*big.Int{}
		return blankSlice, nil
	}

	dequeuedProposalIds, err := governanceContractInstance.GetDequeue(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		return nil, err
	}
	return dequeuedProposalIds, nil

}

func GetReferendumProposals(atBlockNumber *big.Int) ([]*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	//Return zero value if polled before contract deployment
	if !check {
		blankSlice := []*big.Int{}
		return blankSlice, nil
	}

	dequeuedProposalIds, err := GetDequeue(atBlockNumber)

	if err != nil {
		return nil, err
	}

	referendumProposalIds := make([]*big.Int, 0, len(dequeuedProposalIds))

	for _, id := range dequeuedProposalIds {
		stage, err := GetProposalStage(id, atBlockNumber)

		if err != nil {
			return nil, err
		}
		log.Printf("Stage of Proposal #%s is %d\n", id, stage)

		if stage == referendumStageEnumValue {
			referendumProposalIds = append(referendumProposalIds, id)
		}

	}

	return referendumProposalIds, nil

}

func GetProposalStage(proposalID *big.Int, atBlockNumber *big.Int) (uint8, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	if !check {
		return 0, nil
	}

	stage, err := governanceContractInstance.GetProposalStage(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

	if err != nil {
		// enum Stage { None, Queued, Approval, Referendum, Execution, Expiration }

		return 255, err // Hacky - 255 a "nil" equivalent for uint8
	}

	return stage, nil
}

// GetDetailedReferendumProposals returns list of all referendum proposals (with their respective details) at the provided block height
func GetDetailedReferendumProposals(atBlockNumber *big.Int) ([]*DetailedProposal, error) {

	result := []*DetailedProposal{}

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	if !check {
		return result, nil
	}

	approvalStageDuration, err1 := governanceContractInstance.GetApprovalStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err1 != nil {
		log.Println(err1)
		return nil, err1
	}

	referendumStageDuration, err2 := governanceContractInstance.GetReferendumStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err2 != nil {
		log.Println(err2)
		return nil, err2
	}

	executionStageDuration, err3 := governanceContractInstance.GetExecutionStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err3 != nil {
		log.Println(err3)
		return nil, err3
	}

	dequeuedProposalIds, err := GetDequeue(atBlockNumber)

	if err != nil {
		return nil, err
	}

	for idx, proposalID := range dequeuedProposalIds {

		//Check if proposal is not in Referendum stage
		stage, err := GetProposalStage(proposalID, atBlockNumber)

		if err != nil {
			return nil, err
		}
		log.Printf("Stage of Proposal #%s is %d\n", proposalID, stage)

		if stage != referendumStageEnumValue {
			continue // Not a Referendum ProposalID, so skip adding to list of referendum proposals
		}

		index := big.NewInt(int64(idx))
		yesVotes, noVotes, abstainVotes, err1 := governanceContractInstance.GetVoteTotals(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

		if err1 != nil {
			log.Println(err1)
			return nil, err1
		}

		proposer, deposit, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

		if err2 != nil {
			log.Println(err2)
			return nil, err2
		}

		referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
			propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

		detailedProposal := &DetailedProposal{
			ProposalID:   proposalID,
			Index:        index,
			BlockNumber:  atBlockNumber,
			Stage:        "Referendum",
			Proposer:     proposer,
			YesVotes:     yesVotes,
			NoVotes:      noVotes,
			AbstainVotes: abstainVotes,
			Description:  description,

			ProposalEpoch:   propEpoch,
			ReferendumEpoch: referendumEpoch,
			ExecutionEpoch:  executionEpoch,
			ExpirationEpoch: expirationEpoch,
			Deposit:         deposit,
		}

		result = append(result, detailedProposal)

	}

	return result, nil

}

// GetDetailedQueuedProposals returns list of all queued proposals (with their respective details) at the provided block height
func GetDetailedQueuedProposals(atBlockNumber *big.Int) ([]*DetailedQueuedProposal, error) {

	result := []*DetailedQueuedProposal{}

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	if !check {
		return result, nil
	}

	approvalStageDuration, err1 := governanceContractInstance.GetApprovalStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err1 != nil {
		log.Println(err1)
		return nil, err1
	}

	referendumStageDuration, err2 := governanceContractInstance.GetReferendumStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err2 != nil {
		log.Println(err2)
		return nil, err2
	}

	executionStageDuration, err3 := governanceContractInstance.GetExecutionStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err3 != nil {
		log.Println(err3)
		return nil, err3
	}

	queuedProposalIds, upvotes, err := GetQueue(atBlockNumber)

	if err != nil {
		return nil, err
	}

	for idx, proposalID := range queuedProposalIds {

		index := big.NewInt(int64(idx))
		upvotesForProposal := upvotes[idx]

		stage, err := GetProposalStage(proposalID, atBlockNumber)
		if err != nil {
			return nil, err
		}

		proposer, _, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

		if err2 != nil {
			log.Println(err2)
			return nil, err2
		}
		referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
			propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

		detailedQueuedProposal := &DetailedQueuedProposal{
			ProposalID:  proposalID,
			Index:       index,
			BlockNumber: atBlockNumber,
			Stage:       StageMapping[stage],
			Proposer:    proposer,
			Upvotes:     upvotesForProposal,
			Description: description,

			ProposalEpoch:   propEpoch,
			ReferendumEpoch: referendumEpoch,
			ExecutionEpoch:  executionEpoch,
			ExpirationEpoch: expirationEpoch,
		}

		result = append(result, detailedQueuedProposal)

	}

	return result, nil

}

// GetAllProposalsUptoBlockNumber returns  all proposals on Celo blockchain upto block height with details as-on-block-height
func GetAllProposalsUptoBlockNumber(atBlockNumber *big.Int) (*AssortedProposals, error) {

	log.Println("Fetching all proposals at block height : ", atBlockNumber)

	result := &AssortedProposals{
		QueuedProposals:     []*DetailedQueuedProposal{},
		ApprovalProposals:   []*DetailedProposal{},
		ReferendumProposals: []*DetailedProposal{},
		ExecutionProposals:  []*DetailedProposal{},
		ExpiredProposals:    []*DetailedExpiredProposal{},
	}

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Governance])

	if !check {
		return result, nil
	}

	approvalStageDuration, err1 := governanceContractInstance.GetApprovalStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err1 != nil {
		log.Println(err1)
		return result, err1
	}

	referendumStageDuration, err2 := governanceContractInstance.GetReferendumStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err2 != nil {
		log.Println(err2)
		return result, err2
	}

	executionStageDuration, err3 := governanceContractInstance.GetExecutionStageDuration(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err3 != nil {
		log.Println(err3)
		return result, err3
	}

	upVotesMapForQPS, indexMapForQPS, err4 := getUpvotesAndIndexMapsForQueuedProposals(atBlockNumber)

	if err4 != nil {
		log.Println(err4)
		return result, err4
	}

	indexMapForDQPS, err5 := getIndexMapForDequeuedProposals(atBlockNumber)

	if err5 != nil {
		log.Println(err5)
		return result, err5
	}

	allProposals := atlas.GetProposalQueuedEvents(big.NewInt(0), atBlockNumber)

	for _, proposal := range allProposals {
		proposalID := proposal.ProposalId.Big()
		stage, err := GetProposalStage(proposalID, atBlockNumber)

		if err != nil {
			log.Printf("This shouldn't happen - couldn't get stage of Proposal #ID %s, so skipping this proposal \n", proposalID)
			log.Println(err)
			continue
		}

		switch stage {

		case noneStageEnumValue:
			continue

		case queuedStageEnumValue:

			proposer, _, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

			if err2 != nil {
				log.Printf("Was not able to retrive proposer, description and timestamp of Queued Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err2)
				continue
			}

			referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
				propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

			queuedProposal := &DetailedQueuedProposal{
				ProposalID:  proposalID,
				Index:       indexMapForQPS[proposalID.String()],
				BlockNumber: atBlockNumber,
				Stage:       StageMapping[stage],
				Proposer:    proposer,
				Upvotes:     upVotesMapForQPS[proposalID.String()],
				Description: description,

				ProposalEpoch:   propEpoch,
				ReferendumEpoch: referendumEpoch,
				ExecutionEpoch:  executionEpoch,
				ExpirationEpoch: expirationEpoch,

				QueuedAtBlockNumber: proposal.BlockNumber,
				Deposit:             proposal.Deposit,
				QueuedTimestamp:     proposal.Timestamp,
			}

			result.QueuedProposals = append(result.QueuedProposals, queuedProposal)

		case approvalStageEnumValue:

			proposer, _, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

			if err2 != nil {
				log.Printf("Was not able to retrive proposer, description and timestamp of Approval Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err2)
				continue
			}

			referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
				propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

			approvalProposal := &DetailedProposal{
				ProposalID:  proposalID,
				Index:       indexMapForDQPS[proposalID.String()],
				BlockNumber: atBlockNumber,
				Stage:       StageMapping[stage],
				Description: description,

				Proposer:            proposer,
				QueuedAtBlockNumber: proposal.BlockNumber,
				Deposit:             proposal.Deposit,
				QueuedTimestamp:     proposal.Timestamp,

				ProposalEpoch:   propEpoch,
				ReferendumEpoch: referendumEpoch,
				ExecutionEpoch:  executionEpoch,
				ExpirationEpoch: expirationEpoch,
			}

			result.ApprovalProposals = append(result.ApprovalProposals, approvalProposal)

		case referendumStageEnumValue:

			yesVotes, noVotes, abstainVotes, err1 := governanceContractInstance.GetVoteTotals(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

			if err1 != nil {
				log.Printf("Was not able to retrive votes of Referendum Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err1)
				continue
			}

			proposer, _, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)

			if err2 != nil {
				log.Printf("Was not able to retrive proposer, description and timestamp of Referendum Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err2)
				continue
			}

			referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
				propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

			referendumProposal := &DetailedProposal{
				ProposalID:  proposalID,
				Index:       indexMapForDQPS[proposalID.String()],
				BlockNumber: atBlockNumber,
				Stage:       StageMapping[stage],
				Description: description,

				YesVotes:     yesVotes,
				NoVotes:      noVotes,
				AbstainVotes: abstainVotes,

				Proposer:            proposer,
				QueuedAtBlockNumber: proposal.BlockNumber,
				Deposit:             proposal.Deposit,
				QueuedTimestamp:     proposal.Timestamp,

				ProposalEpoch:   propEpoch,
				ReferendumEpoch: referendumEpoch,
				ExecutionEpoch:  executionEpoch,
				ExpirationEpoch: expirationEpoch,
			}

			result.ReferendumProposals = append(result.ReferendumProposals, referendumProposal)

		case executionStageEnumValue:

			yesVotes, noVotes, abstainVotes, err1 := governanceContractInstance.GetVoteTotals(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)


			if err1 != nil {
				log.Printf("Was not able to retrive votes of Execution Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err1)
				continue
			}

			proposer, _, propEpoch, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: atBlockNumber}, proposalID)


			if err2 != nil {
				log.Printf("Was not able to retrive proposer, description and timestamp of Execution Proposal #ID %s, so skipping this proposal \n", proposalID)
				log.Println(err2)
				continue
			}

			referendumEpoch, executionEpoch, expirationEpoch := calculateEpochs(
				propEpoch, approvalStageDuration, referendumStageDuration, executionStageDuration)

			executionProposal := &DetailedProposal{
				ProposalID:  proposalID,
				Index:       indexMapForDQPS[proposalID.String()],
				BlockNumber: atBlockNumber,
				Stage:       StageMapping[stage],
				Description: description,

				YesVotes:     yesVotes,
				NoVotes:      noVotes,
				AbstainVotes: abstainVotes,

				Proposer:            proposer,
				QueuedAtBlockNumber: proposal.BlockNumber,
				Deposit:             proposal.Deposit,
				QueuedTimestamp:     proposal.Timestamp,

				ProposalEpoch:   propEpoch,
				ReferendumEpoch: referendumEpoch,
				ExecutionEpoch:  executionEpoch,
				ExpirationEpoch: expirationEpoch,
			}

			result.ExecutionProposals = append(result.ExecutionProposals, executionProposal)

		case expirationStageEnumValue:

			//No description exists in the smart contract for this expired proposal at the current block height
			//Hence fetching Description by referring to the block number at which this expired proposal was queued
			_, _, _, _, description, err2 := governanceContractInstance.GetProposal(&bind.CallOpts{BlockNumber: proposal.BlockNumber}, proposalID)

			if err2 != nil {
				log.Printf("Couldn't fetch description for expired proposal ID # %s\n, putting in a blank description", proposalID)
				log.Println("Error", err2)
				description = ""
			}

			expiredProposal := &DetailedExpiredProposal{
				ProposalID:          proposalID,
				BlockNumber:         atBlockNumber,
				Stage:               StageMapping[stage],
				Proposer:            proposal.Proposer,
				QueuedAtBlockNumber: proposal.BlockNumber,
				Deposit:             proposal.Deposit,
				Description:         description,
				QueuedTimestamp:     proposal.Timestamp,
				Executed:            atlas.GetExecutionStatusOfProposal(proposal, atBlockNumber),
			}

			result.ExpiredProposals = append(result.ExpiredProposals, expiredProposal)
		}

	}

	return result, nil

}

func getUpvotesAndIndexMapsForQueuedProposals(atBlockNumber *big.Int) (UpvotesMap, IndexMap, error) {

	upvotesMap := make(UpvotesMap)
	indexMap := make(IndexMap)

	queuedProposalIds, upvotes, err := GetQueue(atBlockNumber)

	if err != nil {
		return upvotesMap, indexMap, err
	}

	for idx, proposalID := range queuedProposalIds {
		upvotesMap[proposalID.String()] = upvotes[idx]
		indexMap[proposalID.String()] = big.NewInt(int64(idx))
	}

	return upvotesMap, indexMap, nil

}

func getIndexMapForDequeuedProposals(atBlockNumber *big.Int) (IndexMap, error) {

	indexMap := make(IndexMap)

	dequeuedProposals, err := GetDequeue(atBlockNumber)

	if err != nil {
		return indexMap, err
	}

	for idx, proposalID := range dequeuedProposals {
		indexMap[proposalID.String()] = big.NewInt(int64(idx))
	}

	return indexMap, nil

}

func calculateEpochs(propEpoch *big.Int, approvalStageDuration *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int) (*big.Int, *big.Int, *big.Int) {
	referendumEpoch := new(big.Int).Add(propEpoch, approvalStageDuration)
	executionEpoch := new(big.Int).Add(referendumEpoch, referendumStageDuration)
	expirationEpoch := new(big.Int).Add(executionEpoch, executionStageDuration)

	return referendumEpoch, executionEpoch, expirationEpoch

}
