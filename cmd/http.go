package cmd

import (
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/api"
	"ewallet-fastcampus/internal/repository"
	"ewallet-fastcampus/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}

	registerService := &services.RegisterService{
		RegisterRepo: registerRepo,
	}

	registerApi := &api.RegisterHandler{
		RegisterServices: registerService,
	}

	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerApi.RegisterHandler)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))

	if err != nil {
		log.Fatal(err)
	}
}
