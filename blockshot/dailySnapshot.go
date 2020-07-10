package blockshot

import (
	"database/sql"
	"github.com/ChorusOne/Icarus/blockchain"
	"github.com/ChorusOne/Icarus/snapshot"
	"log"
	"math/big"
)

var genesisBlockNumber = big.NewInt(0)

//TakeDailySnapshots evaluates if the block number is first block of UTC day
//and then takes a snapshot if that is indeed the case.
func TakeDailySnapshots(db *sql.DB, blockNumber *big.Int) error {

	if isFirstBlockOfDay(blockNumber) {

		if isGenesisBlock(blockNumber) {
			snapshot.GenerateAndStoreGenesisSnapshot(db, blockNumber, rpcClient)
		} else {
			snapshot.GenerateAndStoreBlockSnapshots(db, blockNumber, rpcClient)

		}
	}
	return nil
}

func isGenesisBlock(blockNumber *big.Int) bool {
	return blockNumber.Cmp(genesisBlockNumber) == 0
}

func isFirstBlockOfDay(blockNumber *big.Int) bool {

	if isGenesisBlock(blockNumber) {
		return true
	}

	var previousBlockNumber big.Int
	previousBlockNumber.Sub(blockNumber, big.NewInt(1))

	previousBlockHeader, err1 := blockchain.GetHeaderLite(rpcClient, &previousBlockNumber)

	if err1 != nil {
		log.Fatal(err1)
	}

	currentBlockHeader, err2 := blockchain.GetHeaderLite(rpcClient, blockNumber)

	if err2 != nil {
		log.Fatal(err2)
	}

	return !snapshot.SameDateOfTimestamps(previousBlockHeader.Timestamp, currentBlockHeader.Timestamp)

}
