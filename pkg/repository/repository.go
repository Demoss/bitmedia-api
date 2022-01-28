package repository

import (
	"context"
	api "github.com/bitmedia-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Eth interface {
	SaveBlockByNumber(ctx context.Context, blockByNumber api.BlockByNumber) error
	GetTransactionByHash(ctx context.Context, hash string) (api.Result, error)
	GetTransactionByUserFrom(ctx context.Context, hashUserFrom string) (result []api.Result, err error)
	GetTransactionByBlock(ctx context.Context, tag string) (result api.Result, err error)
	GetTransactionByUserTo(ctx context.Context, hashUserFrom string) (result []api.Result, err error)
	GetTransactionByTimestamp(ctx context.Context, timestamp string) (result []api.Result, err error)
}

type Repository struct {
	Eth
}

func NewRepository(database *mongo.Database) *Repository {
	return &Repository{Eth: NewEthMongo(database)}
}
