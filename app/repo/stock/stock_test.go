package stockRepo

import (
	"context"
	rc "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"path"
	"path/filepath"
	"runtime"
	"stock/app/dto"
	"stock/pkg/configuration"
	"stock/pkg/configuration/redis"
	"testing"
)
var redisClient *rc.Client
func initTest(){
	configMain := configuration.ServicesApp{
		EnvVariable: "ST",
		Path:        setPath(),
	}

	configMain.Load()
	initRed := redis.New(configMain, context.Background())
	initRed.InitRedis()
	redisClient =initRed.RC
}

func setPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	basePath := filepath.Dir(d)
	return basePath
}

func TestOrderMoreThanStock(t *testing.T){
	initTest()
	repo := NewStockRepo(redisClient)
	err := repo.OrderRepo(context.Background(), dto.Order{
		ProductID: "PROD001",
		Amount:    9990,
	})
	assert.Error(t, err)
}

func TestOrder(t *testing.T){
	repo := NewStockRepo(redisClient)
	err := repo.OrderRepo(context.Background(), dto.Order{
		ProductID: "PROD001",
		Amount:    10,
	})
	assert.NoError(t, err)
	repo.InitStockRepo(context.Background())
}

func TestRaceConditionPositive(t *testing.T){
	repo := NewStockRepo(redisClient)
	jobs := make(chan error)
	for i:=1; i<=5;i++{
		go func() {
			jobs <- repo.OrderRepo(context.Background(), dto.Order{
				ProductID: "PROD001",
				Amount:    10,
			})
		}()
	}
	arrayErr := []error{}
	for i:=1; i<=5;i++{
		errFound := <-jobs
		if errFound!=nil{
			arrayErr = append(arrayErr,errFound)
		}
	}
	assert.Len(t, arrayErr,0)
	repo.InitStockRepo(context.Background())
}

func TestRaceConditionNegative(t *testing.T){
	repo := NewStockRepo(redisClient)
	jobs := make(chan error)
	for i:=1; i<=5;i++{
		go func() {
			jobs <- repo.OrderRepo(context.Background(), dto.Order{
				ProductID: "PROD001",
				Amount:    33,
			})
		}()
	}
	arrayErr := []error{}
	for i:=1; i<=5;i++{
		errFound := <-jobs
		if errFound!=nil{
			arrayErr = append(arrayErr,errFound)
		}
	}
	//should be 2 request will blocked and return error
	assert.Len(t, arrayErr,2)
	repo.InitStockRepo(context.Background())
}