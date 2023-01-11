package domain

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code         string        `json:"code"`
	ErrorDetails []ErrorDetail `json:"error_details,omitempty"`
	DebugMessage string        `json:"debug_message,omitempty"`
}

type ErrorDetail struct {
	Field        string `json:"field"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type FilterParam struct {
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
	SortBy   string `form:"sort_by"`
	SortType string `form:"sort_type"`
}
