package services

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/core/webclient"
	"goathead/heartbeat-monitor-backend/repositories"
)

func CheckHeartBeat(service *models.Service) int {
	url := service.Url
	statusCode := webclient.Request(url)

	service.Status = statusCode

	 err := repositories.UpdateServiceStatusCode(service)

	 if err != nil {
		 errorLogger := gin.DefaultErrorWriter
		 errorLogger.Write([]byte(err.Error() + "\n"))
	 }

	return statusCode
}