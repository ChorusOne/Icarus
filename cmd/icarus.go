package main

import (
	"github.com/ChorusOne/Icarus/blockshot"
	. "github.com/ChorusOne/Icarus/common"

	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("icarus_%d.log", timestamp)

	file, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file for icarus")
	}

	log.SetOutput(file)

	log.Print("Icarus - Time to Fly.")

	db, err := NewPostgresDBFromConfig("./config.env")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//The following statement flushes Icarus' database.
	//Comment out if you wish to resume the iterator without flushing the database
	//TODO : Configify this
	blockshot.ResetDatabase(db)

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
