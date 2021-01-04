package main

import (
	"github.com/ChorusOne/Icarus/rest"
	"github.com/joho/godotenv"
	"log"
)

func main() {

        err := godotenv.Load("./config.env")
        if err != nil {
                log.Print("./config.env not found; continuing in the case you are using ENV vars.")
        }

	rest.StartAPI()

}
