package model

type ResponseWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FailureResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
