package middleware

import (
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/initializers"
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	fmt.Println(c.GetHeader("Authorization"))
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized2",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized3",
		})
		return
	}

	var user models.User
	if result := initializers.DB.Where("email = ?", claims["sub"]).First(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized4",
		})
		return
	}

	c.Set("user", user)
}
