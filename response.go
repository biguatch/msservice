package msservice

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (error *Error) Error() string {
	return error.Message
}

type ResponseInterface interface {
	SendResponse(w http.ResponseWriter, status int)
}

type Response struct {
	Content interface{} `json:"content,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

func (r *Response) SendResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if r.Content != nil || r.Meta != nil || r.Error != nil {
		_ = json.NewEncoder(w).Encode(r)
	}
}

func SendQuickResponse(w http.ResponseWriter, status int) {
	resp := Response{}
	resp.SendResponse(w, status)
}

func SendError(w http.ResponseWriter, err error, status int) {
	resp := Response{
		Error: &Error{
			Message: err.Error(),
		},
	}

	resp.SendResponse(w, status)
}

func SendData(w http.ResponseWriter, content interface{}, meta interface{}, status int) {
	resp := Response{
		Content: content,
		Meta:    meta,
	}

	resp.SendResponse(w, status)
}
