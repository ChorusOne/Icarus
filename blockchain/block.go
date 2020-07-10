package blockchain

import (
	"context"
	// Note : Conventionally, we would be referencing celo-org/celo-blockchain
	// instead of ethereum/go-ethereum through the entire codebase when making
	// such imports.

	// But this is the "un-convention" that is followed throughout Celo's official codebase too.
	// However, even though we are referencing ethereum/go-ethereum, what is (and should be) 
	// actually referenced is celo-org/celo-blockchain packages - thanks to a workaround.
	// For more on the workaround : refer to this reply
	// https://github.com/celo-org/celo-blockchain/pull/1048#issuecomment-653174763
	

	// TODO : Currently, we manually enforce this workaround on Icarus. Let's automate this with a go.mod soon. 
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
)

//GetBlock returns the full block (including transactional data) for the given block number
func GetBlock(ec *ethclient.Client, blockNumber *big.Int) (*types.Block, error) {

	var result *types.Block
	result, err := ec.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}
