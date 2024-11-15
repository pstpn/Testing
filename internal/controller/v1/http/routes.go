package http

import (
	"github.com/gin-gonic/gin"

	"bee/config"
	"bee/internal/controller/v1/http/admin"
	"bee/internal/controller/v1/http/auth"
	"bee/internal/controller/v1/http/user"
	"bee/internal/service"
	"bee/pkg/logger"
)

func setAdminRoute(
	handler gin.IRoutes,
	l logger.Interface,
	profileService service.ProfileService,
) {
	p := admin.NewProfileController(l, profileService)
	roles := map[string]struct{}{
		"admin": {},
	}

	handler.POST("/profiles", auth.RoleMiddleware(roles), p.CreateProfile)
	handler.PUT("/profiles/:email", auth.RoleMiddleware(roles), p.UpdateProfile)
	handler.DELETE("/profiles/:email", auth.RoleMiddleware(roles), p.DeleteProfile)
}

func setAuthRoute(
	handler gin.IRoutes,
	l logger.Interface,
	authService service.AuthService,
) {
	a := auth.NewAuthController(l, authService)

	handler.POST("/register", a.Register)
}

func setUserRoute(
	handler gin.IRoutes,
	l logger.Interface,
	profileService service.ProfileService,
) {
	p := user.NewProfileController(l, profileService)
	roles := map[string]struct{}{
		"admin": {},
		"user":  {},
	}

	handler.GET("/profiles/:email", auth.RoleMiddleware(roles), p.GetProfile)
}

func SetRoutes(
	handler gin.IRoutes,
	l logger.Interface,
	authService service.AuthService,
	profileService service.ProfileService,
	adminConfig *config.AdminConfig,
) {
	middleware := auth.NewBasicMiddleware(l, authService, adminConfig)
	setAuthRoute(handler, l, authService)
	authGroup := handler.Use(middleware.BasicAuth())
	{
		setAdminRoute(authGroup, l, profileService)
		setUserRoute(authGroup, l, profileService)
	}
}
