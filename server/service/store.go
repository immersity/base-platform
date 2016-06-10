package service

import (
	"github.com/immersity/base-platform/server/model"
)

type AuthStore interface {
	CheckCredentials(email, password string) (string, error)
}

type AccountStore interface {
	CreateAccount(a *model.Account) error
}

type Store interface {
	AuthStore
	AccountStore
}
