package storage

import(
		"gorm.io/gorm"
		"github.com/rohit123sinha456/payredapp/models"
)

type ServiceReportRepository struct {
    DB *gorm.DB
}

func (r *ServiceReportRepository) Create(serviceReport *models.ServiceReport) error {
    var client models.Client
    if err := r.DB.First(&client, serviceReport.ClientID).Error; err != nil {
        return err
    }
    serviceReport.Client = client
    return r.DB.Create(serviceReport).Error
}

func (r *ServiceReportRepository) Read(id uint) (interface{}, error) {
    var serviceReport models.ServiceReport
    if err := r.DB.Preload("Client").First(&serviceReport, id).Error; err != nil {
        return nil, err
    }
    return &serviceReport, nil
}

func (r *ServiceReportRepository) Update(serviceReport models.ServiceReport) error {
    var client models.Client
    if err := r.DB.First(&client, serviceReport.ClientID).Error; err != nil {
        return err
    }
    serviceReport.Client = client
    return r.DB.Save(serviceReport).Error
}

func (r *ServiceReportRepository) Delete(id uint) error {
    return r.DB.Delete(&models.ServiceReport{}, id).Error
}

func (r *ServiceReportRepository) ReadAll() ([]models.ServiceReport, error) {
    var serviceReports []models.ServiceReport
    if err := r.DB.Find(&serviceReports).Error; err != nil {
        return nil, err
    }
    var result []models.ServiceReport
    for i := range serviceReports {
        result = append(result, serviceReports[i])
    }
    return result, nil
}

func (r *ServiceReportRepository) GetServiceReportsByClientID(clientID uint) ([]models.ServiceReport, error) {
    var serviceReports []models.ServiceReport
    if err := r.DB.Preload("Client").Where("client_id = ?", clientID).Order("service_report_date DESC").Find(&serviceReports).Error; err != nil {
        return nil, err
    }
    return serviceReports, nil
}