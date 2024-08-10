package main

import (
	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

func main() {
	Log2()
}

func Log1() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
}

func Log2() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()
}
