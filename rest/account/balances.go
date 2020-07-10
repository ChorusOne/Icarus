// Provide Account related endpoints and fetch functions
package account

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/ChorusOne/Icarus/blockchain"
	"github.com/ChorusOne/Icarus/contract"
	"github.com/ChorusOne/Icarus/rest/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-chi/chi"
	"math/big"
)

// BalancesData requests data for a specific account
func BalancesData(rc *rpc.Client) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		if accountID := chi.URLParam(r, "accountID"); accountID != "" {

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

				accountInfo := AccountWithDetails(accountID, height)
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(accountInfo); err != nil {
					log.Println(err)

				}
			}
		}
	}
}

// AccountWithDetails prepares Account object for a given address with its details at a given block height
func AccountWithDetails(address string, blockNumber *big.Int) *types.Account {

	account := types.NewAccount(address)
	account.Height = blockNumber.String()

	gtb, _ := contract.GoldTokenBalance(address, blockNumber)
	pwb, _ := contract.TotalPendingWithdrawalGoldBalance(address, blockNumber)
	cusd, _ := contract.StableTokenBalanceValue(address, blockNumber)

	account.GoldTokenBalance = gtb.String()
	account.PendingWithdrawalBalance = pwb.String()
	account.CeloUSDValue = cusd.String()

	tlg, _ := contract.TotalLockedGoldBalance(address, blockNumber)
	nvlg, _ := contract.NonVotingLockedGoldBalance(address, blockNumber)
	vlg := big.NewInt(0)
	vlg.Sub(tlg, nvlg)

	account.TotalLockedGoldBalance = tlg.String()
	account.NonVotingLockedGoldBalance = nvlg.String()
	account.VotingLockedGoldBalance = vlg.String()

	avd, _ := contract.GetAccountVotesDistribution(address, blockNumber)

	for group := range *avd {

		voteSplit := (*avd)[group]

		delegation := &types.Delegation{
			Group:        group,
			TotalVotes:   voteSplit.TotalVotes.String(),
			ActiveVotes:  voteSplit.ActiveVotes.String(),
			PendingVotes: voteSplit.PendingVotes.String(),
		}

		account.Delegations = append(account.Delegations, delegation)
	}

	return account

}
