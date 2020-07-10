package system

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/ChorusOne/Icarus/rest/types"

	"database/sql"
	"github.com/ChorusOne/Icarus/blockshot"
	"github.com/go-chi/chi"
)

// TransactionData returns details of a transaction - given a txhash
func TransactionData(Db *sql.DB) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		if txHash := chi.URLParam(r, "txHash"); txHash != "" {

			transaction := GetTransaction(Db, txHash)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(transaction); err != nil {
				log.Println(err)

			}
		}
	}

}

func GetTransaction(db *sql.DB, txHash string) *TransactionResponse {

	var tr TransactionResponse

	tx, err := blockshot.FetchTransactionByTxHash(db, txHash)

	if err != nil {
		tr.Message = "Transaction Not Found. Error"
		tr.Transaction = nil

		log.Printf("Trying to fetch Transaction %s caused an error. Error as follows. ", txHash)
		log.Print(err)

		return &tr

	}

	if tx == nil {
		tr.Message = "No transaction found for the provided transaction hash"
		tr.Transaction = nil
		return &tr
	}

	tr.Message = "Transaction Found"
	tr.Transaction = tx

	return &tr

}

type TransactionResponse struct {
	Message     string                         `json:"message"`
	Transaction *blockshot.TransactionDetailed `json:"transaction"`
}
