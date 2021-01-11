package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

func GetServiceHistory(c *gin.Context) {
	var searchCondition *models.SearchCondition
	var err error
	var serviceList *[]models.Service
	bindErr := c.BindJSON(&searchCondition)
	if bindErr == nil {
		serviceList, err = services.GetServiceHistory(searchCondition)
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
		"message": "하트비트 검사 이력 조회에 성공하였습니다.",
	})
}