package handlers_http_internal_users_v1

import (
	"net/http"

	models_http_common_account_v1 "github.com/golerplate/contracts/models/http/common/account/v1"
	pkghttp "github.com/golerplate/pkg/http"
	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type CreateRequest struct {
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
}

type CreateResponse struct {
	Account *models_http_common_account_v1.Account `json:"account"`
}

func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("handlers.http.internal.users.v1.create.Handler.Create: can not bind request")
		return c.JSON(http.StatusBadRequest, pkghttp.NewHTTPResponse(http.StatusBadRequest, pkghttp.MessageBadRequestError, nil))
	}

	user, err := h.service.CreateUser(ctx, &entities_user_v1.UserCreate{
		ExternalID: req.ExternalID,
		Email:      req.Email,
	})
	if err != nil {
		return c.JSON(pkghttp.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, CreateResponse{
		Account: &models_http_common_account_v1.Account{
			User: &models_http_common_account_v1.User{
				ID:       user.ID,
				Username: user.Username,
			},
		},
	}))
}
