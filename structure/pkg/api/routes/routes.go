package main

import (
	"example/backend/pkg/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	app := gin.Default()
	// default route
	app.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API, Happy to Serve You",
		})
	})

	// Sign Up Route 
	app.POST("/signup", handlers.SignUp)
	// // Sign In Route
	app.POST("/signin", handlers.SignIn)
	// // Refresh Token Route
	// app.POST("/refresh-token", func(c *gin.Context))
	// // Create Organization Route
	// app.POST("/organization", func(c *gin.Context))
	// // Get Organization Route
	// app.GET("/organization/:organization_id", func(c *gin.Context))
	app.Run("localhost:8080")
}