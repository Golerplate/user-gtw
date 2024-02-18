package config

import (
	pkgconfig "github.com/Golerplate/pkg/config"
	pkghttp "github.com/Golerplate/pkg/http"
)

type Config struct {
	pkgconfig.ServiceConfig
	pkghttp.HTTPServerConfig
}
