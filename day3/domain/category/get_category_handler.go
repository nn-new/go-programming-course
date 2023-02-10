package category

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getCategoriesFunc func(context.Context) ([]Category, error)

func (fn getCategoriesFunc) GetCategories(ctx context.Context) ([]Category, error) {
	return fn(ctx)
}

func GetCatalogHandler(svc getCategoriesFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := svc.GetCategories(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, categories)
	}
}
