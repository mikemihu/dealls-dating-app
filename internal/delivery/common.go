package delivery

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"dealls-dating-app/internal/constant"
)

// handlerErrorResponse prevent sensitive error message being exposed to public
func handlerErrorResponse(err error) error {
	if errors.Is(err, constant.ErrUnauthorized) {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if errors.Is(err, constant.ErrNoAccess) {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	if errors.Is(err, constant.ErrNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if errors.Is(err, constant.ErrUnimplemented) {
		return echo.NewHTTPError(http.StatusNotImplemented, err.Error())
	}
	if errors.Is(err, constant.ErrUserNotFound) {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if errors.Is(err, constant.ErrSwipeLimitReached) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
