package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/CSUOS/rabums/pkg/config"
	"github.com/CSUOS/rabums/pkg/database"
	"github.com/CSUOS/rabums/pkg/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate oapi-codegen --package=server --generate chi-server,types,spec -o pkg/server/server.gen.go api/swagger.yaml

func main() {
	log.Info().Msg("Hello")

	ctx := context.Background()

	// Set Logging Format
	if config.Mode == config.ModeProd {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123})
	}
	log.Logger = log.Logger.With().Caller().Timestamp().Logger()
	ctx = log.Logger.WithContext(ctx)

	config.Init(ctx)

	// Initializing pkgs
	database.Init(ctx)
	defer database.Close(ctx)
	r := server.Init(ctx)

	// Listening
	log.Info().Msg("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
