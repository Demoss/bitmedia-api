package service

import (
	"context"
	api "github.com/bitmedia-api"
	"github.com/bitmedia-api/pkg/repository"
)

type EthService struct {
	repo repository.Eth
}

func NewEthService(repo repository.Eth) *EthService {
	return &EthService{repo: repo}
}

func (s *EthService) SaveBlockByNumber(ctx context.Context, blockByNumber api.BlockByNumber) error {

	return s.repo.SaveBlockByNumber(ctx, blockByNumber)
}

func (s *EthService) GetTransactionByHash(ctx context.Context, hash string) (api.Result, error) {
	return s.repo.GetTransactionByHash(ctx, hash)
}

func (s *EthService) GetTransactionByUserFrom(ctx context.Context, hashUserFrom string) (t []api.Result, err error) {
	return s.repo.GetTransactionByUserFrom(ctx, hashUserFrom)
}

func (s *EthService) GetTransactionByBlock(ctx context.Context, tag string) (t api.Result, err error) {
	return s.repo.GetTransactionByBlock(ctx, tag)
}

func (s *EthService) GetTransactionByUserTo(ctx context.Context, hashUserFrom string) (t []api.Result, err error) {
	return s.repo.GetTransactionByUserTo(ctx, hashUserFrom)
}

func (s *EthService) GetTransactionByTimestamp(ctx context.Context, timestamp string) (t []api.Result, err error) {
	return s.repo.GetTransactionByTimestamp(ctx, timestamp)
}
