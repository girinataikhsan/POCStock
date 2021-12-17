package web

import (
	"net/http"
	"stock/app/dto"
	"stock/app/usecase"
	stockUseCase "stock/app/usecase/stock"
	"stock/pkg/configuration/redis"
	"stock/pkg/response"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	routesData     Routes
	rc 				*redis.Redis
	usecase 		usecase.IStockUseCase
}

func NewHTTP(routes Routes,rc *redis.Redis) *HTTP {
	return &HTTP{
		routesData: routes,
		rc: rc,
		usecase:stockUseCase.NewStockUseCase(rc.RC),
	}
}

func (h *HTTP) PingHandler(ctx echo.Context) (e error) {

	ping := dto.Ping{
		Version: h.routesData.Version,
		Name:    h.routesData.AppName,
	}
	e = response.Show(ctx, http.StatusOK, http.StatusOK, "success", ping)
	return
}

func (h *HTTP) InitStockHandler(ctx echo.Context) (e error) {
	h.usecase.InitStock(ctx.Request().Context())
	e = response.Show(ctx, http.StatusOK, http.StatusOK, "success",nil)
	return
}

func (h *HTTP) OrderHandler(ctx echo.Context) (e error) {
	input := dto.Order{}
	if e = ctx.Bind(&input); e != nil {
		e = response.Show(ctx, http.StatusBadRequest, http.StatusBadRequest, "invalid parameters", e.Error())
		return
	}
	e = h.usecase.Order(ctx.Request().Context(), input)
	if e!=nil{
		e = response.Show(ctx, http.StatusBadRequest, http.StatusBadRequest, "failed order", e.Error())
		return
	}
	e = response.Show(ctx, http.StatusOK, http.StatusOK, "success",nil)
	return
}