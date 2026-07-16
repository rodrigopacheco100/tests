package shared

type PaginationResult[T any] struct {
	Items      []T `json:"items"`
	TotalPages int `json:"totalPages"`
	Total      int `json:"total"`
}

type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (r *PaginationRequest) Offset() int {
	return (r.Page - 1) * r.Limit
}
