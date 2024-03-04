package config

import (
	contracts_user_store_svc "github.com/golerplate/contracts/clients/user-store-svc"
	pkgconfig "github.com/golerplate/pkg/config"
	pkghttp "github.com/golerplate/pkg/http"
)

type Config struct {
	pkgconfig.ServiceConfig
	pkghttp.HTTPServerConfig
	contracts_user_store_svc.UserStoreSvcConfig
}
