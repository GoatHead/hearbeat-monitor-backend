package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"goathead/heartbeat-monitor-backend/controllers"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/scheduler"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	database.GetDbConnector()

	scheduler.ScheduledHeartBeat()

	r.GET("/api/hook", controllers.GetHook)
	r.POST("/api/hook", controllers.AddHook)
	r.PUT("/api/hook", controllers.UpdateHook)
	r.DELETE("/api/hook", controllers.DeleteHook)

	r.GET("/api/service", controllers.GetService)
	r.POST("/api/service", controllers.AddService)
	r.PUT("/api/service", controllers.UpdateService)
	r.DELETE("/api/service", controllers.DeleteService)

	r.GET("/api/application-settings", controllers.GetApplicationSettings)
	r.PUT("/api/application-settings", controllers.UpdateApplicationSettings)

	r.POST("/api/heartbeat/test", controllers.TestList)
	r.POST("/api/heartbeat/testAll", controllers.TestAll)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}