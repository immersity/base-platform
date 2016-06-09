package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type AuthStoreMock struct{}

func (s *AuthStoreMock) CheckCredentials(email, password string) error {
	if email != "foo@bar.com" || password != "foobar" {
		return errors.New("store: invalid credentials")
	}
	return nil
}

var authService *AuthService = NewAuthService(&AuthStoreMock{}, "olakease", time.Hour*1)

func TestCreateTokenEndpoint(t *testing.T) {
	cases := []struct {
		body         interface{}
		response     *httptest.ResponseRecorder
		expectedCode int
	}{
		{
			nil,
			httptest.NewRecorder(),
			http.StatusBadRequest,
		},
		{
			map[string]string{
				"email":    "foo@bar.com",
				"lastName": "Bar",
			},
			httptest.NewRecorder(),
			http.StatusBadRequest,
		},
		{
			map[string]string{
				"email":    "foo@bar.com",
				"password": "foobar",
			},
			httptest.NewRecorder(),
			http.StatusCreated,
		},
		{
			map[string]string{
				"email":    "foo@bar.com",
				"password": "barfoo",
			},
			httptest.NewRecorder(),
			http.StatusInternalServerError,
		},
	}
	for _, c := range cases {
		jsonBody, err := json.Marshal(c.body)
		if err != nil {
			t.Errorf(err.Error())
		}
		request, err := http.NewRequest("POST", "", bytes.NewReader(jsonBody))
		if err != nil {
			t.Errorf(err.Error())
		}
		request.Header.Set("Content-Type", "application/json")
		authService.CreateToken(c.response, request)
		if c.expectedCode != c.response.Code {
			t.Errorf("want: %d, got: %d", c.expectedCode, c.response.Code)
		}
	}
}
