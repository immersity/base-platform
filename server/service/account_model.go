package service

import (
	"net/http"

	"github.com/immersity/base-platform/server/model"
	"github.com/mholt/binding"
)

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (self *account) FieldMap(r *http.Request) binding.FieldMap {
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

func (self *account) ToModel() *model.Account {
	return &model.Account{
		Email:    self.Email,
		Password: self.Password,
	}
}
