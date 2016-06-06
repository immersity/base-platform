package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	flag "github.com/ianschenck/envflag"
	"github.com/immersity/base-platform/server/service"
	"github.com/immersity/base-platform/server/store/mongo"
)

func main() {
	var (
		port      = flag.Int("PORT", 3000, "port to listen on")
		mongoDb   = flag.String("MONGO_DB", "immersity", "mongo db")
		mongoDsn  = flag.String("MONGO_DSN", "127.0.0.1", "mongo dsn")
		jwtSecret = flag.String("JWT_SECRET", "olakease", "secret to sign tokens with")
		jwtExpiry = flag.Duration("JWT_EXPIRY", time.Hour*1, "token expiry")
	)
	flag.Parse()

	mongoStore, err := mongo.New(mongo.Config{
		DB:       *mongoDb,
		MongoDsn: *mongoDsn,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer mongoStore.Close()

	addr := fmt.Sprintf("0.0.0.0:%d", *port)

	srvc := service.New(mongoStore, service.Config{
		JwtSecret: *jwtSecret,
		JwtExpiry: *jwtExpiry,
	})

	log.Fatal(http.ListenAndServe(addr, srvc.Handler()))
}
