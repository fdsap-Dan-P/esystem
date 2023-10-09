package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"simplebank/util"

	_ "github.com/denisenkom/go-mssqldb"
)

var testQueriesLocal *QueriesLocal
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open("mssql", config.DBeSystemLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueriesLocal = New(testDB)

	os.Exit(m.Run())
}

// func Copy2Container() {
// 	ctx := context.Background()

// 	nginxC, err := GenericContainer(ctx, GenericContainerRequest{
// 		ContainerRequest: ContainerRequest{
// 			Image:        "nginx:1.17.6",
// 			ExposedPorts: []string{"80/tcp"},
// 			WaitingFor:   wait.ForListeningPort("80/tcp"),
// 		},
// 		Started: true,
// 	})

// 	nginxC.CopyFileToContainer(ctx, "./testresources/hello.sh", "/hello_copy.sh", 700)
// }
