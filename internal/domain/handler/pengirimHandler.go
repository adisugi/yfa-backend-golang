package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/service"
	"yfa-golang/pkg/response"
)

type PengirimHandler struct {
	pengirimService service.IPengirimService
}

func NewPengirimHandler(pengirimService service.IPengirimService) *PengirimHandler {
	var pengirimHandler = PengirimHandler{}
	pengirimHandler.pengirimService = pengirimService
	return &pengirimHandler
}

func (h *PengirimHandler) GetPengirim(c *gin.Context) {
	result, err := h.pengirimService.GetPengirim()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.PengirimModelView{}
	}

	response.ResponseOkWithData(c, result)
}
