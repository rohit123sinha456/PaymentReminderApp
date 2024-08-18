package storage

import(
		"gorm.io/gorm"
		"github.com/rohit123sinha456/payredapp/models"
)

type PaymentRepository struct {
    DB *gorm.DB
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
    var client models.Client
    if err := r.DB.First(&client, payment.ClientID).Error; err != nil {
        return err
    }
    payment.Client = client

    // Create payment
    return r.DB.Create(payment).Error
}

func (r *PaymentRepository) Read(id uint) (interface{}, error) {
    var payment models.Payment
    if err := r.DB.Preload("Client").First(&payment, id).Error; err != nil {
        return nil, err
    }
    return &payment, nil
}

func (r *PaymentRepository) Update(payment models.Payment) error {
    var client models.Client
    if err := r.DB.First(&client, payment.ClientID).Error; err != nil {
        return err
    }
    payment.Client = client
    return r.DB.Save(payment).Error
}

func (r *PaymentRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Payment{}, id).Error
}

func (r *PaymentRepository) ReadAll() ([]models.Payment, error) {
    var payments []models.Payment
    if err := r.DB.Preload("Client").Find(&payments).Error; err != nil {
        return nil, err
    }
    var result []models.Payment
    for i := range payments {
        result = append(result, payments[i])
    }
    return result, nil
}

func (r *PaymentRepository) GetPaymentsByClientID(clientID uint) ([]models.Payment, error) {
    var payments []models.Payment
    if err := r.DB.Preload("Client").Where("client_id = ?", clientID).Order("is_paid ASC, invoice_date DESC").Find(&payments).Error; err != nil {
        return nil, err
    }
    return payments, nil
}