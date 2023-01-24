package privilege

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createPrivileges func(context.Context, Privilege) error

func (fn createPrivileges) CreatePrivilege(ctx context.Context, privilege Privilege) error {
	return fn(ctx, privilege)
}

func CreatePrivilegeHandler(svc createPrivileges) echo.HandlerFunc {
	return func(c echo.Context) error {

		var request Privilege
		if err := c.Bind(&request); err != nil {
			return err
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := svc.CreatePrivilege(c.Request().Context(), request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
