package handlers_http_private_current_user_v1_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	user_store_svc_v1_mocks "github.com/golerplate/contracts/clients/user-store-svc/v1/mocks"
	models_http_common_account_v1 "github.com/golerplate/contracts/models/http/common/account/v1"
	"github.com/golerplate/pkg/constants"
	pkghttp "github.com/golerplate/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
	handlers_http_private_current_user_v1 "github.com/golerplate/user-gtw/internal/handlers/http/private/current-user/v1"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

func Test_Get(t *testing.T) {
	t.Run("ok - username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		// setup variables
		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		// setup service
		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		// setup http handler
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/private/v1/current-user", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("x-cuser", &entities_user_v1.User{
			ID:         userid,
			ExternalID: "external_xx",
			Username:   "testuser",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		})

		handler := handlers_http_private_current_user_v1.NewHandler(context.Background(), s)

		// prepare expected response
		data := handlers_http_private_current_user_v1.GetResponse{
			Account: &models_http_common_account_v1.PrivateAccount{
				User: &models_http_common_account_v1.PrivateUser{
					ID:       userid,
					Username: "testuser",
					Email:    "testuser@test.com",
				},
			},
		}

		formattedResponse := pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, data)
		resp, err := json.Marshal(formattedResponse)
		assert.NoError(t, err)

		// test
		if assert.NoError(t, handler.Get(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(resp), rec.Body.String())
		}
	})

	t.Run("nok - service fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		// setup service
		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		// setup http handler
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/private/v1/current-user", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := handlers_http_private_current_user_v1.NewHandler(context.Background(), s)

		formattedResponse := pkghttp.NewHTTPResponse(http.StatusInternalServerError, pkghttp.MessageInternalServerError, nil)
		resp, err := json.Marshal(formattedResponse)
		assert.NoError(t, err)

		// test
		if assert.NoError(t, handler.Get(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(resp), rec.Body.String())
		}
	})
}
