package apiserver

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
}

func NewServer(router *http.ServeMux) *APIServer {
	return &APIServer{
		config: NewConfigMust(),
		logger: logrus.New(),
		router: router,
	}
}

func (s *APIServer) Run() error {
	if err := s.configureLogger(); err != nil {
		return fmt.Errorf("configure logger: %w", err)
	}

	s.logger.Info("starting API server")

	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		s.logger.Error(err)
	}

	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return fmt.Errorf("parse log level: %w", err)
	}

	s.logger.SetLevel(level)

	return nil
}
