package repository

import (
	"diskon/models"

	"gorm.io/gorm"
)

type CusRepo interface {
	GetCus() ([]models.Customer, error)
}

type CateRepo interface {
	GetCat() ([]models.Category, error)
}

type ProRepo interface {
	GetPro() ([]models.Product, error)
}

type OrRepo interface {
	GetOr() ([]models.Order, error)
}

func RepoConnCus(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCus() ([]models.Customer, error) {
	var cus []models.Customer
	err := r.db.Find(&cus).Error
	return cus, err
}

func (r *repository) GetCat() ([]models.Category, error) {
	var cat []models.Category
	err := r.db.Find(&cat).Error
	return cat, err
}

func (r *repository) GetPro() ([]models.Product, error) {
	var pro []models.Product
	err := r.db.Preload("Category").Find(&pro).Error
	return pro, err
}

func (r *repository) GetOr() ([]models.Order, error) {
	var or []models.Order
	err := r.db.Preload("Product.Category").Preload("Customer").Find(&or).Error
	return or, err
}
