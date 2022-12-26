package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattiabonardi/http-backend-go/controllers"
	"github.com/mattiabonardi/http-backend-go/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	// create router
	router := gin.New()
	// create controllers
	monitoring := new(controllers.MonitoringController)
	authController := new(controllers.AuthenticationController)

	// swagger
	router.StaticFS("/swagger/", http.Dir("swagger"))

	// create resources
	// monitoring
	router.GET("/readyz", monitoring.Status)
	router.GET("/livez", monitoring.Status)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// api
	api := router.Group("api")
	// api v1
	v1 := api.Group("v1")

	// authentication
	authentication := v1.Group("authentication")
	authentication.POST("/login", authController.Login)
	authentication.POST("/refresh_token", middlewares.AuthorizationMiddleware, authController.RefreshToken)

	// start http server
	router.Run()
}
