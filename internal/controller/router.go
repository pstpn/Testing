package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"bee/config"
	docsv1 "bee/docs/v1"
	routesv1 "bee/internal/controller/v1/http"
	"bee/internal/service"
	"bee/pkg/logger"
)

type Controller struct {
	handler     *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	docsv1.SwaggerInfov1.BasePath = "/api/v1"

	v1 := handler.Group("/api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
		v1.GET("/healthcheck", healthCheck)
	}

	return &Controller{
		handler:     handler,
		routerGroup: v1,
	}
}

// healthCheck godoc
//
//	@Summary		Проверка здоровья
//	@Description	Проверка на жизнеспособность
//	@Tags			system
//	@Success		200	{string} string "Сервис жив"
//	@Failure		404	"Сервис мертв"
//	@Router			/healthcheck [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now().String())
}

func (c *Controller) SetV1Routes(
	l logger.Interface,
	authService service.AuthService,
	profileService service.ProfileService,
	adminConfig *config.AdminConfig,
) {
	routesv1.SetRoutes(
		c.routerGroup,
		l,
		authService,
		profileService,
		adminConfig,
	)
}
