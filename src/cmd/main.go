package main

import (
	"context"

	pkgconfig "github.com/Golerplate/pkg/config"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/user-gtw/internal/config"
)

func main() {
	ctx := context.Background()

	// parse configuration
	config := &config.Config{}
	if err := pkgconfig.ParseConfig(config); err != nil {
		log.Fatal().
			Err(err).
			Msg("main: unable to parse config")
	}

	_ = ctx
}
