package msservice

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
