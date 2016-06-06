package mongo

import (
	"gopkg.in/mgo.v2"
)

type Store struct {
	*AccountStore

	_close func() error
}

type Config struct {
	DB       string
	MongoDsn string
}

func New(c Config) (*Store, error) {
	session, err := mgo.Dial(c.MongoDsn)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return &Store{
		NewAccountStore(c.DB, session),
		func() error {
			session.Close()
			return nil
		},
	}, nil
}

func (self *Store) Close() error {
	return self._close()
}
