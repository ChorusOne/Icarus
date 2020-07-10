package blockchain

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type HeaderLite struct {
	Number    string `json:"number"`
	Timestamp string `json:"timestamp"`
}

//GetHeaderLite returns the timestamp of the specified blocknumber wrapped in a HeaderLite object
func GetHeaderLite(c *rpc.Client, blockNumber *big.Int) (*HeaderLite, error) {

	var result *HeaderLite

	err := c.Call(&result, "eth_getBlockByNumber", hexutil.EncodeBig(blockNumber), false)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetGenesisHeaderLite returns the timestamp of the genesis blocknumber wrapped in a HeaderLite object
func GetGenesisHeaderLite(c *rpc.Client) (*HeaderLite, error) {
	var result *HeaderLite

	err := c.Call(&result, "eth_getBlockByNumber", "earliest", false)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetLatestHeaderLite returns the timestamp of the latest blocknumber wrapped in a HeaderLite object
func GetLatestHeaderLite(c *rpc.Client) (*HeaderLite, error) {
	var result *HeaderLite

	err := c.Call(&result, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetLatestBlockNumber returns the latest block number encountered by the full node of the Celo blockchain
func GetLatestBlockNumber(c *rpc.Client) (*big.Int, error) {
	var number string

	err := c.Call(&number, "eth_blockNumber")
	if err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(number)

}
