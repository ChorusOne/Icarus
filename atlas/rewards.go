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

// Atlas Functions relevant to rewards

//RewardInfo stores information representing data for the event EpochRewardsDistributedToVoters(address,uint256)
// plus a superset list of addresses that are potential recipient of shares of this reward
type RewardInfo struct {
	Group        string
	GroupHash    common.Hash
	RewardValue  *big.Int
	BlockNumber  *big.Int
	AddressAtlas map[common.Address]bool
}

//EpochRewardEvent is a struct to represent data for the event EpochRewardsDistributedToVoters(address,uint256)"
type EpochRewardEvent struct {
	Group common.Address
	Value *big.Int
}

//AddAtlasToRewardInfo : For a validator group, a reward and the block height at which the block was received,
// compute a list of all addresses that have (ever) activated a reward for the group upto this block height
// and hence are potential recipients of a share of these rewards.
// This will return a superset of the list of addresses that are actual recipients of the reward shares.
// (depending on whether their vote for the group was active in this epoch cycle )
// Relevant event 
	/* event ValidatorGroupVoteActivated(
	   address indexed account,
	   address indexed group,
	   uint256 value,
	   uint256 units
	*/
func AddAtlasToRewardInfo(ri *RewardInfo) {

	var logVGVA = []byte("ValidatorGroupVoteActivated(address,address,uint256,uint256)")
	var logVGVAHash = crypto.Keccak256Hash(logVGVA)
	//Filter by Event = ValidatorGroupVoteActivated and Group = ri.GroupHash
	var TopicsFilter = [][]common.Hash{{logVGVAHash}, {}, {ri.GroupHash}}

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Election])

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   ri.BlockNumber,
		Topics:    TopicsFilter,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	ri.AddressAtlas = make(map[common.Address]bool)
	for _, vLog := range logs {

		if !(vLog.Topics[2] == ri.GroupHash) {
			log.Panic("TopicFilter didn't work correctly")
		}
		address := common.HexToAddress(vLog.Topics[1].Hex())
		ri.AddressAtlas[address] = true

	}

}

// GetRewardEventsInfo returns list of all validator groups
// (with corresponding reward value and block number at which the reward was received)
// that received epoch rewards between (and including) the two specified block numbers.
func GetRewardEventsInfo(fromBlock *big.Int, toBlock *big.Int) []*RewardInfo {

	var logEpochRewardSig = []byte("EpochRewardsDistributedToVoters(address,uint256)")
	var logEpochRewardSigHash = crypto.Keccak256Hash(logEpochRewardSig)
	var TopicsFilter = [][]common.Hash{{logEpochRewardSigHash}}

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Election])

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

	rewards_info := make([]*RewardInfo, 0, len(logs))

	contractAbi, err := abi.JSON(strings.NewReader(string(binding.ElectionABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		var epochRewardEvent EpochRewardEvent
		err := contractAbi.Unpack(&epochRewardEvent, "EpochRewardsDistributedToVoters", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}
		ri := &RewardInfo{Group: common.HexToAddress(vLog.Topics[1].Hex()).String(),
			GroupHash:   vLog.Topics[1],
			RewardValue: epochRewardEvent.Value,
			BlockNumber: new(big.Int).SetUint64(vLog.BlockNumber)}

		AddAtlasToRewardInfo(ri)

		rewards_info = append(rewards_info, ri)
	}

	return rewards_info
}
