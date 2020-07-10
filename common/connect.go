package common

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// NewRPCClient returns a new rpcclient pointer for the Celo blockchain
func NewRPCClient() (*rpc.Client, error) {

	//First try IPC
	rc, err := rpc.Dial(IPCFilePath)

	if err != nil {
		//Then try RPC
		log.Println("Warning, not able to connect to IPC. Trying RPC but that would make snapshotting very slow")
		rc, err = rpc.Dial(RPCURLPath)
		if err != nil {
			log.Println("Have you set ipc/rpc paths to Celo full node?")
			return nil, err
		}
	}
	return rc, err
}

//NewEthClient returns a new ethclient pointer for the Celo blockchain
func NewEthClient() (*ethclient.Client, error) {

	//First try IPC
	ec, err := ethclient.Dial(IPCFilePath)

	if err == nil {

	} else {
		//Then try RPC
		log.Println("Warning, not able to connect to IPC. Trying RPC but that would make snapshotting very slow")

		ec, err = ethclient.Dial(RPCURLPath)
		if err != nil {
			log.Println("Have you set ipc/rpc paths to Celo full node?")
			return nil, err
		}

	}

	return ec, err

}

func NewPostgresDBFromConfig(config_file string) (*sql.DB, error) {

	err := godotenv.Load(config_file)
	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
