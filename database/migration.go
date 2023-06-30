package database

import (
	"fmt"
	"indochat/models"
	"indochat/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Customer{},
		&models.Category{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)

	if err != nil {
		panic("migration Failed 😢😢")
	}

	fmt.Println("migrasi succsesss coyy 😎👍")
}
