package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"sync"
	"time"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/delivery/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type App struct {
	cfg             *config.Cfg
	logger          *zap.Logger
	middleware      *middleware.Middleware
	echo            *echo.Echo
	jwtConfig       echojwt.Config
	userDelivery    internal.UserDelivery
	profileDelivery internal.ProfileDelivery
	swipeDelivery   internal.SwipeDelivery
	likeDelivery    internal.LikeDelivery
}

func New(
	cfg *config.Cfg,
	logger *zap.Logger,
	jwtConfig echojwt.Config,
	middleware *middleware.Middleware,
	userDelivery internal.UserDelivery,
	profileDelivery internal.ProfileDelivery,
	swipeDelivery internal.SwipeDelivery,
	likeDelivery internal.LikeDelivery,
) *App {
	return &App{
		cfg:             cfg,
		logger:          logger,
		echo:            echo.New(),
		jwtConfig:       jwtConfig,
		middleware:      middleware,
		userDelivery:    userDelivery,
		profileDelivery: profileDelivery,
		swipeDelivery:   swipeDelivery,
		likeDelivery:    likeDelivery,
	}
}

const (
	_ShutdownTimeout = 5 * time.Second
)

func (a *App) Run() error {
	a.logger.Info("starting app")

	// global middleware
	a.echo.Use(echoMiddleware.Logger())
	a.echo.Use(echoMiddleware.RateLimiter(echoMiddleware.NewRateLimiterMemoryStore(rate.Limit(a.cfg.App.RateLimit))))
	if os.Getenv("ENV") == constant.EnvProduction {
		a.echo.Use(echoMiddleware.Recover())
	}

	// register route
	a.registerRoutes()

	// start echo server
	err := a.echo.Start(a.cfg.App.Address)
	if errors.Is(err, http.ErrServerClosed) {
		// ignore shutdown call error
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// Stop gracefully stop the app
func (a *App) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), _ShutdownTimeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := a.echo.Shutdown(ctx)
		if err != nil {
			a.logger.Fatal("error echo shutdown", zap.Error(err))
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}
