package handlers

import (
	"log"
	"net/http"
    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/helpers"
	. "github.com/rohit123sinha456/payredapp/storage"
)


type LoginHandler struct {
	UserRepo *ClientRepository
}

type loginRequest struct {
	Email string `json:"email" binding:"required"`
}

func (h *LoginHandler) Login(c *gin.Context) {
	var email loginRequest
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := h.UserRepo.FindByEmail(email.Email)
	if err != nil {
		log.Printf("Failed to find client ID for email %s: %s", email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := GenerateJWT(client.ID)
	if err != nil {
		log.Printf("Failed to generate token for client ID %d: %s", client.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

