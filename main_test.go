package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	// Set up the Gin router
	router := gin.Default()

	// Define the route handler
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Create a request to send to the route
	req, _ := http.NewRequest("GET", "/hello", nil)
	resp := httptest.NewRecorder()

	// Send the request
	router.ServeHTTP(resp, req)

	// Assert that the response was as expected
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"message":"Hello, World!"}`, resp.Body.String())
}