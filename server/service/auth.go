package service

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mholt/binding"
)

type AuthService struct {
	Store     AuthStore
	JwtSecret string
	JwtExpiry time.Duration
}

func NewAuthService(store AuthStore, secret string, expiry time.Duration) *AuthService {
	return &AuthService{
		store,
		secret,
		expiry,
	}
}

func (self *AuthService) CreateToken(w http.ResponseWriter, r *http.Request) {
	creds := credentials{}
	if binding.Bind(r, &creds).Handle(w) {
		return
	}
	if err := self.Store.CheckCredentials(creds.Email, creds.Password); err != nil {
		renderError(w, r, err)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["sub"] = creds.Email
	token.Claims["exp"] = time.Now().UTC().Unix() + int64(self.JwtExpiry.Seconds())
	tokenStr, err := token.SignedString([]byte(self.JwtSecret))
	if err != nil {
		renderError(w, r, err)
		return
	}
	render(w, http.StatusCreated, struct {
		Token string `json:"token"`
	}{tokenStr})
}
