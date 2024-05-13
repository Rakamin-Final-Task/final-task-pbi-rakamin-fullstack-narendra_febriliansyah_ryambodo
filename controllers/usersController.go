package controllers

import (
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/initializers"
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type RequestBodyRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBodyUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RequestBodyDelete struct {
	Email string `json:"email"`
}

func UserRegisterController(c *gin.Context) {
	var requestBody RequestBodyRegister

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"errors": "Register failed",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"errors": "Register failed",
		})
		return
	}
	user := models.User{Username: requestBody.Username, Email: requestBody.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"errors": "Register failed",
		})
		return
	}

	newData := initializers.DB.First(&user, "email = ?", requestBody.Email)
	fmt.Println("tes", newData)
	c.JSON(200, gin.H{
		"username": requestBody.Username,
		"email":    requestBody.Email,
	})

}

func UserLoginController(c *gin.Context) {
	var requestBody RequestBodyLogin

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"errors": "Login failed",
		})
		return
	}

	user := models.User{Email: requestBody.Email, Password: requestBody.Password}

	result := initializers.DB.Where("email = ?", requestBody.Email).First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"errors": "Login failed1",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"errors": "Login failed",
		})
		return
	}

	key := []byte(os.Getenv("SECRET"))
	var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": requestBody.Email})
	var tokenString, tokenError = t.SignedString(key)

	if tokenError != nil {

		c.JSON(400, gin.H{
			"errors": "Login failed",
			"token":  tokenString,
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func UserUpdateController(c *gin.Context) {
	var requestBody RequestBodyUpdate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"errors": "Update failed",
		})
		return
	}
	user := models.User{Username: requestBody.Username, Email: requestBody.Email}

	result := initializers.DB.Model(&user).Where("email = ?", requestBody.Email).Update("username", requestBody.Username)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"errors": "Update failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"username": requestBody.Username,
		"email":    requestBody.Email,
	})

}

func UserDeleteController(c *gin.Context) {
	var requestBody RequestBodyDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"errors": "Delete failed",
		})
		return
	}
	user := models.User{Email: requestBody.Email}

	result := initializers.DB.Unscoped().Where("email = ?", requestBody.Email).Delete(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"errors": "Delete failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"email": requestBody.Email,
	})

}
