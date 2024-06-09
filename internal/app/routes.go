package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) registerRoutes() {
	a.echo.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	// authentication endpoint
	auth := a.echo.Group("auth")
	auth.POST("/register", a.userDelivery.Register)
	auth.POST("/login", a.userDelivery.Login)

	// user endpoint
	user := a.echo.Group("user")
	user.Use(a.middleware.JWTAuthentication)
	user.GET("/me", a.userDelivery.Me)
	user.POST("/buy-package", a.userDelivery.BuyPackage, a.middleware.JWTAuthentication)

	// profile endpoint
	profile := a.echo.Group("profile")
	profile.Use(a.middleware.JWTAuthentication)
	profile.POST("/update", a.profileDelivery.Update)

	// swipe endpoint
	swipe := a.echo.Group("swipe")
	swipe.Use(a.middleware.JWTAuthentication)
	swipe.GET("/next-profile", a.swipeDelivery.NextProfile)
	swipe.POST("/like", a.swipeDelivery.Like)

	// like endpoint
	like := a.echo.Group("like")
	like.Use(a.middleware.JWTAuthentication)
	like.GET("/mutual", a.likeDelivery.MutualList)
}
