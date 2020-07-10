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

// QueuedProposalsData returns details of all queued proposals at the given block height
func QueuedProposalsData(rc *rpc.Client) types.Handler {
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

			queuedProposalDetails := QueuedProposalsWithDetails(height)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(queuedProposalDetails); err != nil {
				log.Println(err)
			}
		}
	}
}

func QueuedProposalsWithDetails(blockNumber *big.Int) []*contract.DetailedQueuedProposal {

	detailedQueuedProposals, err := contract.GetDetailedQueuedProposals(blockNumber)

	if err != nil {
		return []*contract.DetailedQueuedProposal{}
	}

	return detailedQueuedProposals

}
