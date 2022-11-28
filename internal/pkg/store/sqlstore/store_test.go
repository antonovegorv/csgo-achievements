package sqlstore_test

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=55000 user=postgres password=postgrespw dbname=postgres_test sslmode=disable"
	}

	os.Exit(m.Run())
}
