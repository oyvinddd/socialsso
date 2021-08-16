package provider

import (
	"context"
	"socialsso/account"
)

type Service interface {
	SignIn(context.Context, string) (*account.Account, error)
}
