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
	// create resources

	// monitoring
	router.GET("/readyz", monitoring.Status)
	router.GET("/livez", monitoring.Status)

	// start http server
	router.Run()
}
