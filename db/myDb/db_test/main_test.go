package myDb

import (
	"context"
	"log"
	"os"
	"simple_bank/db/myDb"

	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:Colegio5@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *myDb.Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}
	defer testDB.Close()

	testQueries = myDb.New(testDB)

	os.Exit(m.Run())
}
