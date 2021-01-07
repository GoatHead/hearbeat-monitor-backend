package scheduler

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/repositories"
	"goathead/heartbeat-monitor-backend/services"
	"strconv"
	"time"
)

func ScheduledHeartBeat() {
	heartBeat()

	ticker := time.NewTicker(getDuration())
	go func() {
		for {
			select {
			case <-ticker.C:
				heartBeat()
				ticker.Reset(getDuration())
			}
		}
	}()
}

func getDuration() time.Duration {
	var seconds int
	setting, _ := services.GetApplicationSettings(nil)

	if setting != nil {
		seconds = setting.CycleSec
	} else {
		seconds = 300
	}

	logger := gin.DefaultWriter
	logger.Write([]byte("SCHEDULED SECS: " + strconv.Itoa(seconds) + "\n"))

	return time.Duration(seconds) * time.Second
}

func heartBeat() {

	logger := gin.DefaultWriter
	logger.Write([]byte("Heartbeat Check Start\n"))

	serviceList, err := services.GetService(nil)

	if serviceList != nil {
		if len(*serviceList) > 0 && err == nil {

			for _, service := range *serviceList {

				statusCode := services.CheckHeartBeat(&service)
				service.Status = statusCode

				log := "===HeartBeat===\n"
				log += "url: " + service.Url
				log += "; status: " + strconv.Itoa(statusCode) + "\n"
				log += "===============\n"
				logger.Write([]byte(log))
				_ = repositories.UpdateServiceStatusCode(&service)

			}

		}
	}
}
