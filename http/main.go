package http

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

var VERSION string = "0.0.1"

func (h *Http) Init(config *HttpConfig) error {
	h.Config = config
	h.DefaultRouteRegexp, _ = regexp.Compile("^[1-9][0-9]{2}$")

	// config validation
	if !IsValidFQDN(h.Config.Host) {
		return errors.New("ENLIBHTTP007")
	}
	if !IsValidPortNumber(h.Config.Port) {
		return errors.New("ENLIBHTTP008")
	}

	return nil
}

func (h *Http) Start() error {
	return errors.New("ENLIBCOMM008")
}

func (h *Http) AsyncStart() error {
	return h.Serve()
}

func (h *Http) Restart() error {
	err := h.Closer.Close()
	if err != nil {
		return err
	}
	return h.Serve()
}

func (h *Http) Stop() error {
	return h.Closer.Close()
}

func (h *Http) Status() (string, error) {
	status := map[string]string{
		"host": h.Config.Host,
		"port": strconv.Itoa(h.Config.Port),
	}
	j, _ := json.Marshal(status)
	return string(j), nil
}

func (h *Http) Version() string {
	return VERSION
}
