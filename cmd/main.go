package main

import (
	"os"

	"snwzt/ddacc/internal/cli"

	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}).With().Timestamp().Logger()

	cli.Execute(
		&log,
		os.Exit,
		os.Args[1:],
	)
}
