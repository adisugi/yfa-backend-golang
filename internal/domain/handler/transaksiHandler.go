package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/service"
	"yfa-golang/pkg/response"
)

type TransaksiHandler struct {
	transaksiService service.ITransaksiService
}

func NewTransaksiHandler(transaksiSevice service.ITransaksiService) *TransaksiHandler {
	var transaksiHandler = TransaksiHandler{}
	transaksiHandler.transaksiService = transaksiSevice
	return &transaksiHandler
}

func (h *TransaksiHandler) GetTransaksi(c *gin.Context) {
	result, err := h.transaksiService.GetTransaksi()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &[]entity.TransaksiModelView{}
	}

	response.ResponseOkWithData(c, result)
}

func (h *TransaksiHandler) GetResi(c *gin.Context) {
	resi := c.Param("resi")

	result, err := h.transaksiService.GetResi(resi)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseOkWithData(c, result)
}
