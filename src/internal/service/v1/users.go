package service_v1

import (
	"context"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
	constants "github.com/golerplate/pkg/constants"

	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
)

func (s *Service) CreateUser(ctx context.Context, req *entities_user_v1.UserCreate) (*entities_user_v1.User, error) {
	user, err := s.userStoreSvc.CreateUser(ctx, &user_store_svc_v1_entities.UserCreate{
		ExternalID: req.ExternalID,
		Email:      req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &entities_user_v1.User{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

func (s *Service) UpdateUsername(ctx context.Context, cuser *entities_user_v1.User, username string) (*entities_user_v1.User, error) {
	user, err := s.userStoreSvc.UpdateUsername(ctx, cuser.ID, username)
	if err != nil {
		return nil, err
	}

	return &entities_user_v1.User{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

func (s *Service) GetUserByIdentifier(ctx context.Context, identifier string) (*entities_user_v1.User, error) {
	var user *user_store_svc_v1_entities.User
	var err error

	if constants.User.IsValid(identifier) {
		user, err = s.userStoreSvc.GetUserByID(ctx, identifier)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.userStoreSvc.GetUserByUsername(ctx, identifier)
		if err != nil {
			return nil, err
		}
	}

	return &entities_user_v1.User{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}
