package service

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/mholt/binding"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (self *Credentials) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&self.Email:    newRequiredField("email"),
		&self.Password: newRequiredField("password"),
	}
}

func (self Credentials) Validate(r *http.Request, errs binding.Errors) binding.Errors {
	if !valid.IsEmail(self.Email) {
		errs = append(errs, newValidationError("email", "Correo electrónico inválido"))
	}
	if !valid.IsAlphanumeric(self.Password) {
		errs = append(errs, newValidationError("password", "Contraseña inválida"))
	}
	return errs
}
