package contact

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]ContactData, error)
	FindById(ID string) (ContactData, error)
	Create(contact ContactData) (ContactData, error)
	Delete(ID string) (ContactData, error)
	Update(contact ContactData) (ContactData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]ContactData, error) {
	var contact []ContactData
	err := r.db.Find(&contact).Error
	return contact, err
}

func (r *repository) FindById(ID string) (ContactData, error) {
	var contact ContactData
	err := r.db.First(&contact, "id = ?", ID).Error
	return contact, err
}
func (r *repository) Create(contact ContactData) (ContactData, error) {
	err := r.db.Create(&contact).Error
	return contact, err
}

func (r *repository) Delete(ID string) (ContactData, error) {
	var contact ContactData
	err := r.db.Where("id = ?", ID).Delete(&contact).Error
	return contact, err
}

func (r *repository) Update(contact ContactData) (ContactData, error) {
	err := r.db.Save(&contact).Error
	return contact, err
}
