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

type ResponseToken struct {
	Token 	string `json:"token"`
}
type ResponseError struct {
	Error 	string `json:"error"`
}