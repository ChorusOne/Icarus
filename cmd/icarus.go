package main

import (
	"github.com/ChorusOne/Icarus/blockshot"
	. "github.com/ChorusOne/Icarus/common"

	"log"

	"github.com/joho/godotenv"
)

func main() {

	log.Print("Icarus - Time to Fly.")

	err := godotenv.Load("./config.env")
	if err != nil {
		log.Print("./config.env not found; continuing in the case you are using ENV vars.")
	}

	db, err := NewPostgresDBFromEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if Getenv("RESET_DB", "false") == "true" {
		blockshot.ResetDatabase(db)
	}

	blockshot.BlockIterator(db,
		[]blockshot.Task{
			blockshot.ProcessTransactions,
			//Note : GenerateTransactionTags depends on ProcessTransactions
			//So the following task should only appear after ProcessTransactions task
			//in this list
			blockshot.GenerateTransactionTags,
			blockshot.TakeDailySnapshots,
		})

}
