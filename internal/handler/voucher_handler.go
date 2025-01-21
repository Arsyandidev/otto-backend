package handler

import (
	"be-otto-digital/internal/model"
	"be-otto-digital/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type VoucherHandler struct {
	voucherService service.VoucherService
}

func NewVoucherHandler(voucherService service.VoucherService) *VoucherHandler {
	return &VoucherHandler{voucherService: voucherService}
}

func (h *VoucherHandler) CreateVoucher(w http.ResponseWriter, r *http.Request) {
	var voucher model.Voucher

	if err := json.NewDecoder(r.Body).Decode(&voucher); err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	if err := h.voucherService.CreateVoucher(&voucher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&voucher);
}

func (h *VoucherHandler) GetVoucherByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
        http.Error(w, "Invalid brand ID", http.StatusBadRequest)
        return
    }
	
	voucher, err := h.voucherService.GetVoucherByID(uint(id))
	if err != nil {
		http.Error(w, "Voucher not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&voucher)
}

func (h *VoucherHandler) GetVoucherByBrandID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
        http.Error(w, "Invalid brand ID", http.StatusBadRequest)
        return
    }
	
	voucher, err := h.voucherService.GetVoucherByBrandID(uint(id))
	if err != nil {
		http.Error(w, "Voucher not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&voucher)
}