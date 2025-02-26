package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nytru/easy-news/src/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Error creating logging file")
	}
	defer f.Close()
	log.Logger = zerolog.New(f).With().Timestamp().Logger()

	err = godotenv.Load()
	if err != nil {
		log.
			Warn().
			Err(err).
			Msg("Fail loading env")
	}

	err = cmd.NewApp().Run()
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("App exited with error")
	}
}
