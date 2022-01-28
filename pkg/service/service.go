package service

import (
	"context"
	api "github.com/bitmedia-api"
	"github.com/bitmedia-api/pkg/repository"
)

type Eth interface {
	SaveBlockByNumber(ctx context.Context, blockByNumber api.BlockByNumber) error
	GetTransactionByHash(ctx context.Context, hash string) (api.Result, error)
	GetTransactionByUserFrom(ctx context.Context, hashUserFrom string) (t []api.Result, err error)
	GetTransactionByBlock(ctx context.Context, tag string) (t api.Result, err error)
	GetTransactionByUserTo(ctx context.Context, hashUserFrom string) (t []api.Result, err error)
	GetTransactionByTimestamp(ctx context.Context, timestamp string) (t []api.Result, err error)
}
type Service struct {
	Eth
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Eth: NewEthService(repository)}
}
