package api

type ErrorResponse struct {
	Errors map[string][]string `json:"errors"`
}

func NewErrorResponse() ErrorResponse {
	return ErrorResponse{Errors: map[string][]string{}}
}

type Pager struct {
	Total   int64 `json:"total"`
	Current int32 `json:"current" binding:"required"`
	Size    int32 `json:"size" binding:"required"`
}
