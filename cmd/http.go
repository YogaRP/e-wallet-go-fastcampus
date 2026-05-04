package cmd

import (
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/api"
	"ewallet-fastcampus/internal/interfaces"
	"ewallet-fastcampus/internal/repository"
	"ewallet-fastcampus/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {

	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterApi.RegisterHandler)
	userV1.POST("/login", dependency.LoginApi.Login)
	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutApi.Logout)
	userV1.PUT("/refresh-token", dependency.MiddlewareRefreshToken, dependency.RefreshTokenApi.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))

	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepo

	HealthcheckAPI  interfaces.IHealthcheckHandler
	RegisterApi     interfaces.IRegisterHandler
	LoginApi        interfaces.ILoginHandler
	LogoutApi       interfaces.ILogoutHandler
	RefreshTokenApi interfaces.IRerfreshTokenHandler

	TokenValidationApi *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerService := &services.UserService{
		UserRepo: userRepo,
	}

	registerApi := &api.RegisterHandler{
		RegisterServices: registerService,
	}

	loginService := &services.LoginService{
		UserRepo: userRepo,
	}

	loginApi := &api.LoginHandler{
		LoginService: loginService,
	}

	logoutService := &services.LogoutService{
		UserRepo: userRepo,
	}

	logoutApi := &api.LogoutHandler{
		LogoutService: logoutService,
	}

	refreshTokenService := &services.RefreshTokenService{
		UserRepo: userRepo,
	}

	refreshTokenApi := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenService,
	}

	tokenValidationService := &services.TokenValidationService{
		UserRepo: userRepo,
	}

	tokenValidationApi := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationService,
	}

	return Dependency{
		UserRepository:     userRepo,
		HealthcheckAPI:     healthcheckAPI,
		RegisterApi:        registerApi,
		LoginApi:           loginApi,
		LogoutApi:          logoutApi,
		RefreshTokenApi:    refreshTokenApi,
		TokenValidationApi: tokenValidationApi,
	}

}
