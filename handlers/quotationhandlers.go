package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
	. "github.com/rohit123sinha456/payredapp/storage"
    . "github.com/rohit123sinha456/payredapp/models"
)

type QuotationHandler struct {
    Repository *QuotationRepository
}

func (h *QuotationHandler) CreateQuotation(c *gin.Context) {
    var quotation Quotation
    if err := c.ShouldBindJSON(&quotation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Create(&quotation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, quotation)
}

func (h *QuotationHandler) GetQuotation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quotation ID"})
        return
    }
    quotation, err := h.Repository.Read(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "quotation not found"})
        return
    }
    c.JSON(http.StatusOK, quotation)
}

func (h *QuotationHandler) GetAllQuotations(c *gin.Context) {
    quotations, err := h.Repository.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, quotations)
}

func (h *QuotationHandler) UpdateQuotation(c *gin.Context) {
    var quotation Quotation
    if err := c.ShouldBindJSON(&quotation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repository.Update(&quotation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, quotation)
}

func (h *QuotationHandler) DeleteQuotation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quotation ID"})
        return
    }
    if err := h.Repository.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
