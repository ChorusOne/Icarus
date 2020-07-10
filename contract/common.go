package contract

import (
	. "github.com/ChorusOne/Icarus/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

var clientInstance *ethclient.Client
var rpcClientInstance *rpc.Client

func initClient() {

	//First try to connect via IPC
	client, err := NewEthClient()
	if err != nil {
		log.Fatal(err)
	}

	clientInstance = client
}

func initRPCCLient() {
	rc, err := NewRPCClient()
	if err != nil {
		log.Fatal(err)
	}

	rpcClientInstance = rc
}

func InitClientAndContracts() {
	initClient()
	initRPCCLient()

	InitAccountsContractInstance()
	InitElectionContractInstance()
	InitGTContractInstance()
	InitLGContractInstance()
	InitStableTokenContractInstance()
	InitGovernanceContractInstance()
	InitValidatorsContractInstance()

}

func heightSanityCheck(blockNumber *big.Int, heightDeployed int64) bool {
	return (blockNumber.Cmp(big.NewInt(heightDeployed)) >= 0)

}
