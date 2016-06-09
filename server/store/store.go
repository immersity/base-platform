package store

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go sql/

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	*AccountStore

	close func() error
}

type Config struct {
	Dsn string
}

func New(c Config) (*Store, error) {
	db, err := sql.Open("mysql", c.Dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{
		NewAccountStore(db),
		func() error {
			return db.Close()
		},
	}, nil
}

func (self *Store) Close() error {
	return self.close()
}
