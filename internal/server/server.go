package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	maxHeaderBytes          = 1 << 20
	readTimeout             = 10 * time.Second
	writeTimeout            = 10 * time.Second
	timeForGracefulShutdown = 10 * time.Second
)

type Server struct {
	http.Server
}

func New(handler http.Handler, config Config) *Server {
	return &Server{
		http.Server{
			Addr:           config.Address + ":" + config.Port,
			Handler:        handler,
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
		},
	}
}

func WaitForInterrupt(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	interuptChan := make(chan os.Signal, 1)
	signal.Notify(interuptChan, os.Interrupt)

	go func() {
		<-interuptChan
		log.Println("interruption")
		cancel()
	}()

	return ctx
}

func (s *Server) Run(ctx context.Context) error {
	ctx = WaitForInterrupt(ctx)

	serverErrChan := make(chan error)

	go func() {
		log.Println("start server")
		err := s.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			serverErrChan <- err
		}
	}()

	select {
	case err := <-serverErrChan:
		log.Printf("server error: %v\n", err)
		return err

	case <-ctx.Done():
		ctxShutDown, cancel := context.WithTimeout(context.Background(), timeForGracefulShutdown)

		defer cancel()

		if err := s.Shutdown(ctxShutDown); err != nil {
			log.Println("server didn't exit properly")
			return err
		}

		log.Println("server exited properly")
	}

	return nil
}

type Config struct {
	Address string
	Port    string
}
