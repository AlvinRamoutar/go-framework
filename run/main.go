package run

import (
	"encoding/json"
	"strconv"
)

var VERSION string = "0.0.1"

func (r *Run) Init(config *RunConfig) error {
	r.Config = config
	return nil
}

func (s *Run) Start() error {
	s.isAsync = false
	err := s.prestart()
	if err != nil {
		return err
	}
	s.Runuler.StartBlocking()
	return nil
}

func (s *Run) AsyncStart() error {
	s.isAsync = true
	err := s.prestart()
	if err != nil {
		return err
	}
	s.Runuler.StartAsync()
	return nil
}

func (s *Run) Restart() error {
	s.Stop()
	if s.isAsync {
		s.AsyncStart()
	} else {
		s.Start()
	}
	return nil
}

func (s *Run) Stop() error {
	s.Runuler.Stop()
	s.Runuler.Clear()
	return nil
}

func (s *Run) Status() (string, error) {
	status := map[string]string{
		"time":    s.Config.Time,
		"isAsync": strconv.FormatBool(s.isAsync),
		"running": strconv.FormatBool(s.Runuler.IsRunning()),
	}
	j, _ := json.Marshal(status)
	return string(j), nil
}

func (s *Run) Version() string {
	return VERSION
}
