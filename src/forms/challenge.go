package forms

import (
	"achievements/src/helpers"

	"encoding/json"
	"net/http"
)

type Challenge struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	Properties  helpers.JSONB `json:"properties"`
}

func (f *Challenge) Bind(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		return err
	}

	return nil
}
