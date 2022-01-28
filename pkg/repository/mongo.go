package repository

import (
	"context"
	api "github.com/bitmedia-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const BlocksByNumber = "BlocksByNumber"

type EthMongo struct {
	*mongo.Database
}

func NewEthMongo(db *mongo.Database) *EthMongo {
	return &EthMongo{Database: db}
}

func (r *EthMongo) SaveBlockByNumber(ctx context.Context, blockByNumber api.BlockByNumber) error {
	_, err := r.Database.Collection(BlocksByNumber).InsertOne(ctx, blockByNumber)
	if err != nil {
		return err
	}

	return nil
}

func (r *EthMongo) GetTransactionByHash(ctx context.Context, hash string) (result api.Result, err error) {
	filter := bson.M{"result.transactions.hash": hash}

	var block api.BlockByNumber

	err = r.Database.Collection(BlocksByNumber).FindOne(ctx, filter).Decode(&block)

	for _, transaction := range block.Result.Transactions {
		if transaction.Hash == hash {
			result = block.Result
			break
		}
	}

	return result, nil
}

func (r *EthMongo) GetTransactionByUserFrom(ctx context.Context, hashUserFrom string) (result []api.Result, err error) {
	filter := bson.M{"result.transactions.from": hashUserFrom}

	cur, err := r.Database.Collection(BlocksByNumber).Find(ctx, filter)
	if cur.Err() != nil {
		return nil, err
	}

	var blocks []api.BlockByNumber

	var transactions []api.Transactions

	if err = cur.All(ctx, &blocks); err != nil {
		return result, err
	}
	for i := range blocks {
		for j := range blocks[i].Result.Transactions {
			if blocks[i].Result.Transactions[j].From == hashUserFrom {
				transactions = append(transactions, blocks[i].Result.Transactions[j])
			}
		}
		block := api.BlockByNumber{
			Result: api.Result{
				BaseFeePerGas: blocks[i].Result.BaseFeePerGas,
				Timestamp:     blocks[i].Result.Timestamp,
				Transactions:  transactions,
			},
		}
		result = append(result, block.Result)
	}

	return result, nil
}

func (r *EthMongo) GetTransactionByBlock(ctx context.Context, tag string) (t api.Result, err error) {
	filter := bson.M{"result.transactions.blockNumber": tag}

	var block api.BlockByNumber

	err = r.Database.Collection(BlocksByNumber).FindOne(ctx, filter).Decode(&block)

	return block.Result, err
}

func (r *EthMongo) GetTransactionByUserTo(ctx context.Context, hashUserTo string) (result []api.Result, err error) {
	filter := bson.M{"result.transactions.to": hashUserTo}

	cur, err := r.Database.Collection(BlocksByNumber).Find(ctx, filter)
	if cur.Err() != nil {
		return nil, err
	}

	var blocks []api.BlockByNumber

	var transactions []api.Transactions
	if err = cur.All(ctx, &blocks); err != nil {
		return result, err
	}

	for i := range blocks {
		for j := range blocks[i].Result.Transactions {
			if blocks[i].Result.Transactions[j].To == hashUserTo {
				transactions = append(transactions, blocks[i].Result.Transactions[j])
			}
		}
		block := api.BlockByNumber{
			Result: api.Result{
				BaseFeePerGas: blocks[i].Result.BaseFeePerGas,
				Timestamp:     blocks[i].Result.Timestamp,
				Transactions:  transactions,
			},
		}
		result = append(result, block.Result)
	}
	return result, nil
}

func (r *EthMongo) GetTransactionByTimestamp(ctx context.Context, timestamp string) (result []api.Result, err error) {
	filter := bson.M{"result.timestamp": timestamp}

	cur, err := r.Database.Collection(BlocksByNumber).Find(ctx, filter)
	if cur.Err() != nil {
		return nil, err
	}

	var blocks []api.BlockByNumber

	var transactions []api.Transactions

	if err = cur.All(ctx, &blocks); err != nil {
		return result, err
	}
	for i := range blocks {
		if blocks[i].Result.Timestamp == timestamp {
			transactions = blocks[i].Result.Transactions
		}
		block := api.BlockByNumber{
			Result: api.Result{
				BaseFeePerGas: blocks[i].Result.BaseFeePerGas,
				Timestamp:     blocks[i].Result.Timestamp,
				Transactions:  transactions,
			},
		}
		result = append(result, block.Result)
	}

	return result, nil
}
