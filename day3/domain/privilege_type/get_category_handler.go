package privilegetype

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getPrivilegeTypeFunc func(context.Context) ([]PrivilegeType, error)

func (fn getPrivilegeTypeFunc) GetPrivilegeType(ctx context.Context) ([]PrivilegeType, error) {
	return fn(ctx)
}

func GetPrivilegeTypeHandler(svc getPrivilegeTypeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		privilegeTypes, err := svc.GetPrivilegeType(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, privilegeTypes)
	}
}
