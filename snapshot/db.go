package snapshot

import (
	"database/sql"
	"log"
	"math/big"
)

func DumpVotesSnapshotData(db *sql.DB, vs VotesSnapshot, blockNumber *big.Int) error {

	insertStatement := "INSERT INTO votes_snapshot_data VALUES ($1, $2, $3, $4, $5, $6)"

	for address := range vs {

		accountVoteDistribution := vs[address]

		for group := range *accountVoteDistribution {

			voteSplit := (*accountVoteDistribution)[group]

			_, err := db.Exec(insertStatement,
				blockNumber.String(),
				address,
				group,
				voteSplit.ActiveVotes.String(),
				voteSplit.PendingVotes.String(),
				voteSplit.TotalVotes.String()) //AddVotesColumnWork

			if err != nil {
				return err
			}

		}

	}

	return nil

}

func DumpSystemSnapshotData(db *sql.DB, ss SystemSnapshot, blockNumber *big.Int) error {

	insertStatement := "INSERT INTO system_snapshot_data VALUES ($1, $2, $3, $4, $5)" //AddSystemColumnWork

	_, err := db.Exec(insertStatement,
		blockNumber.String(),
		ss.GoldTokenSupply.String(),
		ss.TotalLockedGoldBalance.String(),
		ss.NonVotingLockedGoldBalance.String(),
		ss.TotalCeloUSDValue.String())
	//AddSystemColumnWork
	//Order Matters

	if err != nil {
		return err
	}

	return nil

}

