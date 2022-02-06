package repository

import (
	"context"
	"fmt"
	api "github.com/bitmedia-api"
	_ "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Transaction = "Transaction"
const perPage int64 = 9

type EthMongo struct {
	*mongo.Database
}

func NewEthMongo(db *mongo.Database) *EthMongo {
	return &EthMongo{Database: db}
}

func (r *EthMongo) SaveTransactionByBlock(ctx context.Context, transactions []api.Transaction) error {

	if transactions == nil {
		return nil
	}

	var res []interface{}

	for _, tr := range transactions {
		res = append(res, tr)
	}

	_, err := r.Database.Collection(Transaction).InsertMany(ctx, res)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *EthMongo) GetTransactionsByHash(ctx context.Context, hash string) (api.Transaction, error) {
	filter := bson.M{"hash": hash}

	var transaction api.Transaction

	err := r.Database.Collection(Transaction).FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		return api.Transaction{}, err
	}

	return transaction, nil
}

func (r *EthMongo) GetTransactionsByUserFrom(ctx context.Context, hashUserFrom string, page int64) (transactions []api.Transaction, err error) {
	filter := bson.M{"from": hashUserFrom}

	findOptions := pagination(page)

	cur, err := r.Database.Collection(Transaction).Find(ctx, filter, findOptions)
	if cur.Err() != nil {
		return nil, err
	}

	if err = cur.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *EthMongo) GetTransactionsByBlock(ctx context.Context, tag string, page int64) (transactions []api.Transaction, err error) {
	filter := bson.M{"blockNumber": tag}

	findOptions := pagination(page)

	cur, err := r.Database.Collection(Transaction).Find(ctx, filter, findOptions)
	if cur.Err() != nil {
		return nil, err
	}

	if err = cur.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *EthMongo) GetTransactionsByUserTo(ctx context.Context, hashUserTo string, page int64) (transactions []api.Transaction, err error) {
	filter := bson.M{"to": hashUserTo}

	findOptions := pagination(page)

	cur, err := r.Database.Collection(Transaction).Find(ctx, filter, findOptions)
	if cur.Err() != nil {
		return nil, err
	}

	if err = cur.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *EthMongo) GetTransactionsByTimestamp(ctx context.Context, timestamp string, page int64) (transactions []api.Transaction, err error) {
	filter := bson.M{"timestamp": timestamp}

	findOptions := pagination(page)

	cur, err := r.Database.Collection(Transaction).Find(ctx, filter, findOptions)
	if cur.Err() != nil {
		return nil, err
	}

	if err = cur.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *EthMongo) GetBlockNumbers(ctx context.Context) (b []interface{}, err error) {
	res, err := r.Database.Collection(Transaction).Distinct(ctx, "blockNumber", bson.D{})
	if err != nil {
		return nil, err
	}

	for _, result := range res {
		b = append(b, result)
	}

	return b, nil
}

func pagination(page int64) *options.FindOptions {
	var findOptions *options.FindOptions
	findOptions = options.Find()
	findOptions.SetLimit(perPage)
	findOptions.SetSkip((page - 1) * perPage)
	return findOptions
}
