package internal

import "github.com/labstack/echo/v4"

type UserDelivery interface {
	// Register creates new user
	Register(c echo.Context) error
	// Login returns user token
	Login(c echo.Context) error
	// Me gets authenticated user info
	Me(c echo.Context) error
	// BuyPackage activates package
	BuyPackage(c echo.Context) error
}
