package log

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

var VERSION string = "0.0.1"

func (l *Log) Init(config *LogConfig) error {
	l.Config = config
	return nil
}

func (l *Log) Start() error {
	writers := []io.Writer{}

	if l.Config.ConsoleLoggingEnabled {
		writers = append(writers, os.Stdout)
	}
	if l.Config.FileLoggingEnabled {
		l.lumberjackWriter = newRollingFile(l.Config)
		writers = append(writers, l.lumberjackWriter)
	}
	mw := io.MultiWriter(writers...)

	zerolog.SetGlobalLevel(l.Config.GetLogLevel())
	if l.Config.LogLevel < 0 || l.Config.LogLevel > 5 {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger := zerolog.New(mw).With().Timestamp().Logger()
	if l.Config.LogLevel == 0 {
		logger = zerolog.New(mw).With().Caller().Timestamp().Logger()
	}

	l.Create = &logger
	return nil
}

func (l *Log) AsyncStart() error {
	return errors.New("ENLIBCOMM007")
}

func (l *Log) Restart() error {
	l.Stop()
	l.Start()
	return nil
}

func (l *Log) Reload(logConfig *LogConfig) error {
	l.Stop()
	l.Init(logConfig)
	l.Start()
	return nil
}

func (l *Log) Stop() error {
	if l.lumberjackWriter != nil {
		l.lumberjackWriter.Close()
	}
	return nil
}

func (l *Log) Status() (string, error) {
	status := map[string]string{
		"loglevel":              strconv.Itoa(l.Config.LogLevel),
		"consoleloggingenabled": strconv.FormatBool(l.Config.ConsoleLoggingEnabled),
		"encodelogsasjson":      strconv.FormatBool(l.Config.EncodeLogsAsJson),
		"fileloggingenabled":    strconv.FormatBool(l.Config.FileLoggingEnabled),
		"fileloggingdirectory":  l.Config.FileLoggingDirectory,
		"fileloggingfilename":   l.Config.FileLoggingFilename,
		"fileloggingmaxsize":    strconv.Itoa(l.Config.FileLoggingMaxSize),
		"fileloggingmaxbackups": strconv.Itoa(l.Config.FileLoggingMaxBackups),
		"fileloggingmaxage":     strconv.Itoa(l.Config.FileLoggingMaxAge),
	}
	j, _ := json.Marshal(status)
	return string(j), nil
}

func (l *Log) Version() string {
	return VERSION
}
