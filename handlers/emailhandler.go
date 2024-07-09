package handlers

import (
    "net/http"
    "log"
    "fmt"
	"gorm.io/gorm"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/config"
	. "github.com/rohit123sinha456/payredapp/engage"
    . "github.com/rohit123sinha456/payredapp/models"
)


// EmailHandler handles the email sending requests
type EmailHandler struct {
	DB *gorm.DB
	Config ConfigStruct
}

func (h *EmailHandler) SendReminders(c *gin.Context) {
	var request struct {
		ClientIDs []uint `json:"client_ids"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if len(request.ClientIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No client IDs provided"})
		return
	}

	var clients []Client
	if err := h.DB.Where("id IN ?", request.ClientIDs).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	for _, client := range clients {
		var payments []Payment
		if err := h.DB.Where("client_id = ? AND is_paid = ?", client.ID, false).Find(&payments).Error; err != nil {
			log.Printf("Failed to fetch payments for client ID %d: %s", client.ID, err)
			continue
		}

		if len(payments) == 0 {
			continue
		}

		to := append(client.Emails, client.LoginEmail)
		subject := "Payment Reminder"
		msgBody := fmt.Sprintf("Dear %s,\n\nYou have the following pending payments:\n\n", client.Name)
		for _, payment := range payments {
			msgBody += fmt.Sprintf("Invoice Number: %s, Amount: %.2f, Due Date: %s\n",
				payment.InvoiceNumber, payment.InvoiceAmount, payment.InvoiceDate)
		}
		msgBody += "\nPlease make your payment at your earliest convenience.\n\nThank you."

		SendEmail(h.Config, to, subject, msgBody)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Emails sent successfully"})
}