package repository

import (
	"context"
	"fmt"
)

type Repository struct {
	storage map[string]string
}

func NewRepository() *Repository {
	st := make(map[string]string)
	return &Repository{storage: st}
}

func (r *Repository) Insert(ctx context.Context, username, password string) error {
	if _, ok := r.storage[username]; ok {
		return fmt.Errorf("username %s already exists", username)
	}
	r.storage[username] = password
	return nil
}
