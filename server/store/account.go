package store

import (
	"database/sql"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/immersity/base-platform/server/model"
)

type AccountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) *AccountStore {
	return &AccountStore{db}
}

func (self *AccountStore) CreateAccount(account *model.Account) error {
	if err := account.PreInsert(); err != nil {
		return NewInternalErr(err)
	}
	res, err := self.db.Exec(
		sqlInsertAccount,
		account.Role,
		account.Verified,
		account.Email,
		account.Password,
		account.FirstName,
		account.LastName,
		account.CreatedOn,
		account.UpdatedOn,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
				return ErrDuplicateAccount
			}
		}
		return NewInternalErr(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return NewInternalErr(err)
	}
	account.ID = id
	return nil
}

func (self *AccountStore) CheckCredentials(email, password string) error {
	account := model.Account{}
	if err := self.db.QueryRow(sqlSelectAccountCredentials, email).Scan(
		&account.ID,
		&account.Role,
		&account.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return ErrInvalidCredentials
		} else {
			return NewInternalErr(err)
		}
	}
	if err := account.ComparePassword(password); err != nil {
		return ErrInvalidCredentials
	}
	return nil
}
