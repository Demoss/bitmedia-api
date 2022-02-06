package repository

import (
	"context"
	api "github.com/bitmedia-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Eth interface {
	SaveTransactionByBlock(ctx context.Context, transactions []api.Transaction) error
	GetTransactionsByHash(ctx context.Context, hash string) (api.Transaction, error)
	GetTransactionsByUserFrom(ctx context.Context, hashUserFrom string, page int64) (result []api.Transaction, err error)
	GetTransactionsByBlock(ctx context.Context, tag string, page int64) (result []api.Transaction, err error)
	GetTransactionsByUserTo(ctx context.Context, hashUserFrom string, page int64) (result []api.Transaction, err error)
	GetTransactionsByTimestamp(ctx context.Context, timestamp string, page int64) (result []api.Transaction, err error)
	GetBlockNumbers(ctx context.Context) (b []interface{}, err error)
}

type Repository struct {
	Eth
}

func NewRepository(database *mongo.Database) *Repository {
	return &Repository{Eth: NewEthMongo(database)}
}
