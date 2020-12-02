package main

import (
	"log"

	"github.com/xjj1/StorageReporter/db"
)

// VERSION contains program version
const VERSION = "1.3"

// GBsize could be 1000
const GBsize = 1024

func main() {
	DB, err := db.InitDB()
	if err != nil {
		log.Fatalln("init db:", err)
	}

	defer DB.Close()
}
