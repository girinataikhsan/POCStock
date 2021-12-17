package repo

import (
	"context"
	"stock/app/dto"
)

type IStockRepo interface {
	InitStockRepo(ctx context.Context)()
	OrderRepo(ctx context.Context, order dto.Order)(err error)
}

