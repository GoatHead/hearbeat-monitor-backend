package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"goathead/heartbeat-monitor-backend/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/api/hook", controllers.GetHook)
	r.POST("/api/hook", controllers.AddHook)
	r.PUT("/api/hook", controllers.UpdateHook)
	r.DELETE("/api/hook", controllers.DeleteHook)

	r.GET("/api/service", controllers.GetService)
	r.POST("/api/service", controllers.AddService)
	r.PUT("/api/service", controllers.UpdateService)
	r.DELETE("/api/service", controllers.DeleteService)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
