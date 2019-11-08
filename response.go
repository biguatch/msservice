package msservice

type Error struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (error *Error) Error() string {
	return error.Message
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}
