package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"stock/app/handlers/web"
	"stock/pkg/configuration"
	"stock/pkg/configuration/redis"
	"time"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})
}

func configAndStartServer() {
	configMain := configuration.ServicesApp{
		EnvVariable: "ST",
		Path:        setPath(),
	}

	configMain.Load()
	rc := redis.New(configMain, context.Background())
	rc.InitRedis()

	htmlEcho := setWebRouter(configMain, rc)

	start(configMain, htmlEcho)
}

func setWebRouter(config configuration.ServicesApp, client *redis.Redis) *echo.Echo {
	e := echo.New()

	// Middleware Logger
	e.Use(middleware.RequestID())
	e.Use(middlewareLogging)

	// Register Routes
	web.NewRoutes(config).RegisterServices(e, client)

	return e
}

func start(config configuration.ServicesApp, htmlEcho *echo.Echo) {
	var err error

	if err = htmlEcho.Start(config.Config.ListenPort); err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}

func setPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	basePath := filepath.Dir(d)
	return basePath
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return log.WithFields(log.Fields{
		"at":        time.Now().Format("2006-01-02 15:04:05"),
		"method":    c.Request().Method,
		"uri":       c.Request().URL.String(),
		"ip":        c.Request().RemoteAddr,
		"requestID": c.Response().Header().Get(echo.HeaderXRequestID),
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("logger", makeLogEntry(c))
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func main() {
	initLogger()
	configAndStartServer()
}

