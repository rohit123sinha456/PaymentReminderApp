package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"
)

type NOCHandler struct {
    Repository *NOCRepository
}

func (h *NOCHandler) CreateNOC(c *gin.Context) {
    var noc NOC
    if err := c.ShouldBindJSON(&noc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&noc); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, noc)
}

func (h *NOCHandler) GetNOC(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid NOC ID"})
        return
    }
    noc, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NOC not found"})
        return
    }
    c.JSON(http.StatusOK, noc)
}

func (h *NOCHandler) GetAllNOCs(c *gin.Context) {
    nocs, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, nocs)
}

func (h *NOCHandler) UpdateNOC(c *gin.Context) {
    var noc NOC
    if err := c.ShouldBindJSON(&noc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Update(&noc); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, noc)
}

func (h *NOCHandler) DeleteNOC(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid NOC ID"})
        return
    }
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
// GetNOCs fetches all NOCs for the logged-in client.
func (h *NOCHandler) GetNOCs(c *gin.Context) {
	clientID, exists := c.Get("clientID")
	if !exists {
		// slog.Error("Client ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var nocs []NOC
	if err := h.Repository.DB.Where("client_id = ?", clientID).Find(&nocs).Error; err != nil {
		// slog.Errorf("Failed to fetch NOCs for client ID %d: %s", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NOCs"})
		return
	}

	c.JSON(http.StatusOK, nocs)
}