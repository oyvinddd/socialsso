package account

import (
	"context"
)

type Service interface {
	SignIn(context.Context, string) (*Account, error)
}
