package logutil

import (
	"fmt"
	"os"
	"strings"

	datetimeconstant "github.com/nihcet/go-lib/pkg/constant/datetime"
	"github.com/rs/zerolog"
)

var log *Logger

type Logger struct {
	Info   func(msg string)                     // plain message
	Infof  func(pattern, msg string)            // message with pattern
	Error  func(msg string, err error)          // plain message and error
	Errorf func(pattern, msg string, err error) // message with pattern and error
}

func GetLogger() Logger {
	return *log
}

func InitializeLog(serviceName string) {
	consoleWriter := getConsoleWriter()
	zlog := zerolog.
		New(consoleWriter).
		With().Timestamp().
		Str("service_name", serviceName).
		Logger()

	log = &Logger{
		Info: func(msg string) {
			zlog.
				Info().
				Msg(msg)
		},
		Infof: func(pattern, msg string) {
			zlog.
				Info().
				Msgf(pattern, msg)
		},
		Error: func(msg string, err error) {
			zlog.
				Error().
				Err(err).
				Msg(msg)
		},
		Errorf: func(pattern, msg string, err error) {
			zlog.
				Error().
				Err(err).
				Msgf(pattern, msg)
		},
	}
}

// also set format output message
func getConsoleWriter() zerolog.ConsoleWriter {
	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: datetimeconstant.DateFormat_Datetime,
	}

	cw.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	cw.FormatErrFieldName = func(i any) string {
		return fmt.Sprintf("%s;", i)
	}

	return cw
}
