package pagination

import "math"

type Request struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type Response struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func GetPagination(req Request) (offset int, limit int) {
	if req.Page <= 0 {
		req.Page = 1
	}

	maxLimit := 100
	limit = req.Limit
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}

	offset = (req.Page - 1) * limit
	return offset, limit
}

func GetPaginationResponse(page int, limit int, total int) Response {
	var result Response
	if page <= 0 {
		page = 1
	}
	result.LastPage = int(math.Ceil(float64(total) / float64(limit)))
	result.CurrentPage = page
	result.Total = total
	result.PerPage = limit
	return result
}
