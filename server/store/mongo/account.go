package mongo

import (
	"errors"
	"time"

	"github.com/immersity/base-platform/server/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AccountStore struct {
	DB   string
	Coll string

	Session *mgo.Session
}

func NewAccountStore(db string, session *mgo.Session) *AccountStore {
	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}

	if err := session.DB(db).C("accounts").EnsureIndex(emailIndex); err != nil {
		panic(err)
	}

	return &AccountStore{
		DB:      db,
		Coll:    "accounts",
		Session: session,
	}
}

func (self *AccountStore) CheckCredentials(email, password string) error {
	account := model.Account{}

	query := bson.M{
		"email": email,
	}

	if err := self.Session.DB(self.DB).C(self.Coll).Find(query).One(&account); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return ErrInvalidCredentials
	}

	return nil
}

func (self *AccountStore) CreateAccount(a *model.Account) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.ID = bson.NewObjectId()
	a.CreatedOn = time.Now().UTC()
	a.Password = string(hashedPassword)
	a.Active = false

	if err := self.Session.DB(self.DB).C(self.Coll).Insert(a); err != nil {
		return err
	}

	return nil
}
