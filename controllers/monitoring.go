package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MonitoringController struct{}

func (h MonitoringController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
