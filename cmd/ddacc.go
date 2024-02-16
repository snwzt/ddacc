package main

import (
	"os"

	"snwzt/ddacc/internal/manager"

	"github.com/rs/zerolog"
)

func main() {
	customlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}).With().Timestamp().Logger()

	manager.Execute(
		&customlogger,
		os.Exit,
		os.Args[1:],
	)
}
