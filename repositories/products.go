package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(Id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(Product models.Product) (models.Product, error)
	FindCategoriesById(categoriesId []int) ([]models.Category, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var product []models.Product
	err := r.db.Preload("Category").Find(&product).Error

	return product, err
}

func (r *repository) GetProduct(Id int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, Id).Error

	return product, err
}
func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Where("order_id IS NULL").Create(&product).Error

	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}

func (r *repository) FindCategoriesById(categoriesId []int) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories, categoriesId).Error

	return categories, err
}
