package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func (r *Repository) GetPassword(ctx context.Context, username string) (string, error) {
	query := `SELECT password FROM participant WHERE username = $1`
	var password string
	err := r.conn.QueryRowContext(ctx, query, username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("username %s not found", username)
		}
		return "", fmt.Errorf("failed to get password for username %s: %w", username, err)
	}
	return password, nil
}

func (r *Repository) UpdatePassword(ctx context.Context, username, password, new_password string) (string, error) {
	query := `UPDATE participant SET password =$3  WHERE username = $1 AND password=$2 `
	_, err := r.conn.ExecContext(ctx, query, username, password, new_password)
	if err != nil {
		return "", fmt.Errorf("Не удалось изменить пароль")
	}
	return "Пароль успешно изменен", nil
}
