package provider

import (
	"fmt"
	"sync"

	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/pkg/authentication"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once

	db     *gorm.DB
	dbOnce sync.Once

	redisClient     *redis.Client
	redisClientOnce sync.Once

	jwtConfig     echojwt.Config
	jwtConfigOnce sync.Once
)

func LoggerProvider() *zap.Logger {
	loggerOnce.Do(func() {
		logger, _ = zap.NewDevelopment()
		_ = zap.ReplaceGlobals(logger)
	})
	return logger
}

func DatabaseProvider(cfg *config.Cfg) *gorm.DB {
	dbOnce.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			cfg.Storage.Postgres.Host,
			cfg.Storage.Postgres.Port,
			cfg.Storage.Postgres.User,
			cfg.Storage.Postgres.Password,
			cfg.Storage.Postgres.Database,
			constant.DefaultTimeZone,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			TranslateError: true,
		})
		if err != nil {
			logger.Fatal("failed open database", zap.Error(err))
		}

		db = db.Debug()
	})
	return db
}

func RedisProvider(cfg *config.Cfg) *redis.Client {
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     cfg.Storage.Redis.Address,
			Password: cfg.Storage.Redis.Password,
		})
	})
	return redisClient
}

func JWTConfigProvider(cfg *config.Cfg) echojwt.Config {
	jwtConfigOnce.Do(func() {
		jwtConfig = echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(authentication.AuthClaims)
			},
			SigningKey: []byte(cfg.Auth.Jwt.Secret),
		}
	})
	return jwtConfig
}
