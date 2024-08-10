package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

func main() {
	Log1()
	Log2()
	PrettyLoging()
	PrettyOutput()
	MultipleLogOutput()
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

func PrettyLoging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("foo", "bar").Msg("Hello world")
}

func PrettyOutput() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Str("foo", "bar").Msg("Hello World")

	// Output: 2006-01-02T15:04:05Z07:00 | INFO  | ***Hello World**** foo:BAR
}

func MultipleLogOutput() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Hello World!")
}
