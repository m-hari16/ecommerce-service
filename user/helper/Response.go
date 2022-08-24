package helper

type Response struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
