package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/sankester/simple_bank/utils"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect db : ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
