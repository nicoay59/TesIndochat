package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindDiscounts() ([]models.Discount, error)
	GetDiscount(ID int) (models.Discount, error)
	CreateDiscount(discount models.Discount) (models.Discount, error)
	UpdateDiscount(discount models.Discount) (models.Discount, error)
	DeleteDiscount(discount models.Discount, ID int) (models.Discount, error)
}

func RepositoryDiscount(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDiscounts() ([]models.Discount, error) {
	var discounts []models.Discount
	err := r.db.Find(&discounts).Error

	return discounts, err
}

func (r *repository) GetDiscount(ID int) (models.Discount, error) {
	var discount models.Discount
	err := r.db.First(&discount, ID).Error

	return discount, err
}

func (r *repository) CreateDiscount(discount models.Discount) (models.Discount, error) {
	err := r.db.Create(&discount).Error

	return discount, err
}

func (r *repository) UpdateDiscount(discount models.Discount) (models.Discount, error) {
	err := r.db.Save(&discount).Error

	return discount, err
}

func (r *repository) DeleteDiscount(discount models.Discount, ID int) (models.Discount, error) {
	err := r.db.Delete(&discount, ID).Scan(&discount).Error

	return discount, err
}
