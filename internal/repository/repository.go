package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	storage map[string]string
	conn    *sqlx.DB
}

func NewRepository() *Repository {
	st := make(map[string]string)

	connectCmd := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		"master", "master", "master", "localhost", "3125")

	conn, err := sqlx.Connect("postgres", connectCmd)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{storage: st, conn: conn}
}

func (r *Repository) Insert(ctx context.Context, username, password string) error {
	if _, ok := r.storage[username]; ok {
		return fmt.Errorf("username %s already exists", username)
	}
	r.storage[username] = password

	query := `INSERT INTO participant (username, password) VALUES ($1, $2)`
	_, err := r.conn.ExecContext(ctx, query, username, password)
	if err != nil {
		return fmt.Errorf("insert username %s failed: %w", username, err)
	}
	return nil
}
