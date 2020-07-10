package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var lockedGoldContractInstance *binding.LockedGold

func InitLGContractInstance() {

	instance, err := binding.NewLockedGold(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][LockedGold]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Locked Gold Contract is loaded")
	lockedGoldContractInstance = instance

}

// TotalLockedGoldBalance returns the total amount of locked gold for an account at the given block height
func TotalLockedGoldBalance(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	holder := common.HexToAddress(accountAddressHex)
	tlg, err := lockedGoldContractInstance.GetAccountTotalLockedGold(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {

		log.Println(err)

	}
	return tlg, err

}

// NonVotingLockedGoldBalance returns the total amount of non-voting locked gold for an account at the given block height
func NonVotingLockedGoldBalance(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	holder := common.HexToAddress(accountAddressHex)
	nvlg, err := lockedGoldContractInstance.GetAccountNonvotingLockedGold(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {
		log.Println(err)
	}
	return nvlg, err
}

func PendingWithdrawals(accountAddressHex string, atBlockNumber *big.Int) ([]*big.Int, []*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		blankSlice := []*big.Int{}
		return blankSlice, blankSlice, nil
	}
	holder := common.HexToAddress(accountAddressHex)
	pw_values, pw_timestamps, err := lockedGoldContractInstance.GetPendingWithdrawals(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {
		log.Println(err)
	}
	return pw_values, pw_timestamps, err

}

// TotalPendingWithdrawalGoldBalance returns the total pending withdrawal gold for an account at the given block height
// Note : This function is not available in the LockedGold contract deployed on Alfajores
func TotalPendingWithdrawalGoldBalance(accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	if NetActive == TestnetAlfajores {
		log.Fatal("This function is not available on Alfajores Testnet")
	}

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	holder := common.HexToAddress(accountAddressHex)
	tpwg, err := lockedGoldContractInstance.GetTotalPendingWithdrawals(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {
		log.Println(err)
	}
	return tpwg, err

}

func SystemNonVotingLockedGoldBalance(atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	snvlg, err := lockedGoldContractInstance.GetNonvotingLockedGold(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return snvlg, err
}

func SystemTotalLockedGoldBalance(atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][LockedGold])

	if !check {
		zero := big.NewInt(0)
		return zero, nil
	}

	stlg, err := lockedGoldContractInstance.GetTotalLockedGold(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return stlg, err
}
