package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"simplebank/util"

	_ "github.com/denisenkom/go-mssqldb"

	local "simplebank/db/datastore/esystemlocal"
)

var LocalDataStore *local.QueriesLocal
var DB *sql.DB

func main() {
	log.Println("Start...")
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DB, err = sql.Open("mssql", config.DBeSystemLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	LocalDataStore = local.New(DB)

	area, err := LocalDataStore.ListArea(context.Background())

	log.Printf("Area: %v", area)
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	for {
		go func() {
			log.Println("Start Loop...")
			for {
				select {
				case <-ticker.C:
					area, err = LocalDataStore.ListArea(context.Background())
					log.Printf("Area: %v", area)
					log.Println("Run...")
				case <-quit:
					log.Println("Stop...")
					ticker.Stop()
					return
				}
			}
		}()
		time.Sleep(5 * time.Second)
	}
	log.Println("End...")
	// os.Exit(m.Run())
}
