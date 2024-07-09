package storage

import(
	"gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/models"
)

type ClientRepository struct {
    DB *gorm.DB
}

func (r *ClientRepository) Create(client *Client) error {
    return r.DB.Create(client).Error
}

func (r *ClientRepository) Read(id uint) (*Client, error) {
    var client Client
    if err := r.DB.First(&client, id).Error; err != nil {
        return nil, err
    }
    return &client, nil
}

func (r *ClientRepository) ReadAll() ([]Client, error) {
    var clients []Client
    if err := r.DB.Find(&clients).Error; err != nil {
        return nil, err
    }
    return clients, nil
}

func (r *ClientRepository) Update(client *Client) error {
    return r.DB.Save(client).Error
}

func (r *ClientRepository) Delete(id uint) error {
    return r.DB.Delete(&Client{}, id).Error
}

func (r *ClientRepository) FindByEmail(email string) (*Client, error) {
    var client Client
    if err := r.DB.Where("login_email = ? OR JSON_CONTAINS(emails, ?)", email, `"`+email+`"`).First(&client).Error; err != nil {
        return nil, err
    }
    return &client, nil
}
