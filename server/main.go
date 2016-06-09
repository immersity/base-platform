package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	flag "github.com/ianschenck/envflag"
	"github.com/immersity/base-platform/server/service"
	"github.com/immersity/base-platform/server/store"
)

func main() {
	var (
		port      = flag.Int("PORT", 3000, "port")
		mysqlDsn  = flag.String("MYSQL_DSN", "crowl:crowl@/immersity", "mysql dsn")
		jwtSecret = flag.String("JWT_SECRET", "olakease", "jwt secret")
		jwtExpiry = flag.Duration("JWT_EXPIRY", time.Hour*1, "jwt expiry")
	)
	flag.Parse()
	st, err := store.New(store.Config{
		Dsn: *mysqlDsn,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
	srvc := service.New(st, service.Config{
		JwtSecret: *jwtSecret,
		JwtExpiry: *jwtExpiry,
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), srvc.Handler()))
}
