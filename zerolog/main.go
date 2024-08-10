package main

import (
	"log"

	"github.com/rs/zerolog"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
}
