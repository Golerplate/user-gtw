package config

import (
	pkgconfig "github.com/golerplate/pkg/config"
	pkghttp "github.com/golerplate/pkg/http"
)

type Config struct {
	pkgconfig.ServiceConfig
	pkghttp.HTTPServerConfig
}
