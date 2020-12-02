package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

func GetHook(c *gin.Context) {
	var hook models.Hook
	var err error
	var hooks *[]models.Hook
	bindErr := c.ShouldBind(&hook)
	if bindErr == nil {
		hooks, err = services.GetHook(hook)
	}
	if err != nil {
		errorLogger := gin.DefaultErrorWriter
		errorLogger.Write([]byte(err.Error() + "\n"))
		c.JSON(http.StatusOK, gin.H{
			"status": "FAIL",
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data": hooks,
	})
}

func AddHook(c *gin.Context) {

}

func UpdateHook(c *gin.Context) {

}

func DeleteHook(c *gin.Context) {

}