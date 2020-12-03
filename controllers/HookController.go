package controllers

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
)

func GetHook(c *gin.Context) {
	var hook *models.Hook
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
		"message": "훅 목록 조회에 성공하였습니다.",
	})
}

func AddHook(c *gin.Context) {
	var hook *models.Hook
	var err error
	bindErr := c.ShouldBind(&hook)
	if bindErr == nil {
		err = services.AddHook(hook)
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
		"message": "훅 추가에 성공하였습니다.",
	})
}

func UpdateHook(c *gin.Context) {
	var hook *models.Hook
	var err error
	bindErr := c.ShouldBind(&hook)
	if bindErr == nil {
		err = services.UpdateHook(hook)
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
		"message": "훅 수정에 성공하였습니다.",
	})
}

func DeleteHook(c *gin.Context) {
	var hook *models.Hook
	var err error
	bindErr := c.ShouldBind(&hook)
	if bindErr == nil {
		err = services.DeleteHook(hook)
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
		"message": "훅 삭제에 성공하였습니다.",
	})
}