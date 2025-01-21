package repository

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	FindByID(id uint) (*model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id uint) error
}

type customerRepository struct {}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (r *customerRepository) Create(customer *model.Customer) error {
	return config.DB.Create(customer).Error
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customer []model.Customer
	err := config.DB.Preload("Transactions").Find(&customer).Error
	return customer, err
}

func (r *customerRepository) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	err := config.DB.Preload("Transactions").First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Update(customer *model.Customer) error {
	return config.DB.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Customer{}, id).Error
}