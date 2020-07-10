package system

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/ChorusOne/Icarus/rest/types"

	"database/sql"
	"github.com/ChorusOne/Icarus/snapshot"
)

// SnapshotsData returns system wide stats for all snapshots
func SnapshotsData(Db *sql.DB) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		snapshots := SystemSnapshots(Db)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(snapshots); err != nil {
			log.Println(err)

		}
	}
}

func SystemSnapshots(Db *sql.DB) []*snapshot.SystemSnapshotRow {

	allSnapshots, err := snapshot.FetchAllSystemSnapshots(Db)

	if err != nil {
		log.Print(err)
		return []*snapshot.SystemSnapshotRow{}
	}

	return allSnapshots

}