func FetchAllSystemSnapshots(db *sql.DB) ([]*SystemSnapshotRow, error) {

	allSnapshots := make([]*SystemSnapshotRow, 0)

	fetchStatement := `SELECT 
	system_snapshot_data.block_number, 
	snapshot_blocks.block_date, 
	system_snapshot_data.system_gold_token_supply, 
	system_snapshot_data.system_total_locked_gold_balance, 
	system_snapshot_data.system_nonvoting_locked_gold_balance, 
	system_snapshot_data.system_celo_usd_value 


	FROM system_snapshot_data
	INNER JOIN snapshot_blocks
	ON system_snapshot_data.block_number = snapshot_blocks.block_number
	ORDER BY system_snapshot_data.block_number DESC`

	rows, err := db.Query(fetchStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ssr SystemSnapshotRow
		err := rows.Scan(
			&ssr.BlockNumber,
			&ssr.SnapshotDate,
			&ssr.GoldTokenSupply,
			&ssr.TotalLockedGoldBalance,
			&ssr.NonVotingLockedGoldBalance,
			&ssr.TotalCeloUSDValue)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		allSnapshots = append(allSnapshots, &ssr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allSnapshots, nil

}

func DumpSnapshotData(db *sql.DB, snapshot Snapshot, blockNumber *big.Int) error {
	//AddColumnWork
	insertStatement := "INSERT INTO snapshot_data VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)" //AddColumnWork

	for address := range snapshot {

		snapshotRow := snapshot[address]
		_, err := db.Exec(insertStatement,
			blockNumber.String(),
			address,
			snapshotRow.GoldTokenBalance.String(),
			snapshotRow.LockedGoldBalance.String(),
			snapshotRow.NonVotingLockedGoldBalance.String(),
			snapshotRow.PendingWithdrawalGoldBalance.String(),
			snapshotRow.CeloUSDValue.String(),
			snapshotRow.Reward.String(),
			snapshotRow.Commission.String())
		//AddColumnWork //Order Matters
		if err != nil {
			return err
		}

	}

	return nil

}

func FetchSnapshotAddresses(db *sql.DB, blockNumber *big.Int) ([]string, error) {

	accts := make([]string, 0)
	fetchStatement := "SELECT address from snapshot_data where block_number = $1"

	rows, err := db.Query(fetchStatement, blockNumber.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var address string
		err := rows.Scan(&address)
		if err != nil {
			return nil, err
		}
		accts = append(accts, address)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return accts, nil

}

func FetchSnapshotData(db *sql.DB, blockNumber *big.Int) (Snapshot, error) {

	var snapshot Snapshot
	snapshot = make(map[string]SnapshotRow)
	//AddColumnWork // Order Matters
	fetchStatement := "SELECT address, gold_token_balance, locked_gold_balance, nonvoting_locked_gold_balance , pending_withdrawal_gold_balance, celo_usd_value, reward, commission_celo_usd from snapshot_data where block_number = $1"

	rows, err := db.Query(fetchStatement, blockNumber.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var spr SnapshotPostgresRow
		err := rows.Scan(
			&spr.Address,
			&spr.GoldTokenBalance,
			&spr.LockedGoldBalance,
			&spr.NonVotingLockedGoldBalance,
			&spr.PendingWithdrawalGoldBalance,
			&spr.CeloUSDValue,
			&spr.Reward,
			&spr.Commission)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		AddRowToSnapshot(&snapshot, &spr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snapshot, nil

}

func CreateSnapshotStatusRow(db *sql.DB, blockNumber *big.Int, blockTimestamp string) {

	_, err := db.Exec("INSERT INTO snapshot_blocks VALUES($1, $2, $3, $4)",
		blockNumber.String(),
		blockTimestamp,
		true,
		UTCTime(blockTimestamp).Format(("02-01-2006")),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func fetchPreviousSnapshotBlockNumber(db *sql.DB) *big.Int {

	var block_number int64
	stmt := `SELECT block_number FROM snapshot_blocks 
			WHERE snapshot_done = TRUE 
			ORDER BY block_number DESC LIMIT 1`

	err := db.QueryRow(stmt).Scan(&block_number)

	if err != nil {
		log.Fatal(err)
	}

	return big.NewInt(block_number)
}

func FlushAllSnapshots(db *sql.DB) {

	_, err := db.Exec("TRUNCATE TABLE snapshot_data")
	if err != nil {
		log.Fatal(err)
	}

}
func AddSnapshotColumn(db *sql.DB, columnName string) {

	stmt := "ALTER TABLE snapshot_data ADD " + columnName + " text NOT NULL"
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

}

func DumpRewardsSnapshotData(db *sql.DB, rs RewardsSnapshot, blockNumber *big.Int) error {
	//AddColumnWork
	insertStatement := "INSERT INTO rewards_snapshot_data VALUES($1, $2, $3, $4, $5, $6, $7, $8)" //AddColumnWork

	for _, row := range rs {

		_, err := db.Exec(insertStatement,
			blockNumber.String(),
			row.EpochBlockNumber.String(),
			row.Address,
			row.Group,
			row.AddressActiveUnits.String(),
			row.GroupActiveUnits.String(),
			row.AddressReward.String(),
			row.GroupReward.String())

		if err != nil {
			return err
		}

	}

	return nil

}

func DumpCommissionsSnapshotData(db *sql.DB, cs CommissionsSnapshot, blockNumber *big.Int) error {
	//AddColumnWork
	insertStatement := "INSERT INTO commissions_snapshot_data VALUES($1, $2, $3, $4, $5, $6)" //AddColumnWork

	for _, row := range cs {

		_, err := db.Exec(insertStatement,
			blockNumber.String(),
			row.EpochBlockNumber.String(),
			row.Validator,
			row.ValidatorPayment.String(),
			row.Group,
			row.GroupPayment.String())

		if err != nil {
			return err
		}

	}

	return nil

}

func FetchAccountVotesSnapshot(db *sql.DB, account string) (AccountVotesSnapshot, error) {

	fetchStatement := `SELECT block_number, for_group, active_votes, pending_votes, total_votes
						FROM votes_snapshot_data
						WHERE address ILIKE $1`

	rows, err := db.Query(fetchStatement, account)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var avs = make(AccountVotesSnapshot)

	for rows.Next() {

		var dr DelegationRow
		var blockNumber string
		err := rows.Scan(
			&blockNumber,
			&dr.Group,
			&dr.ActiveVotes,
			&dr.PendingVotes,
			&dr.TotalVotes)
		if err != nil {
			return nil, err
		}

		if _, ok := avs[blockNumber]; !ok {
			avs[blockNumber] = make([]*DelegationRow, 0, 5)

		}

		avs[blockNumber] = append(avs[blockNumber], &dr)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return avs, nil

}

func FetchAccountSnapshot(db *sql.DB, account string) (AccountSnapshot, error) {

	var accountSnapshot = make(AccountSnapshot)
	//AddColumnWork
	//Order Matters
	fetchStatement := `SELECT block_number, gold_token_balance, locked_gold_balance, 
						nonvoting_locked_gold_balance , pending_withdrawal_gold_balance, 
						celo_usd_value, reward, commission_celo_usd
						FROM snapshot_data 
						WHERE address ILIKE $1`

	rows, err := db.Query(fetchStatement, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var asr AccountSnapshotRow
		var blockNumber string
		err := rows.Scan(
			&blockNumber,
			&asr.GoldTokenBalance,
			&asr.LockedGoldBalance,
			&asr.NonVotingLockedGoldBalance,
			&asr.PendingWithdrawalGoldBalance,
			&asr.CeloUSDValue,
			&asr.Reward,
			&asr.Commission)
		//AddColumnWork
		//Order Matters
		if err != nil {
			return nil, err
		}
		accountSnapshot[blockNumber] = &asr
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return accountSnapshot, nil

}

func DatesOfCompletedSnapshots(db *sql.DB) ([]*SnapshotDate, error) {

	fetchStatement := `SELECT block_number, block_date 
					 	FROM snapshot_blocks 
					 	WHERE snapshot_done = TRUE
						ORDER BY block_number ASC`

	rows, err := db.Query(fetchStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dates := make([]*SnapshotDate, 0, 100)

	for rows.Next() {
		var snapshotdate SnapshotDate
		err := rows.Scan(&snapshotdate.BlockNumber, &snapshotdate.Date)

		if err != nil {
			return nil, err
		}

		dates = append(dates, &snapshotdate)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return dates, nil

}
