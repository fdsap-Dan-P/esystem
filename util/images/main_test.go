package images

import (
	"log"
	"os"
	"testing"

	"simplebank/util"

	_ "github.com/lib/pq"
)

var (
	Config util.Config
	err    error
)

func TestMain(m *testing.M) {
	Config, err = util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Printf("config.HomeFolder: %+v", Config.HomeFolder)

	os.Exit(m.Run())
}
