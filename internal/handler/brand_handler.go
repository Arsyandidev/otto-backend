package handler

import (
	"be-otto-digital/internal/model"
	"be-otto-digital/internal/service"
	"encoding/json"
	"net/http"
)

type BrandHandler struct {
	brandService service.BrandService
}

func NewBrandHandler(brandService service.BrandService) *BrandHandler {
	return &BrandHandler{brandService: brandService}
}

func (h *BrandHandler) CreateBrands(w http.ResponseWriter, r *http.Request) {
	var brand model.Brand

	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	if err := h.brandService.CreateBrand(&brand); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&brand);
}
