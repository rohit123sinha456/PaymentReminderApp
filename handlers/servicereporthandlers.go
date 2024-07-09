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
    if err := c.ShouldBindJSON(&serviceReport); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Update(&serviceReport); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
