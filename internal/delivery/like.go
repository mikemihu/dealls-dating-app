package delivery

import (
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/internal/pkg/contexts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeDelivery struct {
	cfg    *config.Cfg
	likeUC internal.LikeUC
}

func NewLikeDelivery(
	cfg *config.Cfg,
	likeUC internal.LikeUC,
) internal.LikeDelivery {
	return &LikeDelivery{
		cfg:    cfg,
		likeUC: likeUC,
	}
}

func (d *LikeDelivery) MutualList(c echo.Context) error {
	req := entity.MutualListRequest{
		UserID: contexts.GetUser(c).ID,
	}
	profiles, err := d.likeUC.MutualList(c.Request().Context(), req)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, profiles)
}
