//Package blockshot provides functions for iterating over the Celo blockchain block-by-block
//to perform a list of extraction tasks over each block.
package blockshot

import (
	"database/sql"
	"github.com/ChorusOne/Icarus/blockchain"
	. "github.com/ChorusOne/Icarus/common"
	"github.com/ChorusOne/Icarus/contract"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"time"
)

var noBlocksIterated *big.Int = big.NewInt(-1)
var one *big.Int = big.NewInt(1)

const sleepSeconds = 6

//Task is a type alias for extraction tasks
type Task = func(*sql.DB, *big.Int) error

var ethClient *ethclient.Client
var rpcClient *rpc.Client

//InitClients initializes clients for RPC/IPC calls to the Celo blockchain
func InitClients() {

	var err error

	ethClient, err = NewEthClient()
	if err != nil {
		log.Fatal(err)
	}

	rpcClient, err = NewRPCClient()

	if err != nil {
		log.Fatal(err)
	}
}

//BlockIterator iterates over the blocks and performs the supplied extraction tasks
func BlockIterator(db *sql.DB, tasks []Task) {

	InitClients()
	contract.InitClientAndContracts()

	chainHeight, err := blockchain.GetLatestBlockNumber(rpcClient)

	if err != nil {
		log.Fatal(err)
	}

	for {

		nextBlockNumber := big.NewInt(0)
		nextBlockNumber.Add(LatestIteratedBlockNumber(db), one)

		// Sleep if (and until) chainHeight < nextBlockNumber
		for chainHeight.Cmp(nextBlockNumber) < 0 {

			log.Print("Chain Height :", chainHeight.String())
			log.Println(" | Next Block Number : ", nextBlockNumber.String())
			log.Printf("Chain height exceeded. Sleeping for %d seconds \n", sleepSeconds)
			time.Sleep(sleepSeconds * time.Second)

			chainHeight, err = blockchain.GetLatestBlockNumber(rpcClient)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Iterate through all tasks for this block number
		for _, task := range tasks {
			err := task(db, nextBlockNumber)
			if err != nil {
				log.Fatal(err)
			}
		}

		SetIterationStatusComplete(db, nextBlockNumber)

	}
}
