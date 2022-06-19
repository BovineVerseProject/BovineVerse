package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"brave-ox-web/config"
	"brave-ox-web/log"
	"brave-ox-web/web/routing"
)

type server struct {
	s *http.Server
}

func newServer() *server {
	return &server{
		s: &http.Server{
			Handler: routing.InitEngine(),
			Addr:    fmt.Sprintf(":%d", config.Port()),
		},
	}
}

func (s *server) start() {
	log.Infof("server listen at:%s", s.s.Addr)
	if err := s.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error(err)
	}
}

func (s *server) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.s.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown:%v", err)
	}
}
