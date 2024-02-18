package service_v1_test

import (
	"context"
	"testing"
	"time"

	user_store_svc_v1_entities "github.com/Golerplate/contracts/clients/user-store-svc/v1/entities"
	user_store_svc_v1_mocks "github.com/Golerplate/contracts/clients/user-store-svc/v1/mocks"
	"github.com/Golerplate/pkg/constants"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

func Setup() {
}

func TearDown() {
}

func Test_GetUserByIdentifier(t *testing.T) {
	t.Run("ok - user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(&user_store_svc_v1_entities.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, user)
	})

	t.Run("ok - username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&user_store_svc_v1_entities.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByIdentifier(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, user)
	})

	t.Run("nok - user not found by id", func(t *testing.T) {
	})

	t.Run("nok - user not found by username", func(t *testing.T) {
	})

	t.Run("client error", func(t *testing.T) {
	})
}
