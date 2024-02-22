package handlers_http_private_current_user_v1

import (
	"net/http"

	models_http_common_account_v1 "github.com/golerplate/contracts/models/http/common/account/v1"
	pkghttp "github.com/golerplate/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	entities_user_v1 "github.com/golerplate/user-gtw/internal/entities/user/v1"
)

type UpdateUsernameRequest struct {
	Username string `json:"username"`
}

type UpdateUsernameResponse struct {
	Account *models_http_common_account_v1.PrivateAccount `json:"account"`
}

func (h *Handler) UpdateUsername(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("x-cuser").(*entities_user_v1.User)
	if !ok {
		log.Error().Msg("handlers.http.private.current-user.v1.get.Handler.Get: can not get user from context")
		return c.JSON(http.StatusInternalServerError, pkghttp.NewHTTPResponse(http.StatusInternalServerError, pkghttp.MessageInternalServerError, nil))
	}

	var req UpdateUsernameRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("handlers.http.private.current-user.v1.update_username.Handler.UpdateUsername: can not bind request")
		return c.JSON(http.StatusBadRequest, pkghttp.NewHTTPResponse(http.StatusBadRequest, pkghttp.MessageBindError, nil))
	}

	returnedUser, err := h.service.UpdateUsername(ctx, user, req.Username)
	if err != nil {
		return c.JSON(pkghttp.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, UpdateUsernameResponse{
		Account: &models_http_common_account_v1.PrivateAccount{
			User: &models_http_common_account_v1.PrivateUser{
				ID:       returnedUser.ID,
				Username: returnedUser.Username,
				Email:    returnedUser.Email,
			},
		},
	}))
}
