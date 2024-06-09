package contexts

import (
	"dealls-dating-app/internal/entity"

	"github.com/labstack/echo/v4"
)

const (
	_CtxUser = "user"
)

func SetUser(c echo.Context, user entity.User) {
	c.Set(_CtxUser, user)
}

func GetUser(c echo.Context) entity.User {
	if c.Get(_CtxUser) == nil {
		return entity.User{}
	}
	return c.Get(_CtxUser).(entity.User)
}
