package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func LimitRequestGet() gin.HandlerFunc {
	limit := rate.NewLimiter(rate.Every(1*time.Minute), 20)
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "too many request",
			})
		}

		c.Next()
	}
}

func LimitRequestPost() gin.HandlerFunc {
	limit := rate.NewLimiter(rate.Every(1*time.Minute), 2)
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "too many request",
			})
		}

		c.Next()
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get("X-API-KEY")

		if key != os.Getenv("API_KEY") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
