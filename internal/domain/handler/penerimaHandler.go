package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/service"
	"yfa-golang/pkg/response"
)

type PenerimaHandler struct {
	penerimaService service.IPenerimaService
}

func NewPenerimaHandler(penerimaService service.IPenerimaService) *PenerimaHandler {
	var penerimaHandler = PenerimaHandler{}
	penerimaHandler.penerimaService = penerimaService
	return &penerimaHandler
}

func (h *PenerimaHandler) GetPenerima(c *gin.Context) {
	result, err := h.penerimaService.GetPenerima()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.PenerimaModelView{}
	}

	response.ResponseOkWithData(c, result)
}
