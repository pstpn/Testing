package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	routesv2 "course/internal/controller/v2/http"
	httputils "course/internal/controller/v2/http/utils"
	"course/internal/service"
	"course/pkg/logger"
)

type Controller struct {
	handler      *gin.Engine
	routerGroups map[string]*gin.RouterGroup
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Disable CORS
	handler.OPTIONS("/*any", httputils.DisableCors)

	// Swagger settings
	v2 := handler.Group("/api/v2")
	{
		v2.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
		v2.GET("/healthcheck", healthCheck)
	}

	return &Controller{
		handler: handler,
		routerGroups: map[string]*gin.RouterGroup{
			"v2": v2,
		},
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

func (c *Controller) SetV2Routes(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) {
	routesv2.SetRoutes(
		c.routerGroups["v2"],
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		authService,
	)
}
