package handlers

import (
	"example/backend/pkg/controllers"
	"example/backend/pkg/database/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// sign in handler function
func SignUp(c *gin.Context) {
	// get user data from request's body
	var newUser models.User
	// convert request body to user type
	err := c.BindJSON(&newUser)
	// check for errors and return error message
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	// check if user already exists
	if controllers.UserExists(newUser.Email) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return
	}
	// hash user password
	hashedPassword, err := controllers.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password",
		})
		return
	}
	newUser.Password = hashedPassword
	// if user does not exist, save user to database
	// save user to database
	controllers.AddUser(newUser)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func SignIn(c *gin.Context) {
	// get user data from request's body
	var newUser models.User
	// convert request body to user type
	err := c.BindJSON(&newUser)
	// check for errors and return error message
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	// check if user exists via email
	if controllers.UserExists(newUser.Email) {
		user := controllers.GetUserByEmail(newUser.Email)
		// if user exists, check password
		if user != nil {
			// compare password
			if controllers.CheckPasswordHash(newUser.Password, user.Password) {
				// create authentication token
				accessToken := controllers.GenerateAccessToken(user.ID)
				// create refresh token
				refreshToken := controllers.GenerateRefreshToken(user.ID)
				// send response
				c.JSON(http.StatusOK, gin.H{
					"message":       "User signed in successfully",
					"access_token":  accessToken,
					"refresh_token": refreshToken,
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Incorrect password.",
				})
				return

			}
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exist",
		})
		return
	}
}
