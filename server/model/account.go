package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        int64     `json:"id"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

func (a *Account) PreInsert() error {
	if err := a.HashPassword(); err != nil {
		return err
	}

	a.Role = "user"

	now := time.Now().UTC()
	a.CreatedOn = now
	a.UpdatedOn = now

	return nil
}

func (a *Account) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}

func (a *Account) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
