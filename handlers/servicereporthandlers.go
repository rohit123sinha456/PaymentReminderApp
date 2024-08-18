package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"
)

type ServiceReportHandler struct {
    Repository *ServiceReportRepository
}

func (h *ServiceReportHandler) CreateServiceReport(c *gin.Context) {
    var serviceReport ServiceReport
    if err := c.ShouldBindJSON(&serviceReport); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&serviceReport); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, serviceReport)
}

func (h *ServiceReportHandler) GetServiceReport(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service report ID"})
        return
    }
    serviceReport, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "service report not found"})
        return
    }
    c.JSON(http.StatusOK, serviceReport)
}

func (h *ServiceReportHandler) GetAllServiceReports(c *gin.Context) {
    serviceReports, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, serviceReports)
}

func (h *ServiceReportHandler) UpdateServiceReport(c *gin.Context) {
    var serviceReport ServiceReport
    serviceReportID := c.Param("id")
    if err := c.ShouldBindJSON(&serviceReport); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := h.Repository.DB.Model(&ServiceReport{}).Where("id = ?",serviceReportID).Updates(serviceReport); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusOK, serviceReport)
}

func (h *ServiceReportHandler) DeleteServiceReport(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service report ID"})
        return
    }
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
// GetServiceReports fetches all service reports for the logged-in client.
func (h *ServiceReportHandler) GetServiceReports(c *gin.Context) {
	clientID, exists := c.Get("clientID")
	if !exists {
		// slog.Error("Client ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var serviceReports []ServiceReport
	if err := h.Repository.DB.Where("client_id = ?", clientID).Find(&serviceReports).Error; err != nil {
		// slog.Errorf("Failed to fetch service reports for client ID %d: %s", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service reports"})
		return
	}

	c.JSON(http.StatusOK, serviceReports)
}

func (h *ServiceReportHandler) GetServiceReportsByClientID(c *gin.Context) {
    clientID, err := strconv.Atoi(c.Param("client_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
        return
    }

    payments, err := h.Repository.GetServiceReportsByClientID(uint(clientID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, payments)
}