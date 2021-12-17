package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"stock/pkg/configuration"
	"stock/pkg/configuration/redis"
	"time"
)

type Routes struct {
	Env                          string
	RootPath                     string
	URIProxy                     string
	AppName                      string
	Version                      string
	SessionURL                   string
	EncryptionKey                string
	TYkKeyID                     string
	TYKKeySecret                 string
	MccmUrlHost                  string
	TouchPoint                   string
	XAppID                       string
	XAppSecret                   string
	JwtAuth                      string
	JwtSignature                 string
	JwtMinExp                    string
	AerospikeNamespace           string
	AerospikeVoucherGroupNameset string
	AerospikePricePlanNameset    string
	AssetUrl                     string
}

func NewRoutes(config configuration.ServicesApp) *Routes {
	return &Routes{
		RootPath:                     config.Config.RootURL,
		AppName:                      config.Config.AppName,
		Version:                      config.Config.Version,
	}
}

func (route *Routes) RegisterServices(r *echo.Echo, rc *redis.Redis) {
	handler := NewHTTP(*route, rc)
	apiRoutes := r.Group(route.RootPath)

	route.setMiddleware(apiRoutes)

	// Routes Endpoint
	apiRoutes.GET("/ping", handler.PingHandler)
	apiRoutes.POST("/init-stock", handler.InitStockHandler)
	apiRoutes.POST("/order", handler.OrderHandler)
}

func (route *Routes) setMiddleware(rGroup *echo.Group) {
	rGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRealIP},
		AllowMethods: []string{http.MethodGet, http.MethodPut},
	}))

	rGroup.Use(middleware.BodyLimit("2M"))

	rGroup.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second,
	}))
}
