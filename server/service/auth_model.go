package service

import (
	"net/http"

	"github.com/mholt/binding"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (self *credentials) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&self.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&self.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}
