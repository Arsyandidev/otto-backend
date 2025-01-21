package service

import (
	"be-otto-digital/internal/model"
	"be-otto-digital/internal/repository"
	"errors"
)

type VoucherService interface {
	GetVoucherByID(id uint) ([]model.Voucher, error)
	GetVoucherByBrandID(id uint) ([]model.Voucher, error)
	CreateVoucher(voucher *model.Voucher) error
}

type voucherService struct {
	voucherRepo repository.VoucherRepository
}

func NewVoucherService(voucherRepo repository.VoucherRepository) VoucherService {
	return &voucherService{voucherRepo: voucherRepo}
}

func (s *voucherService) CreateVoucher(voucher *model.Voucher) error {
	if voucher.Code == "" {
		return errors.New("Code tidak boleh kosong")
	}

	if voucher.Description == "" {
		return errors.New("Description tidak boleh kosong")
	}

	if voucher.CostInPoints == 0 {
		return errors.New("CostInPoints tidak boleh nol")
	}

	if voucher.ValidFrom == "" {
		return errors.New("Valid From tidak boleh kosong")
	}

	if voucher.ValidUntil == "" {
		return errors.New("Valid Until tidak boleh kosong")
	}

	return s.voucherRepo.Create(voucher)
}

func (s *voucherService) GetVoucherByID(id uint) ([]model.Voucher, error) {
	return s.voucherRepo.FindByID(id)
}

func (s *voucherService) GetVoucherByBrandID(id uint) ([]model.Voucher, error) {
	return s.voucherRepo.FindByBrandID(id)
}
