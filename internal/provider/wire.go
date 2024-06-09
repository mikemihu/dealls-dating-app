//go:build wireinject
// +build wireinject

package provider

import (
	"dealls-dating-app/internal/app"

	"github.com/google/wire"
)

func ProvideApp() *app.App {
	wire.Build(AppSet)

	// Return any struct that exist inside the build
	return &app.App{}
}
