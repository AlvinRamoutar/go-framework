package log

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