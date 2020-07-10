package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/rpc"

	"math/big"
)

//GetBlockState returns the state dump for the given block number
func GetBlockState(c *rpc.Client, blockNumber *big.Int) (*state.Dump, error) {

	var result *state.Dump
	err := c.Call(&result, "debug_dumpBlock", fmt.Sprintf("0x%x", blockNumber))
	if err != nil {
		return nil, err
	}
	return result, nil
}

//GetAddresses extracts all addresses specified in the provided state dump and returns them as a slice
func GetAddresses(dump *state.Dump) []string {
	accts := make([]string, len(dump.Accounts))
	i := 0
	for k := range dump.Accounts {
		accts[i] = k.Hex()
		i++
	}
	return accts
}

//GetStateAddressesWithRelevantData extracts all addresses specified in a state dump and returns
// a mapping of these addresses to their corresponding gold token balances as reported by the state dump
func GetStateAddressesWithRelevantData(dump *state.Dump) map[string]RelevantStateData {
	ard := make(map[string]RelevantStateData)

	for address := range dump.Accounts {
		gtb, _ := new(big.Int).SetString(dump.Accounts[address].Balance, 10)

		ard[address.Hex()] = RelevantStateData{
			GoldTokenBalance: gtb}
		//AddColumnWork

	}

	return ard

}

//GetSystemGoldTokenSupplyFromStateDump returns the total gold token supply of the Celo blockchain
//at a particular block height as reported by the state dump
//This function is relevant to fetch total gold token supply
//when the GoldToken contract hasn't been activated on the blockchain.
// i.e. for block numbers : 0 (genesis) to T-1 (where T is the block height at which the GoldToken contract is initialised)
func GetSystemGoldTokenSupplyFromStateDump(c *rpc.Client, atBlockNumber *big.Int) (*big.Int, error) {
	dump, err := GetBlockState(c, atBlockNumber)

	if err != nil {
		return nil, err
	}

	sum := big.NewInt(0)

	for address := range dump.Accounts {

		gtb, _ := new(big.Int).SetString(dump.Accounts[address].Balance, 10)
		sum.Add(sum, gtb)

	}

	return sum, nil

}

//GetGoldTokenBalanceFromStateDump returns the gold token balanace of an address
//on the Celo blockchain at a particular block height as reported by the state dump.
//This function is relevant to fetch balances for
//when the GoldToken contract hasn't been activated on the blockchain.
// i.e. for block numbers : 0 (genesis) to T-1 (where T is the block height at which the GoldToken contract is initialised)
func GetGoldTokenBalanceFromStateDump(c *rpc.Client, accountAddressHex string, atBlockNumber *big.Int) (*big.Int, error) {

	dump, err := GetBlockState(c, atBlockNumber)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(accountAddressHex)

	val, addressExists := dump.Accounts[address]

	if !addressExists {
		return big.NewInt(0), nil
	}

	gtb, _ := new(big.Int).SetString(val.Balance, 10)
	return gtb, nil

}

//RelevantStateData is a wrapper around state dump data points that we intend
// to return for each address in GetStateAddressesWithRelevantData
//Currently, it only contains Gold Token Balance
type RelevantStateData struct {
	GoldTokenBalance *big.Int
	//AddColumnWork
}
