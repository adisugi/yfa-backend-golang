package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/service"
	"yfa-golang/pkg/response"
)

type KurirHandler struct {
	kurirService service.IKurirService
}

func NewKurirHandler(kurirService service.IKurirService) *KurirHandler {
	var kurirHandler = KurirHandler{}
	kurirHandler.kurirService = kurirService
	return &kurirHandler
}

func (h *KurirHandler) GetKurir(c *gin.Context) {
	result, err := h.kurirService.GetKurir()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.KurirModelView{}
	}

	response.ResponseOkWithData(c, result)
}

func (h *KurirHandler) GetOneKurir(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

	result, err := h.kurirService.GetOneKurir(kurirId)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseOkWithData(c, result)

}