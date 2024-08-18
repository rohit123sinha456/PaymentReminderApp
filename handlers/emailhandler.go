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
		var quotations []Payment
		if err := h.DB.Where("client_id = ? AND is_paid = ?", client.ID, false).Find(&quotations).Error; err != nil {
			log.Printf("Failed to fetch quotations for client ID %d: %s", client.ID, err)
			continue
		}

		if len(quotations) == 0 {
			continue
		}

		to := append(client.Emails, client.LoginEmail)
		subject := "Payment Reminder"
		msgBody := fmt.Sprintf("Dear %s,\n\nYou have the following pending quotations:\n\n", client.Name)
		for _, payment := range quotations {
			msgBody += fmt.Sprintf("Invoice Number: %s, Amount: %.2f, Due Date: %s\n",
				payment.InvoiceNumber, payment.InvoiceAmount, payment.InvoiceDate)
		}
		msgBody += "\nPlease make your payment at your earliest convenience.\n\nThank you."

		SendEmail(h.Config, to, subject, msgBody)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Emails sent successfully"})
}




func (h *EmailHandler) SendQuotationReminders(c *gin.Context) {
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
		var quotations []Quotation
		if err := h.DB.Where("client_id = ? AND is_approved = ?", client.ID, false).Find(&quotations).Error; err != nil {
			log.Printf("Failed to fetch quotations for client ID %d: %s", client.ID, err)
			continue
		}

		if len(quotations) == 0 {
			continue
		}

		to := append(client.Emails, client.LoginEmail)
		subject := "Quotation Reminder"
		msgBody := fmt.Sprintf("Dear %s,\n\nYou have the following pending Quotations:\n\n", client.Name)
		for _, quotation := range quotations {
			msgBody += fmt.Sprintf("Quotation Number: %s, Amount: %.2f, Quotation Date: %s\n",
				quotation.QuotationNumber, quotation.QuotationTotalValue, quotation.QuotationDate)
		}
		msgBody += "\n We await for your kind approval.\n\nThank you."

		SendEmail(h.Config, to, subject, msgBody)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Emails sent successfully"})
}



func (h *EmailHandler) SendServiceReportReminders(c *gin.Context) {
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
		var sreports []ServiceReport
		if err := h.DB.Where("client_id = ? AND attention_required = ?", client.ID, false).Find(&sreports).Error; err != nil {
			log.Printf("Failed to fetch sreports for client ID %d: %s", client.ID, err)
			continue
		}

		if len(sreports) == 0 {
			continue
		}

		to := append(client.Emails, client.LoginEmail)
		subject := "Job Report"
		msgBody := fmt.Sprintf("Dear %s,\n\nThe below mentioned job report needs immediate attemtion:\n\n", client.Name)
		for _, sreport := range sreports {
			msgBody += fmt.Sprintf("Job Report Number: %s, Date: %.2f, Job Report Notes: %s\n",
			sreport.ServiceReportNumber, sreport.ServiceReportDate, sreport.ServiceReportNotes)
		}
		msgBody += "\n We await for your kind approval.\n\nThank you."

		SendEmail(h.Config, to, subject, msgBody)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Emails sent successfully"})
}