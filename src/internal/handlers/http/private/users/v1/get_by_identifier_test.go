package handlers_http_private_users_v1_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
	user_store_svc_v1_mocks "github.com/golerplate/contracts/clients/user-store-svc/v1/mocks"
	models_http_common_account_v1 "github.com/golerplate/contracts/models/http/common/account/v1"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	pkghttp "github.com/golerplate/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	handlers_http_private_users_v1 "github.com/golerplate/user-gtw/internal/handlers/http/private/users/v1"
	service_v1 "github.com/golerplate/user-gtw/internal/service/v1"
)

func Test_GetByIdentifier(t *testing.T) {
	t.Run("ok - username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		// setup variables
		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		// setup mocks
		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&user_store_svc_v1_entities.User{
			ID:         userid,
			ExternalID: "external_xx",
			Username:   "testuser",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		// setup service
		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		// setup http handler
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/private/v1/users/testuser", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetParamNames("identifier")
		c.SetParamValues("testuser")

		handler := handlers_http_private_users_v1.NewHandler(context.Background(), s)

		// prepare expected response
		data := handlers_http_private_users_v1.GetByIdentifierResponse{
			Account: &models_http_common_account_v1.Account{
				User: &models_http_common_account_v1.User{
					ID:       userid,
					Username: "testuser",
				},
			},
		}

		formattedResponse := pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, data)
		resp, err := json.Marshal(formattedResponse)
		assert.NoError(t, err)

		// test
		if assert.NoError(t, handler.GetByIdentifier(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(resp), rec.Body.String())
		}
	})

	t.Run("nok - service fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		// setup mocks
		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(nil, pkgerrors.NewInternalServerError("internal server error"))

		// setup service
		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		// setup http handler
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/private/v1/users/testuser", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetParamNames("identifier")
		c.SetParamValues("testuser")

		handler := handlers_http_private_users_v1.NewHandler(context.Background(), s)

		formattedResponse := pkghttp.NewHTTPResponse(http.StatusInternalServerError, pkghttp.MessageInternalServerError, nil)
		resp, err := json.Marshal(formattedResponse)
		assert.NoError(t, err)

		// test
		if assert.NoError(t, handler.GetByIdentifier(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(resp), rec.Body.String())
		}
	})

	t.Run("nok - missing identifier in url", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		// setup service
		s, err := service_v1.NewService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		// setup http handler
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/private/v1/users/testuser", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := handlers_http_private_users_v1.NewHandler(context.Background(), s)

		formattedResponse := pkghttp.NewHTTPResponse(http.StatusBadRequest, pkghttp.MessageBadRequestError, nil)
		resp, err := json.Marshal(formattedResponse)
		assert.NoError(t, err)

		// test
		if assert.NoError(t, handler.GetByIdentifier(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(resp), rec.Body.String())
		}
	})
}
