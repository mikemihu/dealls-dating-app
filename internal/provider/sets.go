package provider

import (
	"dealls-dating-app/internal/delivery"
	"dealls-dating-app/internal/repository"
	"dealls-dating-app/internal/usecase"

	"github.com/google/wire"

	"dealls-dating-app/internal/app"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/delivery/middleware"
)

var BaseSet = wire.NewSet(
	config.Get,
	LoggerProvider,
	DatabaseProvider,
	RedisProvider,
	JWTConfigProvider,
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepo,
	repository.NewProfileRepo,
	repository.NewPackageRepo,
	repository.NewLikeRepo,
	repository.NewSwipeRepo,
)

var UseCaseSet = wire.NewSet(
	usecase.NewUserUC,
	usecase.NewProfileUC,
	usecase.NewSwipeUC,
	usecase.NewLikeUC,
)

var DeliverySet = wire.NewSet(
	middleware.NewMiddleware,
	delivery.NewUserDelivery,
	delivery.NewProfileDelivery,
	delivery.NewSwipeDelivery,
	delivery.NewLikeDelivery,
)

var AppSet = wire.NewSet(
	BaseSet,
	RepositorySet,
	UseCaseSet,
	DeliverySet,
	app.New,
)
