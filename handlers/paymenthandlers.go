package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"
)


type PaymentHandler struct {
    Repository *PaymentRepository
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
    var payment Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
        return
    }
    payment, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
    payments, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
    var payment Payment
    paymentID := c.Param("id")
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := h.Repository.DB.Model(&Payment{}).Where("id = ?",paymentID).Updates(payment); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
        return
    }
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (h *PaymentHandler) GetPaymentsByClientID(c *gin.Context) {
    clientID, err := strconv.Atoi(c.Param("client_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
        return
    }

    payments, err := h.Repository.GetPaymentsByClientID(uint(clientID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetPaymentReminders(c *gin.Context) {
	clientID, exists := c.Get("clientID")
	if !exists {
		// slog.Error("Client ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var payments []Payment
	if err := h.Repository.DB.Where("client_id = ?", clientID).Find(&payments).Error; err != nil {
		// slog.Errorf("Failed to fetch payments for client ID %d: %s", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

