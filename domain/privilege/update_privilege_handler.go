package privilege

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type updatePrivileges func(context.Context, Privilege) error

func (fn updatePrivileges) UpdatePrivilege(ctx context.Context, privilege Privilege) error {
	return fn(ctx, privilege)
}

func UpdatePrivilegeHandler(svc updatePrivileges) echo.HandlerFunc {
	return func(c echo.Context) error {

		var request Privilege
		if err := c.Bind(&request); err != nil {
			return err
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := svc.UpdatePrivilege(c.Request().Context(), request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "updated successfully",
		})
	}
}
