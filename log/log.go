package log

import (
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	Create           *zerolog.Logger
	Config           *LogConfig
	lumberjackWriter *lumberjack.Logger
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

func newRollingFile(config *LogConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path.Join(config.FileLoggingDirectory, config.FileLoggingFilename),
		MaxBackups: config.FileLoggingMaxBackups, // files
		MaxSize:    config.FileLoggingMaxSize,    // megabytes
		MaxAge:     config.FileLoggingMaxAge,     // days
	}
}