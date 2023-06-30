package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	FindOrders(userId int) ([]models.Order, error)
	GetOrders(Id int) (models.Order, error)
	CreateOrders(orders models.Order) (models.Order, error)
	DeleteOrders(orders models.Order) (models.Order, error)
	GetOrderByUSer(Id int) ([]models.Order, error)
}

func RepositoryOrders(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders(userId int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Product.Category").Preload("Customer").Find(&orders).Error

	return orders, err
}

func (r *repository) GetOrders(Id int) (models.Order, error) {
	var orders models.Order
	err := r.db.Preload("Product").Preload("Customer").First(&orders, Id).Error

	return orders, err
}

func (r *repository) GetOrderByUSer(ID int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("costumer_id=?", ID).Preload("Customer").Preload("Product").Find(&orders).Error

	return orders, err

}

func (r *repository) CreateOrders(orders models.Order) (models.Order, error) {
	err := r.db.Create(&orders).Error

	return orders, err
}

func (r *repository) DeleteOrders(orders models.Order) (models.Order, error) {
	err := r.db.Delete(&orders).Error

	return orders, err
}
