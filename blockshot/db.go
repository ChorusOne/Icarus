package blockshot

import (
	"database/sql"
	"encoding/json"

	"log"
	"math/big"
	"strings"
)

//ResetDatabase deletes the block iterator tables if they exist
//and then creates fresh block iterator tables
// Table - block_iterator_status
// Table - block_iterator_transactions
func ResetDatabase(db *sql.DB) {

	DeleteTable("block_iterator_status", db)
	DeleteTable("block_iterator_transactions", db)
	DeleteTable("block_iterator_event_logs", db)
	DeleteTable("block_iterator_tagged_transactions", db)

	DeleteTable("snapshot_blocks", db)
	DeleteTable("snapshot_data", db)
	DeleteTable("system_snapshot_data", db)
	DeleteTable("votes_snapshot_data", db)
	DeleteTable("rewards_snapshot_data", db)

	CreateStatusTable(db)
	CreateTransactionsTable(db)
	CreateEventLogsTable(db)
	CreateTagsTable(db)

	CreateSnapshotBlocksTable(db)
	CreateSnapshotDataTable(db)
	CreateSystemSnapshotDataTable(db)
	CreateVotesSnapshotDataTable(db)
	CreateRewardsSnapshotDataTable(db)

}

//DeleteTable deletes a table from a database
func DeleteTable(tableName string, db *sql.DB) {
	DBExec("DROP TABLE IF EXISTS "+tableName, db)

}

//CreateStatusTable creates the block_iterator_status table
func CreateStatusTable(db *sql.DB) {

	statement := `
	CREATE TABLE block_iterator_status(
		block_number integer NOT NULL PRIMARY KEY,
		iteration_done boolean NOT NULL DEFAULT FALSE
	);
	`
	DBExec(statement, db)

}

//CreateTransactionsTable creates the block_iterator_transactions table
func CreateTransactionsTable(db *sql.DB) {

	statement := `
	CREATE TABLE block_iterator_transactions(
		block_number integer NOT NULL,
		timestamp text NOT NULL,
		tx_hash text NOT NULL PRIMARY KEY, 
		tx_details JSONB NOT NULL
	);
	`
	DBExec(statement, db)

}

func CreateTagsTable(db *sql.DB) {

	statement := `
	CREATE TABLE block_iterator_tagged_transactions(
		block_number integer NOT NULL,
		tx_hash text NOT NULL PRIMARY KEY, 
		from_address text NOT NULL, 
		to_address text, 
		events JSONB,
		tags JSONB
	);
	`
	DBExec(statement, db)

}

//CreateLogsTable creates the block_iterator_event_logs table
func CreateEventLogsTable(db *sql.DB) {

	statement := `
	CREATE TABLE block_iterator_event_logs(
		block_number integer NOT NULL,
		block_log_index integer NOT NULL, 

		tx_hash text NOT NULL, 
		tx_log_index integer NOT NULL, 

		topic text NOT NULL, 
		details JSONB NOT NULL, 
		PRIMARY KEY (block_number, block_log_index)
	);
	`

	DBExec(statement, db)

}

func CreateSnapshotBlocksTable(db *sql.DB) {
	statement := `
	CREATE TABLE snapshot_blocks(
		block_number integer NOT NULL PRIMARY KEY,
		block_timestamp text NOT NULL,
		snapshot_done boolean NOT NULL DEFAULT FALSE, 
		block_date text NOT NULL
	);
	`
	DBExec(statement, db)

}

func CreateSnapshotDataTable(db *sql.DB) {

	// AddColumnWork
	// Order Matters
	statement := `
	CREATE TABLE snapshot_data(
		block_number integer NOT NULL,
		address text NOT NULL,
		gold_token_balance text NOT NULL,
		locked_gold_balance text NOT NULL,
		nonvoting_locked_gold_balance text NOT NULL,
		pending_withdrawal_gold_balance text NOT NULL,
		celo_usd_value text NOT NULL,
		reward text NOT NULL,
		PRIMARY KEY (block_number, address)


		);
	`

	DBExec(statement, db)
}

