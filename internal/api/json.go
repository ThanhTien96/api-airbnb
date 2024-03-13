package api

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Paging  interface{} `json:"paging,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
	// Filter interface{} `json:"filter"`
}

func PagingSuccessResponse(message string, data, paging interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    200,
		Success: true,
		Data:    data,
		Paging:  paging,
	}
}

func DataSuccessResponse(message string, data interface{}, success bool) *SuccessResponse {
	return &SuccessResponse{Code: 200, Message: message, Data: data, Success: success}
}

func jsonError(code int, message string) interface{} {
	obj := map[string]interface{}{
		"success": false,
		"code":    code,
		"errors": []string{
			message,
		},
	}
	return obj
}

func JsonError(code int, message string) interface{} {
	return jsonError(code, message)
}

