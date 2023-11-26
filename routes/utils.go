package routes

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func bindBodyOrError(c *gin.Context, body any) error {
	err := c.ShouldBindJSON(body)
	if err == nil {
		return nil
	}

	structType := reflect.TypeOf(body).Elem()
	var jsonTags []string
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonTag := field.Tag.Get("json")
		jsonTags = append(jsonTags, jsonTag)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": fmt.Sprintf("Required fields: %s.", strings.Join(jsonTags, ", ")),
	})
	return err
}

func generateJWT(userID uint, userRole int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userID,
		"role": userRole,
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}
