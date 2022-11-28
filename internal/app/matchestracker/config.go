package matchestracker

// Config ...
type Config struct {
	DatabaseURL            string `toml:"database_url"`
	SteamAPIKey            string `toml:"steam_api_key"`
	SteamAPIEndpoint       string `toml:"steam_api_endpoint"`
	TickerTimeoutInSeconds int    `toml:"ticker_timeout_in_seconds"`
	LoggerLevel            string `toml:"logger_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
