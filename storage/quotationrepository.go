package storage

import(
		"gorm.io/gorm"
		"github.com/rohit123sinha456/payredapp/models"
)

type QuotationRepository struct {
    DB *gorm.DB
}

func (r *QuotationRepository) Create(quotation *models.Quotation) error {
    var client models.Client
    if err := r.DB.First(&client, quotation.ClientID).Error; err != nil {
        return err
    }
    quotation.Client = client
    return r.DB.Create(quotation).Error
}

func (r *QuotationRepository) Read(id uint) (interface{}, error) {
    var quotation models.Quotation
    if err := r.DB.Preload("Client").First(&quotation, id).Error; err != nil {
        return nil, err
    }
    return &quotation, nil
}

func (r *QuotationRepository) Update(quotation models.Quotation) error {
    var client models.Client
    if err := r.DB.First(&client, quotation.ClientID).Error; err != nil {
        return err
    }
    quotation.Client = client
    return r.DB.Save(quotation).Error
}

func (r *QuotationRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Quotation{}, id).Error
}

func (r *QuotationRepository) ReadAll() ([]models.Quotation, error) {
    var quotations []models.Quotation
    if err := r.DB.Preload("Client").Find(&quotations).Error; err != nil {
        return nil, err
    }
    var result []models.Quotation
    for i := range quotations {
        result = append(result, quotations[i])
    }
    return result, nil
}

func (r *QuotationRepository) GetQuotationByClientID(clientID uint) ([]models.Quotation, error) {
    var quotations []models.Quotation
    if err := r.DB.Preload("Client").Where("client_id = ?", clientID).Order("quotation_total_value DESC, quotation_date DESC").Find(&quotations).Error; err != nil {
        return nil, err
    }
    return quotations, nil
}