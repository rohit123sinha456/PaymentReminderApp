package storage

import(
		"gorm.io/gorm"
		. "github.com/rohit123sinha456/payredapp/models"
)

type PaymentRepository struct {
    DB *gorm.DB
}

func (r *PaymentRepository) Create(payment interface{}) error {
    return r.DB.Create(payment).Error
}

func (r *PaymentRepository) Read(id uint) (interface{}, error) {
    var payment Payment
    if err := r.DB.First(&payment, id).Error; err != nil {
        return nil, err
    }
    return &payment, nil
}

func (r *PaymentRepository) Update(payment interface{}) error {
    return r.DB.Save(payment).Error
}

func (r *PaymentRepository) Delete(id uint) error {
    return r.DB.Delete(&Payment{}, id).Error
}

func (r *PaymentRepository) ReadAll() ([]interface{}, error) {
    var payments []Payment
    if err := r.DB.Find(&payments).Error; err != nil {
        return nil, err
    }
    var result []interface{}
    for i := range payments {
        result = append(result, &payments[i])
    }
    return result, nil
}
