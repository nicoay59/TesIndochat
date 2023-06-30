package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.Customer, error)
	GetUser(ID int) (models.Customer, error)
	CreateUser(user models.Customer) (models.Customer, error)
	UpdateUser(user models.Customer) (models.Customer, error)
	DeleteUser(user models.Customer, ID int) (models.Customer, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.Customer, error) {
	var users []models.Customer
	err := r.db.Preload("Product").Find(&users).Error // add this code

	return users, err
}

func (r *repository) GetUser(ID int) (models.Customer, error) {
	var user models.Customer
	err := r.db.Preload("Product").First(&user, ID).Error // add this code

	return user, err
}

func (r *repository) CreateUser(user models.Customer) (models.Customer, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.Customer) (models.Customer, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.Customer, ID int) (models.Customer, error) {
	err := r.db.Delete(&user, ID).Scan(&user).Error

	return user, err
}
