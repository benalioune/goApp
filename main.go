package main

import (
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterBookRoutes(r)
	r.Run() // Default listens on :8080
}
