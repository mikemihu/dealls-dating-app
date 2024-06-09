package delivery

import (
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/internal/pkg/contexts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SwipeDelivery struct {
	cfg     *config.Cfg
	swipeUC internal.SwipeUC
}

func NewSwipeDelivery(
	cfg *config.Cfg,
	swipeUC internal.SwipeUC,
) internal.SwipeDelivery {
	return &SwipeDelivery{
		cfg:     cfg,
		swipeUC: swipeUC,
	}
}

func (d *SwipeDelivery) NextProfile(c echo.Context) error {
	req := entity.NextProfileRequest{
		User: contexts.GetUser(c),
	}
	profiles, err := d.swipeUC.NextProfile(c.Request().Context(), req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, profiles)
}

func (d *SwipeDelivery) Like(c echo.Context) error {
	var req entity.SwipeRequest
	err := c.Bind(&req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	req.UserID = contexts.GetUser(c).ID

	profiles, err := d.swipeUC.Like(c.Request().Context(), req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, profiles)
}
