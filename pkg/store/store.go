package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	linkRepository *LinkRepository
}

func (s *Store) Open() error {
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("Database url not found")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Link() *LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}

	s.linkRepository = &LinkRepository{
		store: s,
	}

	return s.linkRepository

}
