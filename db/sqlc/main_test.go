package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/annguyen34/simple-bank/util"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(testDB)

	os.Exit(m.Run())

}
