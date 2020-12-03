package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

func GetService(c *gin.Context) {
	var service *models.Service
	var err error
	var serviceList *[]models.Service
	bindErr := c.ShouldBind(&service)
	if bindErr == nil {
		serviceList, err = services.GetService(service)
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
		"data": serviceList,
		"message": "서비스 목록 조회에 성공하였습니다.",
	})
}

func AddService(c *gin.Context) {
	var service *models.Service
	var err error
	bindErr := c.ShouldBind(&service)
	if bindErr == nil {
		err = services.AddService(service)
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
		"message": "서비스 추가에 성공하였습니다.",
	})
}

func UpdateService(c *gin.Context) {
	var service *models.Service
	var err error
	bindErr := c.ShouldBind(&service)
	if bindErr == nil {
		err = services.UpdateService(service)
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
		"message": "서비스 수정에 성공하였습니다.",
	})
}

func DeleteService(c *gin.Context) {
	var service *models.Service
	var err error
	bindErr := c.ShouldBind(&service)
	if bindErr == nil {
		err = services.DeleteService(service)
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
		"message": "서비스 삭제에 성공하였습니다.",
	})
}

func PingService(c *gin.Context)  {

}