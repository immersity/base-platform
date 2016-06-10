package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/immersity/base-platform/server/model"
	"github.com/immersity/base-platform/server/store"
)

type AccountStoreMock struct {
	Accounts map[string]bool
}

func (s *AccountStoreMock) CreateAccount(a *model.Account) error {
	if exists := s.Accounts[a.Email]; exists {
		return store.ErrDuplicateAccount
	}
	s.Accounts[a.Email] = true
	return nil
}

var service *AccountService = NewAccountService(&AccountStoreMock{
	Accounts: make(map[string]bool, 0),
})

func TestCreateAccountEndpoint(t *testing.T) {
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
				"email":     "foobar.com",
				"password":  "foo-bar",
				"firstName": "Foo",
				"lastName":  "Bar",
			},
			httptest.NewRecorder(),
			http.StatusBadRequest,
		},
		{
			map[string]string{
				"email":     "foo@bar.com",
				"password":  "foobar",
				"firstName": "Foo",
				"lastName":  "Bar",
			},
			httptest.NewRecorder(),
			http.StatusCreated,
		},
		{
			map[string]string{
				"email":     "foo@bar.com",
				"password":  "foobar",
				"firstName": "Foo",
				"lastName":  "Bar",
			},
			httptest.NewRecorder(),
			http.StatusConflict,
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
		service.CreateAccount(c.response, request)
		if c.expectedCode != c.response.Code {
			t.Errorf("want: %d, got: %d", c.expectedCode, c.response.Code)
		}
	}
}
