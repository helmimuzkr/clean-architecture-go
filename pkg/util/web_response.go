package util

type PaginationResponse struct {
	Page        int `json:"page"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
}

type WithPagination struct {
	Pagination PaginationResponse `json:"pagination"`
	Data       interface{}        `json:"data"`
	Message    string             `json:"message"`
}

func SuccessResponse(code int, message string, data ...any) (int, map[string]interface{}) {
	response := make(map[string]interface{})
	response["message"] = message
	response["status"] = code

	for _, v := range data {
		response["data"] = v
	}

	return code, response
}
