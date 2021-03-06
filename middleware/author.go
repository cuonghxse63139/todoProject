package middleware

import (
	"log"
	"net/http"
	"strings"
	"todoProject/config"
	"todoProject/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		splitToken := strings.Split(header, "Bearer ")

		if len(splitToken) != 2 {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		reqToken := splitToken[1]

		if reqToken == "" {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.SECRET_JWT_TOKEN), nil
		})

		if err != nil {
			handleUnauthorized(c, err.Error())
		} else if !token.Valid {
			log.Printf("Token is invalid %s\n", reqToken)
			handleUnauthorized(c, "Invalid Token")
		}

		claims, isOk := token.Claims.(jwt.MapClaims)

		if isOk {
			c.Set(config.TOKEN_CURRENT_USER_ID, claims[config.TOKEN_CURRENT_USER_ID])
			c.Set(config.TOKEN_CURRENT_USERNAME, claims[config.TOKEN_CURRENT_USERNAME])
			c.Set(config.TOKEN_CURRENT_USER_ROLE, claims[config.TOKEN_CURRENT_USER_ROLE])
		} else {
			log.Panicf("Cannot extract claims from token %s\n", reqToken)
			handleUnauthorized(c, "Invalid Token")
		}

	}
}

func handleUnauthorized(c *gin.Context, errorMsg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.BadRequestResponse{ErrorMessage: errorMsg})
}
