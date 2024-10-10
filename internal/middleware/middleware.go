package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"situs-forum/internal/configs"
	"situs-forum/pkg/jwt"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Missing Token"))
			return
		}

		userId, userName, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userId)
		c.Set("username", userName)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Missing Token"))
			return
		}

		userId, userName, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userId)
		c.Set("username", userName)
		c.Next()
	}
}
