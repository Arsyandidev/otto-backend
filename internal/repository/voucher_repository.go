package repository

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
)

type VoucherRepository interface {
	FindByID(id uint) ([]model.Voucher, error)
	FindByBrandID(id uint) ([]model.Voucher, error)
	Create(voucher *model.Voucher) error
}

type voucherRepository struct {}

func NewVoucherRepository() VoucherRepository {
	return &voucherRepository{}
}

func (r *voucherRepository) Create(voucher *model.Voucher) error {
	return config.DB.Create(voucher).Error
}

func (r *voucherRepository) FindByID(id uint) ([]model.Voucher, error) {
	var voucher []model.Voucher
	err := config.DB.Preload("Brand").Where("id = ?", id).Find(&voucher).Error
	return voucher, err
}

func (r *voucherRepository) FindByBrandID(id uint) ([]model.Voucher, error) {
	var voucher []model.Voucher
	err := config.DB.Preload("Brand").Where("brand_id = ?", id).Find(&voucher).Error
	return voucher, err
}
