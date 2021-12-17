package stockRepo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"stock/app/dto"
	"stock/app/repo"
)

type stock struct {
	rc *redis.Client
}

func NewStockRepo(rc *redis.Client) repo.IStockRepo {
	return &stock{
		rc: rc,
	}
}

func (s stock)InitStockRepo(ctx context.Context)(){
	s.rc.Set(ctx, "PROD001",100,0)
	s.rc.Set(ctx, "PROD002",100,0)
	s.rc.Set(ctx, "PROD003",100,0)
}

func (s stock) OrderRepo(ctx context.Context, order dto.Order)(err error){
	exist := s.rc.Exists(ctx, order.ProductID).Val()
	if exist==0{
		err = fmt.Errorf("invalid productID")
		return
	}
	limit, err := s.rc.Get(ctx, order.ProductID).Int64()
	if limit<=0{
		err = fmt.Errorf("Stock already empty")
		return
	}
	limit = s.rc.DecrBy(ctx, order.ProductID, int64(order.Amount)).Val()
	if limit<0{
		err = fmt.Errorf("Stock already empty")
		log.Info("return stock because insufficient stock")
		s.rc.IncrBy(ctx, order.ProductID, int64(order.Amount)).Val()
	}
	return
}
