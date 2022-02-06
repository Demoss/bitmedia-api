package service

import (
	"context"
	"fmt"
	api "github.com/bitmedia-api"
	"github.com/bitmedia-api/pkg/repository"
	"github.com/pkg/errors"
	"sync"
)

type EthService struct {
	repo repository.Eth
}

func NewEthService(repo repository.Eth) *EthService {
	return &EthService{repo: repo}
}

func (s *EthService) findBlockNumbersOnce(ctx context.Context) (numbers []string, err error) {
	var once sync.Once
	onceValue := func() {
		blockNumbers, err := s.repo.GetBlockNumbers(ctx)
		if err != nil {
			return
		}

		for _, number := range blockNumbers {
			numbers = append(numbers, fmt.Sprintf("%v", number))
		}
		return

	}
	once.Do(onceValue)
	return numbers, err
}

func (s *EthService) SaveTransactionsByBlock(ctx context.Context, blockByNumber api.BlockByNumber) error {

	once, err := s.findBlockNumbersOnce(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to find numbers")
	}

	for i := range once {
		if len(blockByNumber.Result.Transactions) == 0 || once[i] == blockByNumber.Result.Transactions[0].BlockNumber {
			return nil
		}
	}

	var transactions []api.Transaction

	for _, transaction := range blockByNumber.Result.Transactions {
		tr := api.Transaction{
			BaseFeePerGas: blockByNumber.Result.BaseFeePerGas,
			Timestamp:     blockByNumber.Result.Timestamp,
			BlockNumber:   transaction.BlockNumber,
			From:          transaction.From,
			To:            transaction.To,
			GasPrice:      transaction.GasPrice,
			Hash:          transaction.Hash,
		}
		transactions = append(transactions, tr)
	}

	return s.repo.SaveTransactionByBlock(ctx, transactions)
}

func (s *EthService) GetTransactionsByHash(ctx context.Context, hash string) (api.Transaction, error) {
	return s.repo.GetTransactionsByHash(ctx, hash)
}

func (s *EthService) GetTransactionsByUserFrom(ctx context.Context, hashUserFrom string, page int64) (t []api.Transaction, err error) {
	return s.repo.GetTransactionsByUserFrom(ctx, hashUserFrom, page)
}

func (s *EthService) GetTransactionsByBlock(ctx context.Context, tag string, page int64) (t []api.Transaction, err error) {
	return s.repo.GetTransactionsByBlock(ctx, tag, page)
}

func (s *EthService) GetTransactionsByUserTo(ctx context.Context, hashUserFrom string, page int64) (t []api.Transaction, err error) {
	return s.repo.GetTransactionsByUserTo(ctx, hashUserFrom, page)
}

func (s *EthService) GetTransactionsByTimestamp(ctx context.Context, timestamp string, page int64) (t []api.Transaction, err error) {
	return s.repo.GetTransactionsByTimestamp(ctx, timestamp, page)
}
