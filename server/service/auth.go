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
	credentials := Credentials{}
	if binding.Bind(r, &credentials).Handle(w) {
		return
	}
	role, err := self.Store.CheckCredentials(credentials.Email, credentials.Password)
	if err != nil {
		renderError(w, r, err)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims["sub"] = credentials.Email
	token.Claims["sub"] = struct {
		Role  string `json:"role"`
		Email string `json:"email"`
	}{role, credentials.Email}
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
