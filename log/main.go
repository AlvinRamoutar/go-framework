package log

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var VERSION string = "0.0.1"

type LogConfig struct {
	LogLevel              int    `yaml:"loglevel"`
	ConsoleLoggingEnabled bool   `yaml:"consoleloggingenabled"`
	EncodeLogsAsJson      bool   `yaml:"encodelogasjson"`
	FileLoggingEnabled    bool   `yaml:"fileloggingenabled"`
	FileLoggingDirectory  string `yaml:"fileloggingdirectory"`
	FileLoggingFilename   string `yaml:"fileloggingfilename"`
	FileLoggingMaxSize    int    `yaml:"fileloggingmaxsize"`
	FileLoggingMaxBackups int    `yaml:"fileloggingmaxbackups"`
	FileLoggingMaxAge     int    `yaml:"fileloggingmaxage"`
}

func (l *LogConfig) New() *LogConfig {
	lc := LogConfig{}

	lc.LogLevel = 1
	lc.ConsoleLoggingEnabled = true
	lc.EncodeLogsAsJson = true
	lc.FileLoggingEnabled = false
	lc.FileLoggingDirectory = "./"
	lc.FileLoggingFilename = "log"
	lc.FileLoggingMaxSize = 10
	lc.FileLoggingMaxBackups = 99999999
	lc.FileLoggingMaxAge = 99999999

	return &lc
}

// panic (zerolog.PanicLevel, 5)
// fatal (zerolog.FatalLevel, 4)
// error (zerolog.ErrorLevel, 3)
// warn (zerolog.WarnLevel, 2)
// info (zerolog.InfoLevel, 1)
// debug (zerolog.DebugLevel, 0)
// trace (zerolog.TraceLevel, -1)
func (l *LogConfig) GetLogLevel() zerolog.Level {
	switch l.LogLevel {
	case 0:
		return zerolog.DebugLevel
	case 1:
		return zerolog.InfoLevel
	case 2:
		return zerolog.WarnLevel
	case 3:
		return zerolog.ErrorLevel
	case 4:
		return zerolog.FatalLevel
	case 5:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

type Log struct {
	Create           *zerolog.Logger
	Config           *LogConfig
	lumberjackWriter *lumberjack.Logger
}

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

func newRollingFile(config *LogConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path.Join(config.FileLoggingDirectory, config.FileLoggingFilename),
		MaxBackups: config.FileLoggingMaxBackups, // files
		MaxSize:    config.FileLoggingMaxSize,    // megabytes
		MaxAge:     config.FileLoggingMaxAge,     // days
	}
}
