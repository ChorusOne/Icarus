package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	"github.com/ChorusOne/Icarus/blockchain"
	. "github.com/ChorusOne/Icarus/common"
)

var goldTokenContractInstance *binding.GoldToken

func InitGTContractInstance() {

	instance, err := binding.NewGoldToken(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][GoldToken]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Gold Token Contract is loaded")
	goldTokenContractInstance = instance

}

func GoldTokenBalance(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][GoldToken])

	//If polled before contract deployment, get value from State Dump
	if !check {
		g, err := blockchain.GetGoldTokenBalanceFromStateDump(rpcClientInstance, accountAddressHex, atBlockNumber)
		if err != nil {
			log.Println(err)
		}
		return g, err
	}

	holder := common.HexToAddress(accountAddressHex)
	g, err := goldTokenContractInstance.BalanceOf(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {

		log.Println(err)
	}
	return g, err

}

func SystemGoldTokenSupply(atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][GoldToken])

	//If polled before contract deployment, get value from State Dump
	if !check {

		sg, err := blockchain.GetSystemGoldTokenSupplyFromStateDump(rpcClientInstance, atBlockNumber)
		if err != nil {
			log.Println(err)
		}
		return sg, err
	}

	sg, err := goldTokenContractInstance.TotalSupply(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return sg, err
}
