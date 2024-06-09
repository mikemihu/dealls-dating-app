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

type ProfileDelivery interface {
	// Update updates profile info
	Update(c echo.Context) error
}

type SwipeDelivery interface {
	// NextProfile returns next profile to swipe
	NextProfile(c echo.Context) error
	// Like marks target user as liked, returns matched status
	Like(c echo.Context) error
}

type LikeDelivery interface {
	// MutualList returns mutual list
	MutualList(c echo.Context) error
}
