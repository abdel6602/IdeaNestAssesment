package controllers

import (
	"example/backend/pkg/database/models"
	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@gmail.com", Password: "password"},
	{ID: 2, Name: "Jane Doe", Email: "Jane@gmail.com", Password: "password"},
}


func UserExists(username string) bool {
	for i := range users {
		if users[i].Email == username {
			return true
		}
	}
	return false
}

func GetUserByEmail(email string) *models.User {
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(user models.User){
	users = append(users, user)
}