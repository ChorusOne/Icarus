package system

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/ChorusOne/Icarus/blockchain"
	"github.com/ChorusOne/Icarus/contract"
	"github.com/ChorusOne/Icarus/rest/types"
	"github.com/ethereum/go-ethereum/rpc"

	"math/big"
)

// ValidatorGroupsData returns details of all validator groups at the given block height
func ValidatorGroupsData(rc *rpc.Client) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		latest, err := blockchain.GetLatestBlockNumber(rc)

		if err != nil {
			log.Print(err)

		} else {

			heightQuery := r.URL.Query().Get("block_height")
			height := big.NewInt(0)
			var success bool
			height, success = height.SetString(heightQuery, 10)

			//If height is invalid, return details for latest block number
			if !success ||
				height.Cmp(big.NewInt(0)) < 0 ||
				height.Cmp(latest) > 0 {

				height = latest

			}

			vgDetails := ValidatorGroupsWithDetails(height)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(vgDetails); err != nil {
				log.Println(err)

			}
		}
	}
}

func ValidatorGroupsWithDetails(blockNumber *big.Int) []*contract.ValidatorGroup {

	vgs, err := contract.GetRegisteredValidatorGroups(blockNumber)

	if err != nil {
		return []*contract.ValidatorGroup{}
	}

	vgDeets, err2 := contract.GetDetailsOfValidatorGroups(blockNumber, vgs)

	if err2 != nil {
		return []*contract.ValidatorGroup{}
	}

	return vgDeets

}
