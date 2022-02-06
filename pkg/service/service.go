package service

import (
	"context"
	api "github.com/bitmedia-api"
	"github.com/bitmedia-api/pkg/repository"
)

type Eth interface {
	SaveTransactionsByBlock(ctx context.Context, blockByNumber api.BlockByNumber) error
	GetTransactionsByHash(ctx context.Context, hash string) (api.Transaction, error)
	GetTransactionsByUserFrom(ctx context.Context, hashUserFrom string, page int64) (t []api.Transaction, err error)
	GetTransactionsByBlock(ctx context.Context, tag string, page int64) (t []api.Transaction, err error)
	GetTransactionsByUserTo(ctx context.Context, hashUserFrom string, page int64) (t []api.Transaction, err error)
	GetTransactionsByTimestamp(ctx context.Context, timestamp string, page int64) (t []api.Transaction, err error)
}
type Service struct {
	Eth
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Eth: NewEthService(repository)}
}
