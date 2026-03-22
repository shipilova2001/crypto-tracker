package dtos

type Response struct {
	Message string
	Data    *UserResponse
	Code    int
}
type ResponseJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
