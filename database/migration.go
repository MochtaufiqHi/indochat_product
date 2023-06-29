package database

import (
	"diskon/models"
	"diskon/pkg/mysql"
	"fmt"
)

func MigrateDB() {
	err := mysql.DB.AutoMigrate(
		// &models.Products{},
		&models.Product{},
		&models.Category{},
		&models.Customer{},
		&models.Order{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}
	fmt.Println("Migration Success")
}
