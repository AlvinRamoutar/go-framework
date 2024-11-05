package run

import (
)

type RunConfig struct {
}

func (r *RunConfig) New() *RunConfig {
	rc := RunConfig{}
	return &rc
}