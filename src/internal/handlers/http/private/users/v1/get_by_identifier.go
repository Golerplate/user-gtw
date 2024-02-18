package handlers_http_private_users_v1

import (
	"net/http"

	models_http_common_account_v1 "github.com/golerplate/contracts/models/http/common/account/v1"
	pkghttp "github.com/golerplate/pkg/http"
	"github.com/labstack/echo/v4"
)

type GetByIdentifierResponse struct {
	Account *models_http_common_account_v1.Account `json:"account"`
}

func (h *Handler) GetByIdentifier(c echo.Context) error {
	ctx := c.Request().Context()

	identifier := c.Param("identifier")
	if identifier == "" {
		return c.JSON(http.StatusBadRequest, pkghttp.NewHTTPResponse(http.StatusBadRequest, pkghttp.MessageBadRequestError, nil))
	}

	user, err := h.service.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		return c.JSON(pkghttp.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, GetByIdentifierResponse{
		Account: &models_http_common_account_v1.Account{
			User: &models_http_common_account_v1.User{
				ID:             user.ID,
				Username:       user.Username,
				IsVerified:     user.IsVerified,
				ProfilePicture: user.ProfilePicture,
				CreatedAt:      user.CreatedAt,
			},
		},
	}))
}
