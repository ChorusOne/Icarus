package account

import (
	"encoding/json"
	"github.com/ChorusOne/Icarus/rest/types"
	"github.com/ChorusOne/Icarus/snapshot"
	"github.com/go-chi/chi"
	"log"

	"database/sql"
	"net/http"
)

// SnapshotsData returns data for a specific account at all snapshot points
func SnapshotsData(Db *sql.DB) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		// Handle accountID = ""
		if accountID := chi.URLParam(r, "accountID"); accountID != "" {

			datedAccounts, _ := AccountWithHistoricalDetails(accountID, Db)

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(datedAccounts); err != nil {
				log.Println(err)

			}

		}
	}
}

func AccountWithHistoricalDetails(accountID string, Db *sql.DB) ([]*types.DatedAccount, error) {

	snapshotDates, err := snapshot.DatesOfCompletedSnapshots(Db)

	if err != nil {
		return nil, err
	}

	var datedAccounts = make([]*types.DatedAccount, 0, len(snapshotDates))

	// Fetching account rows from snapshot_data
	accountSnapshot, err2 := snapshot.FetchAccountSnapshot(Db, accountID)
	if err2 != nil {
		return nil, err2
	}

	accountVotesSnapshot, err3 := snapshot.FetchAccountVotesSnapshot(Db, accountID)
	if err3 != nil {
		return nil, err3
	}

	for _, snapshotDate := range snapshotDates {

		datedAccount := types.NewDatedAccount(snapshotDate.Date, accountID)
		datedAccount.Height = snapshotDate.BlockNumber

		if asr, ok := accountSnapshot[snapshotDate.BlockNumber]; ok {

			datedAccount.GoldTokenBalance = asr.GoldTokenBalance
			datedAccount.TotalLockedGoldBalance = asr.LockedGoldBalance
			datedAccount.NonVotingLockedGoldBalance = asr.NonVotingLockedGoldBalance
			datedAccount.PendingWithdrawalBalance = asr.PendingWithdrawalGoldBalance
			datedAccount.CeloUSDValue = asr.CeloUSDValue
			datedAccount.SnapshotReward = asr.Reward

		}

		// Fetch Votes Snapshot Data and Add To Accounts
		if avsr, ok := accountVotesSnapshot[snapshotDate.BlockNumber]; ok {

			dels := make([]*types.Delegation, 0, len(avsr))

			for _, dr := range avsr {
				drDelegation := types.Delegation(*dr)
				dels = append(dels, &drDelegation)
			}

			datedAccount.Delegations = dels

		}

		// Fetch Rewards Data and Add To Account
		datedAccounts = append(datedAccounts, datedAccount)
	}

	return datedAccounts, nil

}
