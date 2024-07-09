package handlers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "github.com/rohit123sinha456/payredapp/helpers"
	. "github.com/rohit123sinha456/payredapp/storage"
	. "github.com/rohit123sinha456/payredapp/models"

)


type UserloginHandler struct {
	UserRepo *UserRepository
}

// userloginRequest represents the login request payload.
type userloginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// userloginResponse represents the login response payload.
type userloginResponse struct {
	Token string `json:"token"`
}

// Login handles user login and JWT token generation.
func (h *UserloginHandler) Login(c *gin.Context) {
	var req userloginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := h.UserRepo.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, userloginResponse{Token: token})
}
