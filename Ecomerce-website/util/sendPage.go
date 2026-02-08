package util

import (
	"ecomerce/domain"
	"net/http"
)

type PaginatedData struct {
	Data any `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Data []*domain.Products `json:"data"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}


func SendPage(w http.ResponseWriter,data any,page,limit,cnt int64) {

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Limit: limit,
		TotalItems: cnt,
		TotalPages: cnt/limit,

		},
	}
  
	SendData(w,http.StatusOK,paginatedData)

}