package catalog

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getCatalogsFunc func(context.Context) ([]Catalog, error)

func (fn getCatalogsFunc) GetCatalogs(ctx context.Context) ([]Catalog, error) {
	return fn(ctx)
}

func GetCatalogHandler(svc getCatalogsFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		catalogs, err := svc.GetCatalogs(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, catalogs)
	}
}
