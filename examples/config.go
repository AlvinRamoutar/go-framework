package main

import (
	"errors"
	"os"

	"alvinr.ca/go-framework/http"
	"alvinr.ca/go-framework/log"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Config *AppSettings `yaml:"config"`
}

type AppSettings struct {
	Language string           `yaml:"language"`
	Logging  *log.LogConfig   `yaml:"logging"`
	Http     *http.HttpConfig `yaml:"http"`
}

func (c *Config) New() *Config {
	c.Config = &AppSettings{}

	c.Config.Language = "EN"

	c.Config.Logging = &log.LogConfig{}
	c.Config.Logging = c.Config.Logging.New()
	c.Config.Http = &http.HttpConfig{}
	c.Config.Http = c.Config.Http.New()

	return c
}

func (c *Config) Init(logConfig *log.LogConfig) {
	c = c.New()
	c.Config.Logging = logConfig
}

func (c *Config) Start() error {
	var cf []byte
	var err error

	if len(os.Args) > 1 {
		cf, err = os.ReadFile(os.Args[1])
		err = yaml.Unmarshal(cf, &c)
	}
	err = envconfig.Process(APP_NAME, c.Config)
	return err
}

func (c *Config) AsyncStart() error {
	return errors.New("ENLIBCOMM007")
}

func (c *Config) Restart() error {
	lc := &c.Config.Logging
	c.Stop()
	c.Init(*lc)
	return c.Start()
}

func (c *Config) Stop() error {
	c.Config = nil
	return nil
}

func (c *Config) Status() (string, error) {
	return "", nil
}

func (c *Config) Version() string {
	return VERSION
}
