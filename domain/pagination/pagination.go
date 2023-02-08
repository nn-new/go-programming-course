package pagination

import (
	"strconv"
	"strings"
)

type Pagination struct {
	Search    string `query:"search"`
	Sort      string `query:"sort"`
	Direction string `query:"direction"`
	Page      string `query:"page"`
	PageSize  string `query:"pageSize"`
}

func (p Pagination) GetDirection() int {
	if strings.ToLower(p.Direction) == "asc" {
		return 1
	}
	return -1
}

func (p Pagination) GetPage() int64 {
	page, err := strconv.Atoi(p.Page)
	if err != nil {
		page = 1
	}
	return int64(page)
}

func (p Pagination) GetPageSize() int64 {
	pageSize, err := strconv.Atoi(p.PageSize)
	if err != nil {
		pageSize = 10
	}
	return int64(pageSize)
}
