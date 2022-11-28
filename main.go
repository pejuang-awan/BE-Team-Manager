package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pejuang-awan/BE-Team-Manager/config"
	"github.com/pejuang-awan/BE-Team-Manager/router"
	"github.com/pejuang-awan/BE-Team-Manager/services"
)

func main() {
	config.InitializeConfig()

	if config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	services.InitializeDatabase()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic("[ERROR] Failed to listen and serve")
	}
}
