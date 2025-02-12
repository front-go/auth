//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package service

import "context"

type DbRepo interface {
	Insert(ctx context.Context, username, password string) error
	GetPassword(ctx context.Context, username string) (string, error)
	UpdatePassword(ctx context.Context, username, password, new_password string) (string, error)
}
