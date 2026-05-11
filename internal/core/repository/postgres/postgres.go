package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Store struct {
	config *Config
	db     *sql.DB
	logger *logrus.Logger
}

func NewStore() *Store {
	return &Store{
		config: NewConfigMust(),
		logger: logrus.New(),
	}
}

func (s *Store) configureLogger() error {
	lvl, err := logrus.ParseLevel("DEBUG")
	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)

	return nil
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func (s *Store) Open() error {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.config.User,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.Database,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	logrus.Info("repository was open")

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
