package account

import "context"

type Repository interface {
	GetOrCreate(ctx context.Context, email string) (*Account, error)
}
