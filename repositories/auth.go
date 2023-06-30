package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.Customer) (models.Customer, error)
	Login(email string) (models.Customer, error)
	CheckAuth(ID int) (models.Customer, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.Customer) (models.Customer, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.Customer, error) {
	var user models.Customer
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckAuth(ID int) (models.Customer, error) {
	var user models.Customer
	err := r.db.Preload("Products").First(&user, ID).Error // add this code

	return user, err
}
