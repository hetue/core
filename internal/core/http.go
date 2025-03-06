package core

import (
	"github.com/goexl/http"
	"github.com/hetue/core/internal/config"
)

type Http struct {
	proxy *config.Proxy
}

func newHttp(proxy *config.Proxy) *Http {
	return &Http{
		proxy: proxy,
	}
}

func (h *Http) New() *http.Client {
	http.New()
}

func (h *Http) setup() {
	client := http.New()
	if nil == h.Proxies {
		h.Proxies = make([]*config.Proxy, 0)
	}
	if nil != h.Proxy {
		h.Proxies = append(h.Proxies, h.Proxy)
	}

	for _, _proxy := range h.Proxies {
		builder := client.Proxy()
		builder.Host(_proxy.Host)
		builder.Scheme(_proxy.Scheme)
		builder.Port(_proxy.Port)
		builder.Target(_proxy.Target)
		builder.Exclude(_proxy.Exclude)
		client = builder.Build()
	}
	h.http = client.Build()
}