func CreateSystemSnapshotDataTable(db *sql.DB) {
	// AddSystemColumnWork
	// Order Matters
	statement := `
	CREATE TABLE system_snapshot_data(

		block_number integer NOT NULL PRIMARY KEY,
		system_gold_token_supply text NOT NULL,
		system_total_locked_gold_balance text NOT NULL,
		system_nonvoting_locked_gold_balance text NOT NULL,
		system_celo_usd_value text NOT NULL 
	);
	`
	DBExec(statement, db)
}

func CreateVotesSnapshotDataTable(db *sql.DB) {

	// AddColumnWork
	// Order Matters
	statement := `
	CREATE TABLE votes_snapshot_data(
		block_number integer NOT NULL,
		address text NOT NULL,
		for_group text NOT NULL,
		active_votes text NOT NULL,
		pending_votes text NOT NULL,
		total_votes text NOT NULL,
		PRIMARY KEY (block_number, address, for_group)

		);
	`

	DBExec(statement, db)
}

func CreateRewardsSnapshotDataTable(db *sql.DB) {

	statement := `
	CREATE TABLE rewards_snapshot_data(
		block_number integer NOT NULL,
		epoch_block_number integer NOT NULL,
		address text NOT NULL, 
		for_group text NOT NULL,
		address_active_units text NOT NULL, 
		group_active_units text NOT NULL, 
		address_reward text NOT NULL, 
		group_reward text NOT NULL, 
		PRIMARY KEY (block_number, epoch_block_number, address, for_group)
		);
	`

	DBExec(statement, db)

}

//DBExec is a helper function that executes an SQL statement
//on a database connection
func DBExec(statement string, db *sql.DB) {
	_, err := db.Exec(statement)
	if err != nil {
		log.Panic(err)
	}

}

func LatestIteratedBlockNumber(db *sql.DB) *big.Int {

	var blockNumber int64
	stmt := `SELECT block_number FROM block_iterator_status
			WHERE iteration_done = TRUE 
			ORDER BY block_number DESC LIMIT 1`

	err := db.QueryRow(stmt).Scan(&blockNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			return noBlocksIterated
		} else {
			log.Fatal(err)
		}
	}

	return big.NewInt(blockNumber)

}

//SetIterationStatusComplete adds blocknumber to database and sets its value to be true
func SetIterationStatusComplete(db *sql.DB, blockNumber *big.Int) {

	stmt := "INSERT INTO block_iterator_status VALUES($1, $2)"
	_, err := db.Exec(stmt, blockNumber.String(), "TRUE")

	if err != nil {
		log.Fatal(err)
	}

}

//DumpTransaction inserts transaction data into block_iterator_transactions table
func DumpTransaction(db *sql.DB, blockNumber *big.Int, timestamp string, txHash string, transaction *Transaction) {

	stmt := "INSERT INTO block_iterator_transactions VALUES ($1, $2, $3, $4)"

	_, err := db.Exec(stmt, blockNumber.String(), timestamp, txHash, transaction)

	if err != nil {
		log.Fatal(err)
	}

}

//DumpEventLog inserts eventLog data into block_iterator_event_logs table
func DumpEventLog(db *sql.DB, blockNumber *big.Int, eLog *EventLog) {

	stmt := "INSERT INTO block_iterator_event_logs VALUES ($1, $2, $3, $4, $5, $6)"
	details, err := json.Marshal(eLog.Details)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		stmt,
		blockNumber.String(),
		eLog.BlockLogIndex,
		eLog.TxHash.Hex(),
		eLog.TxLogIndex,
		eLog.Topic,
		details)

	if err != nil {
		log.Fatal(err)
	}

}

func DumpTagsRow(db *sql.DB, blockNumber *big.Int, txHash string, from string, to string, eLogs EventLogs, tags Tags) {

	stmt := "INSERT INTO block_iterator_tagged_transactions VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := db.Exec(stmt,
		blockNumber.String(),
		txHash,
		from,
		to,
		eLogs,
		tags)

	if err != nil {
		log.Fatal(err)
	}

}

