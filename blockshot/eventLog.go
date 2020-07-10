package blockshot

import (
	"context"

	"database/sql"
	"github.com/ethereum/go-ethereum"
	"math/big"
)

//ProcessEventLogs retrieves all event logs generated in the block corresponding to
//the specified block number and then stores them in the database
func ProcessEventLogs(db *sql.DB, blockNumber *big.Int) error {

	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
	}

	logs, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	for _, vLog := range logs {
		eLog := EventLogFrom(&vLog)
		DumpEventLog(db, blockNumber, eLog)
	}

	return nil
}
