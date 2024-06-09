package delivery

import (
	"net/http"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/internal/pkg/contexts"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	cfg    *config.Cfg
	userUC internal.UserUC
}

func NewUserDelivery(
	cfg *config.Cfg,
	userUC internal.UserUC,
) internal.UserDelivery {
	return &UserDelivery{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (u *UserDelivery) Register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email or password is empty")
	}

	err := u.userUC.Register(c.Request().Context(), email, password)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (u *UserDelivery) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "email or password is empty")
	}

	token, err := u.userUC.Login(c.Request().Context(), email, password)
	if err != nil {
		return handlerErrorResponse(err)
	}

	resp := entity.AuthResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *UserDelivery) Me(c echo.Context) error {
	user := contexts.GetUser(c)
	return c.JSON(http.StatusOK, user)
}

func (u *UserDelivery) BuyPackage(c echo.Context) error {
	req := entity.BuyPackageRequest{}
	err := c.Bind(&req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	req.UserID = contexts.GetUser(c).ID

	err = u.userUC.BuyPackage(c.Request().Context(), req)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, nil)
}
