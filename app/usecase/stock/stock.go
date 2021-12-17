package stockUseCase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"stock/app/dto"
	"stock/app/repo"
	stockRepo "stock/app/repo/stock"
	"stock/app/usecase"
)

type stock struct {
	rc *redis.Client
	repo repo.IStockRepo
}

func NewStockUseCase(rc *redis.Client)(usecase.IStockUseCase){
	return &stock{
		rc: rc,
		repo: stockRepo.NewStockRepo(rc),
	}
}

func (s stock)InitStock(ctx context.Context)(){
	s.repo.InitStockRepo(ctx)
}

func (s stock) Order(ctx context.Context, order dto.Order)(err error){
	err = s.repo.OrderRepo(ctx, order)
	return
}