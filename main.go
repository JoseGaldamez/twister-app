package main

import (
	"log"
	"twister/app/db"
	"twister/app/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No connection to DB")
		return
	}

	handlers.Controllers()
}
