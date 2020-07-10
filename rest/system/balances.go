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

// BalancesData requests system wide stats at the latest block height
func BalancesData(rc *rpc.Client) types.Handler {
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

			system := SystemWithDetails(height)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(system); err != nil {
				log.Println(err)

			}
		}
	}
}

func SystemWithDetails(blockNumber *big.Int) *types.System {

	system := types.NewSystem(blockNumber.String())

	gt, _ := contract.SystemGoldTokenSupply(blockNumber)
	tlg, _ := contract.SystemTotalLockedGoldBalance(blockNumber)
	nvlg, _ := contract.SystemNonVotingLockedGoldBalance(blockNumber)
	cusd, _ := contract.SystemStableTokenSupplyValue(blockNumber)

	system.GoldTokenSupply = gt.String()
	system.TotalLockedGoldBalance = tlg.String()
	system.NonVotingLockedGoldBalance = nvlg.String()
	system.TotalCeloUSDValue = cusd.String()

	return system

}
