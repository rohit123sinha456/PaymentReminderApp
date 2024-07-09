package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"
)

type ClientHandler struct {
    Repository *ClientRepository
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
    var client Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) GetClient(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
        return
    }
    client, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
        return
    }
    c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) GetAllClients(c *gin.Context) {
    clients, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
    var client Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Update(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
        return
    }
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (h *ClientHandler) GetClientByEmail(c *gin.Context) {
    email := c.Query("email")
    if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email query parameter is required"})
        return
    }

    client, err := h.Repository.FindByEmail(email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"client_id": client.ID})
}
