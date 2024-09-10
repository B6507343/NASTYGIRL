package main

import (
	"net/http"

	"example.com/project/config"
	"example.com/project/controller"
	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func main() {
	config.ConnectionDB()
	config.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Group API Routes
	router := r.Group("/api")
	{
		router.POST("/checkin/:ticket_id", controller.CreateTicketCheck)
 // แก้ไขจาก tickectcheck เป็น ticketcheck
	}

	// Route สำหรับการเช็คสถานะของ API
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server
	r.Run("localhost:" + PORT)
}

// CORS Middleware เพื่อจัดการการเข้าถึงจาก client ต่าง ๆ
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
