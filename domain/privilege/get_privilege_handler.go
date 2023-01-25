package privilege

import (
	"context"
	"net/http"
	"privilege/domain/pagination"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getPrivileges func(context.Context, pagination.Pagination) ([]Privilege, error)

func (fn getPrivileges) GetPrivilege(ctx context.Context, pag pagination.Pagination) ([]Privilege, error) {
	return fn(ctx, pag)
}

func GetPrivilegeHandler(svc getPrivileges) echo.HandlerFunc {
	return func(c echo.Context) error {
		var pagination pagination.Pagination
		err := c.Bind(&pagination)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		privileges, err := svc.GetPrivilege(c.Request().Context(), pagination)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, privileges)
	}
}

type getPrivilegesById func(context.Context, primitive.ObjectID) (Privilege, error)

func (fn getPrivilegesById) GetPrivilegeById(ctx context.Context, id primitive.ObjectID) (Privilege, error) {
	return fn(ctx, id)
}

func GetPrivilegeByIdHandler(svc getPrivilegesById) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}

		privilege, err := svc.GetPrivilegeById(c.Request().Context(), objID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, privilege)
	}
}
