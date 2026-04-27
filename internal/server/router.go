package server

import (
	"net/http"

	"api-skeleton/internal/notes"

	"github.com/gin-gonic/gin"
)

func NewRouter(notesHandler *notes.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     "true",
			"status": "healthy",
		})
	})

	notes.RegisterRoutes(api, notesHandler)

	return r

}
