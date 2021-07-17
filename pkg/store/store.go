package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	db             *pgx.Conn
	linkRepository *LinkRepository
}

func (s *Store) Migrate() error {
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("Database url not found")
	}

	m, err := migrate.New(
		"shorter-api/migrations/20210717194858_create_links.up.sql",
		dbUrl)
	if err != nil {
		fmt.Println("Error 1:", err)
		return err
	}
	if err := m.Up(); err != nil {
		fmt.Println("Error 2:", err)
		return err
	}

	return nil
}

func (s *Store) Open() (*pgx.Conn, error) {
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("Database url not found")
	}
	db, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Println("con faled")
		log.Fatal(err)
	}

	s.db = db

	return s.db, nil
}

func (s *Store) Close() {
	s.db.Close(context.Background())
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
