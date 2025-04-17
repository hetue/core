package config

import (
	"github.com/harluo/boot"
)

type Wrapper struct {
	Retry   *Retry   `default:"{}" json:"retry,omitempty"`
	Runtime *Runtime `default:"{}" json:"runtime,omitempty"`
}

func newWrapper(config *boot.Config) (wrapper *Wrapper, err error) {
	wrapper = new(Wrapper)
	err = config.Build().Get(wrapper)

	return
}
