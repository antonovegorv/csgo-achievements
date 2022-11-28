package matchestracker

import (
	"database/sql"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/store/sqlstore"
	_ "github.com/lib/pq"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	client := newClient(
		store,
		config.SteamAPIKey,
		config.SteamAPIEndpoint,
		config.TickerTimeoutInSeconds,
		config.LoggerLevel,
	)

	return client.trackMatches()
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
