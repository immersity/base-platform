package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"
	flag "github.com/ianschenck/envflag"
	"github.com/immersity/base-platform/server/service"
	"github.com/immersity/base-platform/server/store"
)

func main() {
	var (
		httpAddr  = flag.String("HTTP_ADDRESS", "0.0.0.0:3000", "http address")
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

	s := service.New(st, service.Config{
		JwtSecret: *jwtSecret,
		JwtExpiry: *jwtExpiry,
	})

	errChan := make(chan error, 1)
	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = s.Handler()

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	log.Printf("HTTP service listening on %s...\n", *httpAddr)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case sig := <-signalChan:
			log.Printf("Captured %v. Exiting...\n", sig)
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
