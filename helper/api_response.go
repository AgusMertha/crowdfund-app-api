package helper

type response struct {
	Meta meta `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	metadata := meta{message, code, status}

	responseData := response{metadata,data}

	return responseData

}