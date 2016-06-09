package service

import (
	"net/http"

	"github.com/mholt/binding"
)

type AccountService struct {
	Store AccountStore
}

func NewAccountService(store AccountStore) *AccountService {
	return &AccountService{
		Store: store,
	}
}

func (self *AccountService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	candidate := Account{}
	if binding.Bind(r, &candidate).Handle(w) {
		return
	}
	account := candidate.ToModel()
	if err := self.Store.CreateAccount(account); err != nil {
		renderError(w, r, err)
		return
	}
	render(w, http.StatusCreated, account)
}
