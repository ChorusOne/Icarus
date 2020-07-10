package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var electionContractInstance *binding.Election

func InitElectionContractInstance() {

	instance, err := binding.NewElection(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Election]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Election Contract is loaded")
	electionContractInstance = instance

}

func TotalVotesByAccount(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	holder := common.HexToAddress(accountAddressHex)
	tv, err := electionContractInstance.GetTotalVotesByAccount(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {
		log.Println(err)
	}
	return tv, err

}

type Group = string

type AccountVotesDistribution = map[Group]*VoteSplit

type VoteSplit struct {
	ActiveVotes  *big.Int
	PendingVotes *big.Int
	TotalVotes   *big.Int
}

func GetAccountVotesDistribution(accountAddressHex string, atBlockNumber *big.Int) (*AccountVotesDistribution, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		emptyMap := make(AccountVotesDistribution)
		return &emptyMap, nil
	}
	account := common.HexToAddress(accountAddressHex)
	groups, err := electionContractInstance.GetGroupsVotedForByAccount(&bind.CallOpts{BlockNumber: atBlockNumber}, account)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var avd AccountVotesDistribution = make(AccountVotesDistribution)

	for _, group := range groups {

		av, err1 := electionContractInstance.GetActiveVotesForGroupByAccount(&bind.CallOpts{BlockNumber: atBlockNumber}, group, account)

		if err1 != nil {
			log.Println(err1)
			return nil, err1
		}

		pv, err2 := electionContractInstance.GetPendingVotesForGroupByAccount(&bind.CallOpts{BlockNumber: atBlockNumber}, group, account)

		if err2 != nil {
			log.Println(err2)
			return nil, err2
		}

		//We can just do av+pv too as the method does that.
		//However, let's just rely on the method abstraction, so we don't have to change here if the method definition changes
		tv, err3 := electionContractInstance.GetTotalVotesForGroupByAccount(&bind.CallOpts{BlockNumber: atBlockNumber}, group, account)

		if err3 != nil {
			log.Println(err3)
			return nil, err3
		}

		avd[group.String()] = &VoteSplit{ActiveVotes: av, PendingVotes: pv, TotalVotes: tv}

	}

	return &avd, nil

}

func GetActiveVoteUnitsForGroup(groupAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	group := common.HexToAddress(groupAddressHex)

	tavu, err := electionContractInstance.GetActiveVoteUnitsForGroup(&bind.CallOpts{BlockNumber: atBlockNumber}, group)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tavu, nil

}

func GetActiveVoteUnitsForGroupByAccount(groupAddressHex string, accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	group := common.HexToAddress(groupAddressHex)
	account := common.HexToAddress(accountAddressHex)

	aavu, err := electionContractInstance.GetActiveVoteUnitsForGroupByAccount(
		&bind.CallOpts{BlockNumber: atBlockNumber},
		group,
		account)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return aavu, nil

}

func GetTotalVotes(atBlockNumber *big.Int) (*big.Int, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	tv, err := electionContractInstance.GetTotalVotes(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tv, nil

}

func getTotalVotesForEligibleValidatorGroups(atBlockNumber *big.Int) (map[common.Address]*big.Int, error) {

	var resultMap = make(map[common.Address]*big.Int)

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		return resultMap, nil
	}

	ret, err := electionContractInstance.GetTotalVotesForEligibleValidatorGroups(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
		return resultMap, err

	}

	groups := ret.Groups
	values := ret.Values

	for i := range groups {
		resultMap[groups[i]] = values[i]

	}
	return resultMap, nil

}

func GetNumVotesReceivable(groupAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Election])

	//Return zero value if polled before contract deployment
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	group := common.HexToAddress(groupAddressHex)

	nvr, err := electionContractInstance.GetNumVotesReceivable(&bind.CallOpts{BlockNumber: atBlockNumber}, group)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nvr, nil

}
