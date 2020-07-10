package rest

import (
	"database/sql"

	"github.com/ChorusOne/Icarus/contract"

	. "github.com/ChorusOne/Icarus/common"
	"github.com/ChorusOne/Icarus/rest/account"
	"github.com/ChorusOne/Icarus/rest/system"
	"github.com/ChorusOne/Icarus/rest/types"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

var ec *ethclient.Client
var rc *rpc.Client
var db *sql.DB

func initClients() {
	var err error

	ec, err = NewEthClient()

	if err != nil {
		log.Fatal(err)
	}

	rc, err = NewRPCClient()

	if err != nil {
		log.Fatal(err)
	}

}

func initDb(configFile string) {
	var err error
	db, err = NewPostgresDBFromConfig(configFile)

	if err != nil {
		log.Fatal(err)
	}
}

// StartAPI creates an HTTP server with a REST API for Anthem to consume. This
// function blocks and so should be spawned as a goroutine.
func StartAPI(configFile string) {

	initDb(configFile)

	initClients()
	contract.InitClientAndContracts()

	r := chi.NewRouter()
	r.Get("/", IndexResponder())

	r.Route("/accounts/{accountID}", func(r chi.Router) {

		r.Get("/", account.BalancesData(rc))
		r.Get("/history", account.SnapshotsData(db))
		r.Get("/transactions", account.TransactionsData(db, rc))

	})

	r.Route("/system", func(r chi.Router) {
		r.Get("/balances", system.BalancesData(rc))
		r.Get("/history", system.SnapshotsData(db))
		r.Get("/validator_groups", system.ValidatorGroupsData(rc))
		r.Get("/referendum_proposals", system.ReferendumProposalsData(rc))
		r.Get("/queued_proposals", system.QueuedProposalsData(rc))
		r.Get("/proposals_history", system.HistoricalProposalsData(rc))
		r.Get("/transactions/{txHash}", system.TransactionData(db))
	})

	http.ListenAndServe(":10101", r) // Todo : Config-ify
}

// IndexResponder acts as a health check.
func IndexResponder() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Repeat after me: I am not a tyrant, I am not a tyrant, I am not a tyrant..."))
	}
}

// -----------------------------------------------------------------------------
