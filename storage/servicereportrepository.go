package storage

import(
		"gorm.io/gorm"
		. "github.com/rohit123sinha456/payredapp/models"
)

type ServiceReportRepository struct {
    DB *gorm.DB
}

func (r *ServiceReportRepository) Create(serviceReport interface{}) error {
    return r.DB.Create(serviceReport).Error
}

func (r *ServiceReportRepository) Read(id uint) (interface{}, error) {
    var serviceReport ServiceReport
    if err := r.DB.First(&serviceReport, id).Error; err != nil {
        return nil, err
    }
    return &serviceReport, nil
}

func (r *ServiceReportRepository) Update(serviceReport interface{}) error {
    return r.DB.Save(serviceReport).Error
}

func (r *ServiceReportRepository) Delete(id uint) error {
    return r.DB.Delete(&ServiceReport{}, id).Error
}

func (r *ServiceReportRepository) ReadAll() ([]interface{}, error) {
    var serviceReports []ServiceReport
    if err := r.DB.Find(&serviceReports).Error; err != nil {
        return nil, err
    }
    var result []interface{}
    for i := range serviceReports {
        result = append(result, &serviceReports[i])
    }
    return result, nil
}
