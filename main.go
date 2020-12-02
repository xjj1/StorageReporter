package main

import (
	"log"

	"github.com/xjj1/StorageReporter/cmd"
	"github.com/xjj1/StorageReporter/db"
)

// GBsize could be 1000
const GBsize = 1024

func main() {
	repo, err := db.InitSQLiteRepo()
	if err != nil {
		log.Fatalln("init db:", err)
	}
	defer repo.Close()

	err = cmd.NewApp(repo).Execute()
	if err != nil {
		log.Println(err)
	}
}
