package atlas

import (
	. "github.com/ChorusOne/Icarus/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var atlasEthClient *ethclient.Client

func init() {

	client, err := NewEthClient()
	if err != nil {
		log.Fatal(err)
	}

	atlasEthClient = client
}