//GetHashOfTransactions returns all tx hashes in the given block (number)
func GetHashOfTransactions(db *sql.DB, blockNumber *big.Int) ([]string, error) {

	fetchStatement := "SELECT tx_hash from block_iterator_transactions where block_number = $1"

	rows, err := db.Query(fetchStatement, blockNumber.String())
	if err != nil {
		return nil, err
	}

	hashes := make([]string, 0)

	defer rows.Close()
	for rows.Next() {
		var hash string
		err := rows.Scan(&hash)

		if err != nil {
			return nil, err
		}

		hashes = append(hashes, hash)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hashes, nil
}

//FetchTransactions returns all transactions involving the given address
//upto the given blockheight (with tagging metadata)
func FetchTransactions(db *sql.DB, address string, uptoBlockNumber *big.Int, limit int, offset int, sourceContract string) ([]*TransactionDetailed, error) {

	/*
	   Let's do a left order join
	   That way, we will know which transactions are not tagged already
	   Because tagged column for them will be null
	   We need to handle them at controller level or at frontend level - maybe try and fetch individual tx after a delay
	*/

	selectStatement := `SELECT 
						block_iterator_transactions.block_number, 
						block_iterator_transactions.timestamp, 
						block_iterator_transactions.tx_hash, 
						block_iterator_transactions.tx_details,
						block_iterator_tagged_transactions.from_address,
						block_iterator_tagged_transactions.to_address,
						block_iterator_tagged_transactions.events,
						block_iterator_tagged_transactions.tags

						FROM block_iterator_transactions
						LEFT OUTER JOIN block_iterator_tagged_transactions 
						ON block_iterator_transactions.tx_hash = block_iterator_tagged_transactions.tx_hash
						

						WHERE block_iterator_transactions.block_number <= $1 
						AND  ( (block_iterator_tagged_transactions.from_address ILIKE $2) OR (block_iterator_tagged_transactions.to_address ILIKE $2) )`

	filterBySourceStatement := ""

	switch strings.ToLower(sourceContract) {

	case "governance":
		filterBySourceStatement = ` AND block_iterator_tagged_transactions.tags @> '[{"source":"Governance"}]'`

		//Add more cases if/when required by Anthem

	}

	orderStatement := ` ORDER BY block_iterator_tagged_transactions.block_number DESC LIMIT $3 OFFSET $4`

	fetchStatement := selectStatement + filterBySourceStatement + orderStatement

	rows, err := db.Query(fetchStatement, uptoBlockNumber.String(), address, limit, offset)

	if err != nil {
		return nil, err
	}

	transactions := make([]*TransactionDetailed, 0)

	defer rows.Close()
	for rows.Next() {

		var tdr TransactionDetailed

		err := rows.Scan(
			&tdr.BlockNumber,
			&tdr.Timestamp,
			&tdr.TxHash,
			&tdr.Details,
			&tdr.From,
			&tdr.To,
			&tdr.ELogs,
			&tdr.Tags)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, &tdr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil

}

//FetchTransactionByTxHash returns all transactions involving the given address
//upto the given blockheight (with tagging metadata)
func FetchTransactionByTxHash(db *sql.DB, txHash string) (*TransactionDetailed, error) {

	fetchStatement := `SELECT 
	block_iterator_transactions.block_number, 
	block_iterator_transactions.timestamp, 
	block_iterator_transactions.tx_hash, 
	block_iterator_transactions.tx_details,
	block_iterator_tagged_transactions.from_address,
	block_iterator_tagged_transactions.to_address,
	block_iterator_tagged_transactions.events,
	block_iterator_tagged_transactions.tags

	FROM block_iterator_transactions
	LEFT OUTER JOIN block_iterator_tagged_transactions 
	ON block_iterator_transactions.tx_hash = block_iterator_tagged_transactions.tx_hash

	WHERE block_iterator_transactions.tx_hash ILIKE $1`

	var tdr TransactionDetailed
	err := db.QueryRow(fetchStatement, txHash).Scan(
		&tdr.BlockNumber,
		&tdr.Timestamp,
		&tdr.TxHash,
		&tdr.Details,
		&tdr.From,
		&tdr.To,
		&tdr.ELogs,
		&tdr.Tags)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &tdr, nil

}
