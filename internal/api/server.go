package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pizza-app/internal/data"
	"pizza-app/internal/sms"
	"strings"
	"syscall"
	"time"
)

// NewApp is the Factory method for api server app
func NewApp(repo data.Repo, msg sms.Messenger) *App {
	return &App{
		Repo: repo,
		Msg:  msg,
	}
}

const (
	DefaultReadTimeout = 5 * time.Second

	DefaultWriteTimeout = 10 * time.Second
)

// graceful shutdown.
type Service interface {
	Shutdown(context.Context)
}

// Serve starts a service backed by an http.Server using default options.
func Serve(s Service, port string, h http.Handler) {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	server := http.Server{
		Addr:         port,
		Handler:      h,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
	}
	ServeWithHTTPServer(s, &server)
}

func ServeWithHTTPServer(s Service, hs *http.Server) {
	go func() {
		log.Printf("listening on port %s...\n", hs.Addr)
		err := hs.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

// Shutdown is called for graceful shutdown
func (a *App) Shutdown(c context.Context) {
	log.Printf("API server shut down gracefully.")
}
