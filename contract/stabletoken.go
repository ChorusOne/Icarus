package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var stableTokenContractInstance *binding.StableToken

func InitStableTokenContractInstance() {

	instance, err := binding.NewStableToken(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][StableToken]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("StableToken Contract is loaded")
	stableTokenContractInstance = instance

}

func StableTokenBalanceValue(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][StableToken])
	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	holder := common.HexToAddress(accountAddressHex)
	st, err := stableTokenContractInstance.BalanceOf(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {

		log.Println(err)
	}
	return st, err

}

func SystemStableTokenSupplyValue(atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][StableToken])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}
	sst, err := stableTokenContractInstance.TotalSupply(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return sst, err
}
