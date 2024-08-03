package myDb_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"simple_bank/db/myDb"
	"simple_bank/util"

	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *myDb.Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println(config.DBSource)
	fmt.Println(config.ServerAddress)
	fmt.Println(config.AccessTokenDuration)

	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}
	defer testDB.Close()

	testQueries = myDb.New(testDB)

	os.Exit(m.Run())
}
