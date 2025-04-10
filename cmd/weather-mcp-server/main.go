package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/TuanKiri/weather-mcp-server/internal/server"
)

func main() {
	addr := flag.String("address", "", "The host and port to start the sse server")
	flag.Parse()

	cfg := &server.Config{
		ListenAddr:        *addr,
		WeatherAPIKey:     os.Getenv("WEATHER_API_KEY"),
		WeatherAPITimeout: 1 * time.Second,
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
