package delivery

import (
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/internal/pkg/contexts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileDelivery struct {
	cfg       *config.Cfg
	profileUC internal.ProfileUC
}

func NewProfileDelivery(
	cfg *config.Cfg,
	profileUC internal.ProfileUC,
) internal.ProfileDelivery {
	return &ProfileDelivery{
		cfg:       cfg,
		profileUC: profileUC,
	}
}

func (d *ProfileDelivery) GetList(c echo.Context) error {
	profiles, err := d.profileUC.GetList(c.Request().Context(), entity.ProfileFilter{})
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, profiles)
}

func (d *ProfileDelivery) Update(c echo.Context) error {
	var req entity.ProfileUpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	req.UserID = contexts.GetUser(c).ID

	err = d.profileUC.Update(c.Request().Context(), req)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, nil)
}
