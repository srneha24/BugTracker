package utils

import "math"

type PaginatedResponse[T any] struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       []T    `json:"data"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	ToalPages  int    `json:"total_page"`
	NextPage   *int   `json:"next_page"`
	PrevPage   *int   `json:"prev_page"`
}

type PaginationQueryParams struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=10" binding:"min=1"`
}

func Paginate[T any](data []T, page, limit, totalCount int) PaginatedResponse[T] {
	var paginatedResponse PaginatedResponse[T]

	paginatedResponse.Success = true
	paginatedResponse.Message = "Request Successful"
	paginatedResponse.Data = data
	paginatedResponse.TotalCount = totalCount
	paginatedResponse.ToalPages = int(math.Ceil(float64(totalCount) / float64(limit)))
	paginatedResponse.Page = page
	paginatedResponse.Limit = limit

	if paginatedResponse.ToalPages-paginatedResponse.Page > 0 {
		next := paginatedResponse.Page + 1
		paginatedResponse.NextPage = &next
	}

	if paginatedResponse.ToalPages-paginatedResponse.Page >= 0 && paginatedResponse.Page-1 > 0 {
		prev := paginatedResponse.Page - 1
		paginatedResponse.PrevPage = &prev
	}

	return paginatedResponse
}
