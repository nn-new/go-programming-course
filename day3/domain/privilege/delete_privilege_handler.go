package privilege

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deletePrivilege func(context.Context, primitive.ObjectID) error

func (fn deletePrivilege) DeletePrivilege(ctx context.Context, id primitive.ObjectID) error {
	return fn(ctx, id)
}

func DeletePrivilegeHandler(svc deletePrivilege) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}

		err = svc.DeletePrivilege(c.Request().Context(), objID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "deleted successfully",
		})
	}
}
