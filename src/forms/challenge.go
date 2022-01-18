package forms

import (
	"encoding/json"
	"net/http"
)

type Challenge struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (f *Challenge) Bind(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		return err
	}

	return nil
}
