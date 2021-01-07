package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

func GetApplicationSettings(c *gin.Context) {
	var applicationSettings *models.ApplicationSettings
	var err error
	var result *models.ApplicationSettings
	bindErr := c.ShouldBind(&applicationSettings)
	if bindErr == nil {
		result, err = services.GetApplicationSettings(applicationSettings)
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
		"data": result,
		"message": "애플리케이션 설정 조회에 성공하였습니다.",
	})
}

func UpdateApplicationSettings(c *gin.Context) {
	var applicationSettings *models.ApplicationSettings
	var err error
	bindErr := c.ShouldBind(&applicationSettings)
	if bindErr == nil {
		err = services.UpdateApplicationSettings(applicationSettings)
	}
	if err != nil {
		errorLogger := gin.DefaultErrorWriter
		errorLogger.Write([]byte(err.Error() + "\n"))
		c.JSON(http.StatusOK, gin.H{
			"status": "FAIL",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"message": "애플리케이션 설정 수정에 성공하였습니다.",
	})
}