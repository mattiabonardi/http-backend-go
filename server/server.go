package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mattiabonardi/http-backend-go/controllers"
)

func Init() {
	// create router
	router := gin.New()
	// create controllers
	monitoring := new(controllers.MonitoringController)
	authentication := new(controllers.AuthenticationController)

	// create resources
	// monitoring
	router.GET("/readyz", monitoring.Status)
	router.GET("/livez", monitoring.Status)

	// api v1
	v1 := router.Group("v1")

	// authentication
	v1.POST("/login", authentication.Login)

	// start http server
	router.Run()
}
