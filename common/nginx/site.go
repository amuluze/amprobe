// Package nginx
// Date: 2024/8/6 16:42
// Author: Amu
// Description:
package nginx

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type UrlPath struct {
	Op      string `json:"op,omitempty"`
	UrlPath string `json:"url_path,omitempty"`
}

type Site struct {
	Name          string          `json:"name,omitempty" validate:"required,excludesall= _"`
	ServerNames   []string        `json:"server_names,omitempty" validate:"required,gt=0,dive"`
	UrlPaths      []*UrlPath      `json:"url_paths,omitempty"`
	ProxyConfig   *ProxyConfig    `json:"proxy_config,omitempty"`
	ListenConfigs []*ListenConfig `json:"listen_configs,omitempty" validate:"gt=0,dive"`
}

type SiteGroup struct {
	Sites []*Site `json:"sites,omitempty"`
}

type ProxyConfig struct {
	Upstreams        []*Upstream `json:"upstreams,omitempty" validate:"gt=0,dive"`
	Keepalive        int64       `json:"keepalive"`
	KeepaliveTimeout int64       `json:"keepalive_timeout,omitempty"`
	Schema           string      `json:"schema,omitempty" validate:"oneof=http https"`
}

type Upstream struct {
	Address string `json:"address,omitempty"`
	Port    uint16 `json:"port,omitempty"`
	Weight  uint   `json:"weight,omitempty"`
}

func (s *Upstream) Name() string {
	return net.JoinHostPort(s.Address, strconv.Itoa(int(s.Port)))
}

type ListenConfig struct {
	Address string `json:"address,omitempty"`
	Port    uint16 `json:"port,omitempty"`
}

func (l *ListenConfig) ListenAddress() string {
	if l.Port == 0 && strings.HasPrefix(l.Address, "unix:") {
		return l.Address
	} else if l.Address != "" {
		return net.JoinHostPort(l.Address, strconv.Itoa(int(l.Port)))
	} else {
		return fmt.Sprintf("%d", l.Port)
	}
}
