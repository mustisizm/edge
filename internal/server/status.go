package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET Signer reports if the server is operational.
//
// GET /status
func (s *Config) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Edge",
		"Version": "0.0.2",
	})
}
