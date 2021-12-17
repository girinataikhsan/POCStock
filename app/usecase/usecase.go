package usecase

import (
	"context"
	"stock/app/dto"
)

type IStockUseCase interface {
	InitStock(ctx context.Context)()
	Order(ctx context.Context, order dto.Order)(err error)
}
