package routes

import (
	"be-otto-digital/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, brandHandler *handler.BrandHandler, voucherHandler *handler.VoucherHandler, transactionHandler *handler.TransactionHandler) {
	r.HandleFunc("/brand", brandHandler.CreateBrands).Methods("POST")

	r.HandleFunc("/voucher", voucherHandler.CreateVoucher).Methods("POST")
	r.HandleFunc("/voucher", voucherHandler.GetVoucherByID).Methods("GET")
	r.HandleFunc("/voucher/brand", voucherHandler.GetVoucherByBrandID).Methods("GET")

	r.HandleFunc("/transaction/redemption", transactionHandler.CreateTransaction).Methods("POST")
}