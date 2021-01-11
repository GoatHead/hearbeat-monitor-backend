package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

type pagedServiceHistory struct {
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	ItemCount int `json:"itemCount"`
	Data []models.Service `json:"data"`
}

func GetServiceHistory(c *gin.Context) {
	var searchCondition *models.SearchCondition
	var err1 error
	var err2 error
	var cnt int
	var serviceList *[]models.Service
	bindErr := c.ShouldBind(&searchCondition)
	var result pagedServiceHistory
	if bindErr == nil {

		cnt, err1 = services.GetServiceHistoryCnt(searchCondition)
		serviceList, err2 = services.GetServiceHistory(searchCondition)
		result = pagedServiceHistory{
			Page: searchCondition.PageStart,
			PageSize: searchCondition.PageSize,
			Data: nil,
		}

		if serviceList != nil {
			result.ItemCount = cnt
			result.Data = *serviceList
		}
	}

	errorString := ""
	if err1 != nil {
		errorString += err1.Error() + "\n"
	}
	if err2 != nil {
		errorString += err2.Error() + "\n"
	}

	if errorString != "" {
		errorLogger := gin.DefaultErrorWriter
		errorLogger.Write([]byte(errorString + "\n"))
		c.JSON(http.StatusOK, gin.H{
			"status": "FAIL",
			"data": nil,
			"message": errorString,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data": result,
		"message": "하트비트 검사 이력 조회에 성공하였습니다.",
	})
}