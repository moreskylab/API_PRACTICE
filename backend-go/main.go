package main

import (
	"backend-go/config"
	"backend-go/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Apply customized CORS configuration
	r.Use(config.SetupCORS())

	objController := controller.NewObjectController()

	// Map application endpoints
	r.GET("/objects", objController.GetManyObjects)
	r.GET("/objects/:id", objController.GetOneObject)
	r.POST("/objects", objController.PostOneObject)
	r.PUT("/objects/:id", objController.PutOneObject)
	r.PATCH("/objects/:id", objController.PatchOneObject)
	r.DELETE("/objects/:id", objController.DeleteOneObject)

	r.Run(":8081") // Launches on port 8080
}
