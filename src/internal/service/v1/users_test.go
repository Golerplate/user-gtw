package service_v1_test

import (
	"context"
	"testing"
	"time"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
	user_store_svc_v1_mocks "github.com/golerplate/contracts/clients/user-store-svc/v1/mocks"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

func Test_GetUserByIdentifier(t *testing.T) {
	t.Run("ok - user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(&user_store_svc_v1_entities.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     true,
			ProfilePicture: "https://test.com/testuser.jpg",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			IsVerified:     true,
			ProfilePicture: "https://test.com/testuser.jpg",
			CreatedAt:      created,
		}, user)
	})

	t.Run("ok - username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&user_store_svc_v1_entities.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     true,
			ProfilePicture: "https://test.com/testuser.jpg",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			IsVerified:     true,
			ProfilePicture: "https://test.com/testuser.jpg",
			CreatedAt:      created,
		}, user)
	})

	t.Run("nok - user not found by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)

		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(nil, pkgerrors.NewNotFoundError("not found"))

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), userid)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.True(t, pkgerrors.IsNotFoundError(err))
	})

	t.Run("nok - user not found by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(nil, pkgerrors.NewNotFoundError("not found"))

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), "testuser")
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.True(t, pkgerrors.IsNotFoundError(err))
	})

	t.Run("nok - client error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(nil, pkgerrors.NewInternalServerError("failed"))

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), "testuser")
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.True(t, pkgerrors.IsInternalServerError(err))
	})
}
