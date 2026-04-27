package notes

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	notes := rg.Group("/notes")
	{
		notes.POST("", h.CreateNote)
		notes.GET("", h.ListNotes)
		notes.GET("/:id", h.GetNoteByID)
		notes.PUT("/:id", h.UpdateNote)
		notes.DELETE("/:id", h.DeleteNote)
	}
}

