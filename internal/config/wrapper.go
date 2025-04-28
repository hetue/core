package config

import (
	"github.com/harluo/config"
)

type Wrapper struct {
	Retry   *Retry   `default:"{}" json:"retry,omitempty"`
	Runtime *Runtime `default:"{}" json:"runtime,omitempty"`
}

func newWrapper(getter config.Getter) (wrapper *Wrapper, err error) {
	wrapper = new(Wrapper)
	err = getter.Get(wrapper)

	return
}
