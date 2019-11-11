package msservice

import (
	"encoding/json"
	"net/http"
	"reflect"
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
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

func (r *Response) SendResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(r)
}

func SendQuickResponse(w http.ResponseWriter, status int) {
	resp := Response{}
	resp.SendResponse(w, status)
}

func SendError(w http.ResponseWriter, err interface{}, status int) {
	switch reflect.ValueOf(&err).Elem().Interface().(type) {
	case error:
		err = &Error{Message: err.(string)}
	default:
		// And here I'm feeling dumb. ;)
		err = &Error{Message: "unknown error"}
	}

	resp := Response{
		Error: err.(*Error),
	}

	resp.SendResponse(w, status)
}
