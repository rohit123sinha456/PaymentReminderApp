package storage

import(
		"gorm.io/gorm"
		. "github.com/rohit123sinha456/payredapp/models"
)

type NOCRepository struct {
    DB *gorm.DB
}

func (r *NOCRepository) Create(noc interface{}) error {
    return r.DB.Create(noc).Error
}

func (r *NOCRepository) Read(id uint) (interface{}, error) {
    var noc NOC
    if err := r.DB.First(&noc, id).Error; err != nil {
        return nil, err
    }
    return &noc, nil
}

func (r *NOCRepository) Update(noc interface{}) error {
    return r.DB.Save(noc).Error
}

func (r *NOCRepository) Delete(id uint) error {
    return r.DB.Delete(&NOC{}, id).Error
}

func (r *NOCRepository) ReadAll() ([]interface{}, error) {
    var nocs []NOC
    if err := r.DB.Find(&nocs).Error; err != nil {
        return nil, err
    }
    var result []interface{}
    for i := range nocs {
        result = append(result, &nocs[i])
    }
    return result, nil
}
