package service

import (
	"net/http"
	"time"

	"github.com/goware/cors"
	"github.com/goware/jwtauth"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	//"golang.org/x/net/context"
)

type Config struct {
	JwtSecret string
	JwtExpiry time.Duration
}

type Service struct {
	config Config
	*AuthService
	*AccountService
}

func New(store Store, config Config) *Service {
	return &Service{
		config,
		NewAuthService(store, config.JwtSecret, config.JwtExpiry),
		NewAccountService(store),
	}
}

func (s *Service) Handler() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	JwtAuth := jwtauth.New("HS256", []byte(s.config.JwtSecret), nil)

	router.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/tokens", s.CreateToken)
		})

		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", s.CreateAccount)
		})

		r.Group(func(r chi.Router) {
			r.Use(JwtAuth.Verifier)
			r.Use(jwtauth.Authenticator)
		})
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: false,
	})

	return c.Handler(router)
}
