package privilege

import (
	"net/http"
	"privilege/domain/pagination"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/xuri/excelize/v2"
)

func DownloadPrivilegeHandler(svc getPrivileges) echo.HandlerFunc {
	return func(c echo.Context) error {
		var pagination pagination.Pagination
		err := c.Bind(&pagination)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		pagination.IsDownload = true

		privileges, err := svc.GetPrivilege(c.Request().Context(), pagination)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		now := time.Now().UnixMilli()
		filename := "privilege_export_" + strconv.Itoa(int(now)) + ".xlsx"
		c.Response().Header().Set("Content-Type", "application/octet-stream")
		c.Response().Header().Set("Content-Disposition", "attachment;filename="+filename)

		err = createExcel(privileges, c)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		return c.Attachment(filename, filename)
	}
}

func createExcel(privilege []Privilege, c echo.Context) error {
	f := excelize.NewFile()

	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	if err != nil {
		return err
	}

	header := []interface{}{}
	for _, cell := range []string{
		"Title", "Full Description", "Sort Description",
	} {
		header = append(header, cell)
	}

	if err := streamWriter.SetRow("A1", header, excelize.RowOpts{StyleID: style}); err != nil {
		return err
	}

	for i, u := range privilege {
		row := make([]interface{}, len(header))
		row[0] = u.Title
		row[1] = u.FullDescription
		row[2] = u.ShortDescription

		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		if err := streamWriter.SetRow(cell, row); err != nil {
			return err
		}
	}

	if err := streamWriter.Flush(); err != nil {
		return err
	}

	if err := f.Write(c.Response()); err != nil {
		return err
	}

	return nil
}
