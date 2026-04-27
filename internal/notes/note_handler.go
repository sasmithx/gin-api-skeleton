package notes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateNote(c *gin.Context) {
	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.repo.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing note id"})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *Handler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing note id"})
		return
	}

	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := h.repo.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing note id"})
		return
	}

	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete note"})
		return
	}

	c.Status(http.StatusNoContent)
}

