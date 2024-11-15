package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"bee/config"
	"bee/internal/controller"
	"bee/internal/service"
	storage "bee/internal/storage/inmemory"
	"bee/pkg/logger"
	httpserver "bee/pkg/server/http"
	"bee/pkg/storage/inmemory"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(c.Logger.Level, loggerFile)

	authStorage := storage.NewAuthStorage(inmemory.NewStorage())
	profileStorage := storage.NewProfileStorage(inmemory.NewStorage())

	authService := service.NewAuthService(l, authStorage)
	profileService := service.NewProfileService(l, profileStorage)

	handler := gin.New()
	con := controller.NewRouter(handler)
	con.SetV1Routes(l, authService, profileService, &c.HTTP.Admin)

	router := httpserver.New(handler, httpserver.Port(c.HTTP.Port))

	err = router.Start()
	if err != nil {
		log.Fatal(err)
	}
}
