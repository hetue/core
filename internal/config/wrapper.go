package config

import (
	"github.com/pangum/pangu"
)

type Wrapper struct {
	Logging *Logging
	Proxy   *Proxy
	Retry   *Retry
	Runtime *Runtime
}

func newWrapper(config *pangu.Config) (wrapper *Wrapper, err error) {
	wrapper = new(Wrapper)
	err = config.Build().Get(wrapper)

	return
}
