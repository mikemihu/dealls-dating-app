package middleware

import (
	"net/http"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/pkg/authentication"
	"dealls-dating-app/internal/pkg/contexts"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg           *config.Cfg
	logger        *zap.Logger
	jwtMiddleware echo.MiddlewareFunc
	userUC        internal.UserUC
}

func NewMiddleware(
	cfg *config.Cfg,
	logger *zap.Logger,
	jwtConfig echojwt.Config,
	userUC internal.UserUC,
) *Middleware {
	jwtMiddleware, err := jwtConfig.ToMiddleware()
	if err != nil {
		logger.Fatal("failed to create jwt middleware", zap.Error(err))
	}

	return &Middleware{
		cfg:           cfg,
		logger:        logger,
		jwtMiddleware: jwtMiddleware,
		userUC:        userUC,
	}
}

func (m *Middleware) JWTAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return m.jwtMiddleware(func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			m.logger.Error("failed get token from echo context")
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		claims, ok := token.Claims.(*authentication.AuthClaims)
		if !ok {
			m.logger.Error("failed claims")
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		user, err := m.userUC.Get(c.Request().Context(), claims.UserID)
		if err != nil {
			m.logger.Error("failed get user", zap.Error(err))
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		contexts.SetUser(c, user)

		return next(c)
	})
}
