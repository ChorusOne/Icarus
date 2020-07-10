package account

import (
	"encoding/json"

	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/ChorusOne/Icarus/blockchain"

	"github.com/ChorusOne/Icarus/blockshot"
	"github.com/ChorusOne/Icarus/rest/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-chi/chi"
	"math/big"
)

// TransactionsData returns transactions data for a specific account
func TransactionsData(db *sql.DB, rc *rpc.Client) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		if accountID := chi.URLParam(r, "accountID"); accountID != "" {

			latest, err := blockchain.GetLatestBlockNumber(rc)

			if err != nil {
				log.Print(err)

			} else {

				heightQuery := r.URL.Query().Get("block_height")
				sourceFilterQuery := r.URL.Query().Get("filter_by_source")
				height := big.NewInt(0)
				var successH bool
				height, successH = height.SetString(heightQuery, 10)

				//If height is invalid, return details for latest block number
				if !successH ||
					height.Cmp(big.NewInt(0)) < 0 ||
					height.Cmp(latest) > 0 {

					height = latest

				}

				limit, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
				//If limit is missing or invalid, set limit to default = 50
				if err1 != nil || limit <= 0 {
					limit = 50
				}

				page, err2 := strconv.Atoi(r.URL.Query().Get("page"))
				//If page is missing or invalid, set page to default = 1
				if err2 != nil || page <= 0 {
					page = 1
				}

				transactionsInfo := GetTransactions(db, accountID, height, limit, page, sourceFilterQuery)
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(transactionsInfo); err != nil {
					log.Println(err)

				}
			}
		}
	}
}

func GetTransactions(db *sql.DB, address string, uptoBlockNumber *big.Int, limit int, page int, sourceFilterQuery string) []*blockshot.TransactionDetailed {

	offset := (page - 1) * limit

	txs, err := blockshot.FetchTransactions(db, address, uptoBlockNumber, limit, offset, sourceFilterQuery)

	if err != nil {
		log.Println(err)

	}

	return txs

}
