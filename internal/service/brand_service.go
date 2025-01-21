package service

import (
	"be-otto-digital/internal/model"
	"be-otto-digital/internal/repository"
	"errors"
)

type BrandService interface {
	CreateBrand(brand *model.Brand) error
}

type brandService struct {
	brandRepo repository.BrandRepository
}

func NewBrandService(brandRepo repository.BrandRepository) BrandService {
	return &brandService{brandRepo: brandRepo}
}

func (s *brandService) CreateBrand(brand *model.Brand) error {
	if brand.Name == "" {
		return errors.New("Nama brand tidak boleh kosong")
	}

	if brand.Description == "" {
		return errors.New("Deskripsi brand tidak boleh kosong")
	}

	return s.brandRepo.Create(brand)
}
