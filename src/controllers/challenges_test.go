package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"achievements/src/models"
	mock_repository "achievements/src/repository/mock"

	"github.com/golang/mock/gomock"
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
	mockCtrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockChallengesRepository(mockCtrl)

	tests := []struct {
		name    string
		c       *Challenges
		useStab bool
		args    args
	}{
		{
			"return a error during body parse",
			&Challenges{},
			false,
			args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, "/challenges", bytes.NewBuffer([]byte{})),
				http.StatusBadRequest,
			},
		},
		{
			"return created",
			ChallengesWrapper(mockRepo),
			true,
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
			c := tt.c

			if tt.useStab {
				challenge := models.Challenge{
					Name:        "challenge 12",
					Description: "descrition 12",
				}
				mockRepo.EXPECT().Create(&challenge).Return(nil)
			}

			c.Create(tt.args.w, tt.args.r)

			if tt.args.w.Code != tt.args.code {
				t.Errorf("got HTTP status code %d, expected %d", tt.args.w.Code, tt.args.code)
			}

		})
	}
}
