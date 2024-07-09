package storage

import(
		"gorm.io/gorm"
		. "github.com/rohit123sinha456/payredapp/models"
)

type QuotationRepository struct {
    DB *gorm.DB
}

func (r *QuotationRepository) Create(quotation interface{}) error {
    return r.DB.Create(quotation).Error
}

func (r *QuotationRepository) Read(id uint) (interface{}, error) {
    var quotation Quotation
    if err := r.DB.First(&quotation, id).Error; err != nil {
        return nil, err
    }
    return &quotation, nil
}

func (r *QuotationRepository) Update(quotation interface{}) error {
    return r.DB.Save(quotation).Error
}

func (r *QuotationRepository) Delete(id uint) error {
    return r.DB.Delete(&Quotation{}, id).Error
}

func (r *QuotationRepository) ReadAll() ([]interface{}, error) {
    var quotations []Quotation
    if err := r.DB.Find(&quotations).Error; err != nil {
        return nil, err
    }
    var result []interface{}
    for i := range quotations {
        result = append(result, &quotations[i])
    }
    return result, nil
}
