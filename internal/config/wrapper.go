package config

import (
	"github.com/pangum/pangu"
)

type Wrapper struct {
	Retry   *Retry   `default:"{}" json:"retry,omitempty"`
	Runtime *Runtime `default:"{}" json:"runtime,omitempty"`
}

func newWrapper(config *pangu.Config) (wrapper *Wrapper, err error) {
	wrapper = new(Wrapper)
	err = config.Build().Get(wrapper)

	return
}
