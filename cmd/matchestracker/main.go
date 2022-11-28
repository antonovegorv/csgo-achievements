package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/antonovegorv/csgo-achievements/internal/app/matchestracker"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./configs/matchestracker.toml", "config path")
}

func main() {
	flag.Parse()

	config := matchestracker.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := matchestracker.Start(config); err != nil {
		log.Fatal(err)
	}
}
