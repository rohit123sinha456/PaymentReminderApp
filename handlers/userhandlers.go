package handlers

import (
    "net/http"
    "strconv"
    "log"
    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"

)

// UserHandler handles requests related to users
type UserHandler struct {
    Repository *UserRepository
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }
    user, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
    users, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    var user User
    userID := c.Param("id")
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := h.Repository.DB.Model(&User{}).Where("id = ?",userID).Updates(user); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }
    log.Printf("%d",uint(id))
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

