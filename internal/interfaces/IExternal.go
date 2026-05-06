package interfaces

import (
	"context"
	"ewallet-fastcampus/external"
)

type IWallet interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
}
