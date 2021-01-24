package main

import (
	"github.com/caarlos0/env"
	"github.com/jekabolt/quotes/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	c := &server.Config{}
	err := env.Parse(c)
	if err != nil {
	}
	setLogLevel(c)
	log.Info().Str("config", c.String()).Send()

	s, err := c.InitServer()
	if err != nil {
		log.Fatal().Err(err).Msg("InitServer failed")
	}
	err = s.Serve()
	if err != nil {
		log.Fatal().Err(err).Msg("Run failed")
	}
}

func setLogLevel(cfg *server.Config) {
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("debug is enabled")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
