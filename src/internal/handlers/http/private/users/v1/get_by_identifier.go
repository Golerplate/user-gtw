package handlers_http_private_users_v1

import "github.com/labstack/echo"

type GetByIdentifierResponse struct {
}

func (h *Handler) GetByIdentifier(c echo.Context) error {
	ctx := c.Request().Context()

	identifier := c.Param("identifier")
	if identifier == "" {
		return c.JSON(400, "identifier is required")
	}

	user, err := h.service.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, user)
}
