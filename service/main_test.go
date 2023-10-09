package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testLaptop *LaptopServer
var testDB *sql.DB

// var test

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	testDB, err = sql.Open(config.DBDriver, postgresqlDbInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testLaptop = NewLaptopServer(testDB)

	// testStore = NewStoreAccount(testDB)

	os.Exit(m.Run())
}
