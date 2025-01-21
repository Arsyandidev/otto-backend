package repository

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
)

type BrandRepository interface {
	Create(brand *model.Brand) error
}

type brandRepository struct {}

func NewBrandRepository() BrandRepository {
	return &brandRepository{}
}

func (r *brandRepository) Create(brand *model.Brand) error {
	return config.DB.Create(brand).Error
}