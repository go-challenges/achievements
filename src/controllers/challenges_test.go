package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChallenges_Create(t *testing.T) {
	type args struct {
		w    *httptest.ResponseRecorder
		r    *http.Request
		code int
	}
	stub_body := map[string]interface{}{
		"name":        "challenge 12",
		"description": "descrition 12",
	}
	json_stub, _ := json.Marshal(stub_body)
	//mockCtrl := gomock.NewController(t)
	//challenge_stub := mock_models.NewMockChallengeRepository(mockCtrl)

	tests := []struct {
		name string
		c    *Challenges
		args args
	}{
		{
			"return a error during body parse",
			&Challenges{},
			args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, "/challenges", bytes.NewBuffer([]byte{})),
				http.StatusBadRequest,
			},
		},
		{
			"return created",
			&Challenges{},
			args{
				httptest.NewRecorder(),
				httptest.NewRequest(
					http.MethodPost,
					"/challenges",
					bytes.NewReader(json_stub),
				),
				http.StatusCreated,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Challenges{}
			c.Create(tt.args.w, tt.args.r)

			if tt.args.w.Code != tt.args.code {
				t.Errorf("got HTTP status code %d, expected %d", tt.args.w.Code, tt.args.code)
			}

		})
	}
}
