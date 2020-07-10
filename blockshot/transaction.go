package blockshot

import (
	"database/sql"
	"github.com/ChorusOne/Icarus/blockchain"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
)

//ProcessTransactions is a task to retrieve all transactions of the given block number,
//extract relevant information for each transaction and store it in the database
func ProcessTransactions(db *sql.DB, blockNumber *big.Int) error {

	b, err := blockchain.GetBlock(ethClient, blockNumber)

	if err != nil {
		return err
	}

	transactions := b.Transactions()
	timestamp := strconv.FormatUint(b.Time(), 10)

	for _, tx := range transactions {
		gasReceipt := blockchain.GetGasReceipt(ethClient, tx.Hash())
		processTransaction(db, blockNumber, timestamp, tx, gasReceipt)
	}

	return nil

}

//processTransaction parses and dumps to database - the information of a single transaction
func processTransaction(db *sql.DB, blockNumber *big.Int, timestamp string, tx *types.Transaction, gasReceipt *blockchain.GasReceipt) error {

	transaction := TransactionFrom(tx, gasReceipt)
	DumpTransaction(db, blockNumber, timestamp, tx.Hash().String(), transaction)

	return nil

}
