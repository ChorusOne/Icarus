package common

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	_ "github.com/lib/pq"
	"log"
	"os"

)

// NewRPCClient returns a new rpcclient pointer for the Celo blockchain
func NewRPCClient() (*rpc.Client, error) {

	//First try IPC
	rc, err := rpc.Dial(Getenv("CELO_URI", IPCFilePath))

	if err != nil {
		log.Println("Warning, not able to connect to CELO_URI")
		return nil, err
	}
	return rc, err
}

//NewEthClient returns a new ethclient pointer for the Celo blockchain
func NewEthClient() (*ethclient.Client, error) {

	//First try IPC
	ec, err := ethclient.Dial(Getenv("CELO_URI", IPCFilePath))
	if err != nil {
		log.Println("Warning, not able to connect to CELO_URI")
		return nil, err
	}

	return ec, err

}

func NewPostgresDBFromEnvironment() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslMode := Getenv("DB_SSLMODE", "require")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " +
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
