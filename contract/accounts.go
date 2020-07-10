package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var accountsContractInstance *binding.Accounts

func InitAccountsContractInstance() {

	instance, err := binding.NewAccounts(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Accounts]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Accounts Contract is loaded")
	accountsContractInstance = instance

}

func GetName(accountAddressHex string, atBlockNumber *big.Int) (string, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Accounts])
	if !check {

		return "", nil

	}

	holder := common.HexToAddress(accountAddressHex)
	name, err := accountsContractInstance.GetName(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {

		log.Println(err)
	}
	return name, err

}
func GetMetadataURL(accountAddressHex string, atBlockNumber *big.Int) (string, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Accounts])
	if !check {

		return "", nil

	}

	holder := common.HexToAddress(accountAddressHex)
	metadataURL, err := accountsContractInstance.GetMetadataURL(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {

		log.Println(err)
	}
	return metadataURL, err

}
