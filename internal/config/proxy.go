package config

import (
	"fmt"
	"net/url"
)

type Proxy struct {
	// 是否开启
	Enabled *bool `default:"true"`
	// 主机
	Host string `validate:"ip|hostname"`
	// 端口
	Port int `validate:"min=1,max=65535"`
	// 代理类型
	Scheme string `default:"scheme,omitempty" validate:"required,oneof=socks4 socks5 http https"`
	// 目标
	Target string
	// 排除
	Exclude string
	// 代理认证用户名
	Username string
	// 代理认证密码
	Password string
}

func newProxy(wrapper *Wrapper) *Proxy {
	return wrapper.Proxy
}

func (p *Proxy) Addr() (addr string) {
	if "" != p.Username && "" != p.Password {
		addr = fmt.Sprintf("%s://%s:%s@%s", p.Scheme, url.QueryEscape(p.Username), url.QueryEscape(p.Password), p.Host)
	} else {
		addr = fmt.Sprintf("%s://%s", p.Scheme, p.Host)
	}

	return
}
