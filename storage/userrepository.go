package storage

import(
		"gorm.io/gorm"
		. "github.com/rohit123sinha456/payredapp/models"
)

// UserRepository is an example implementation for User
type UserRepository struct {
    DB *gorm.DB
}

func (r *UserRepository) Create(user interface{}) error {
    return r.DB.Create(user).Error
}


func (r *UserRepository) Read(id uint) (interface{}, error) {
    var user User
    if err := r.DB.First(&user, id).Error; err != nil {
        return User{}, err
    }
    return user, nil
}

func (r *UserRepository) Update(user interface{}) error {
    return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
    return r.DB.Delete(&User{}, id).Error
    // return r.DB.Where("id = ?", id).Delete(&user).Error
}

func (r *UserRepository) ReadAll() ([]interface{}, error) {
    var users []User
    if err := r.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    var result []interface{}
    for i := range users {
        result = append(result, &users[i])
    }
    return result, nil
}