package service_v1

import (
	"context"

	user_store_svc_v1 "github.com/golerplate/contracts/clients/user-store-svc/v1"
)

type Service struct {
	userStoreSvc user_store_svc_v1.UserStoreSvc
}

func NewService(ctx context.Context,
	userStoreSvc user_store_svc_v1.UserStoreSvc,
) (*Service, error) {
	return &Service{
		userStoreSvc: userStoreSvc,
	}, nil
}
