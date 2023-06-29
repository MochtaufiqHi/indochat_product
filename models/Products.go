package models

import "time"

type Product struct {
	ID           int      `json:"id_product" gorm:"primary_key:auto_increment"`
	Name         string   `json:"name" gorm:"type: varchar(255)"`
	Price        int      `json:"price" gorm:"type: varchar(255)"`
	Description  string   `json:"description" grom:"type: varchar(255)"`
	Image        string   `json:"image" gorm:"type: varchar(255)"`
	CategoriesID int      `json:"category_ID"`
	Categories   Category `json:"category" gorm:"foreignKey:CategoriesID"`
	Orders       []Order  `gorm:"many2many:order_products"`
}

type Category struct {
	ID   int    `json:"id_category" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}

type Order struct {
	ID         int       `json:"id_order" gorm:"primary_key:auto_increment"`
	CustomerID int       `json:"customer_id"`
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID"`
	Product    []Product `json:"product" gorm:"product"`
	Date       time.Time `gorm:"date"`
	Status     string    `gorm:"status"`
	Products   []Product `gorm:"many2many:order_products"`
}

type Customer struct {
	ID       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}

type ResponCategory struct {
	Name string
}

func (ResponCategory) TableName() string {
	return "categories"
}

type ResponCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ResponCustomer) TableName() string {
	return "customers"
}

type ResponProducts struct {
	Name         string     `json:"name"`
	Price        int        `json:"price"`
	Description  string     `json:"description"`
	Image        string     `json:"image"`
	CategoriesID int        `json:"category_id"`
	Categories   []Category `json:"categories"`
}

func (ResponProducts) TableName() string {
	return "products"
}
