package socialsso

import "context"

type Repository interface {
	// GetOrCreate creates a new account. If account already
	// exists it should simply return the existing account.
	GetOrCreate(context.Context, *Account) (*Account, error)
}
