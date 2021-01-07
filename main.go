package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"goathead/heartbeat-monitor-backend/controllers"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

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
