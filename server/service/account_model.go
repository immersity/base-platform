package service

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/immersity/base-platform/server/model"
	"github.com/mholt/binding"
)

type NewAccount struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (self *NewAccount) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&self.Email:     newRequiredField("email"),
		&self.Password:  newRequiredField("password"),
		&self.FirstName: newRequiredField("firstName"),
		&self.LastName:  newRequiredField("lastName"),
	}
}

func (self NewAccount) Validate(r *http.Request, errs binding.Errors) binding.Errors {
	if !valid.IsEmail(self.Email) {
		errs = append(errs, newValidationError("email", "Correo electrónico inválido"))
	}
	if !valid.IsAlphanumeric(self.Password) {
		errs = append(errs, newValidationError("password", "Contraseña inválida"))
	}
	return errs
}

func (self *NewAccount) ToModel() *model.Account {
	return &model.Account{
		Email:     self.Email,
		Password:  self.Password,
		FirstName: self.FirstName,
		LastName:  self.LastName,
	}
}
