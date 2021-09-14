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

func (h *KurirHandler) UpdateKurir(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	namaKurir := c.PostForm("namaKurir")
	noTelpKurir := c.PostForm("noTelpKurir")
	alamat := c.PostForm("alamat")
	nik := c.PostForm("nik")
	ttl := c.PostForm("ttl")
	file := c.PostForm("file")

	kurir := entity.KurirModelView{
		IdKurir:     kurirId,
		NamaKurir:   namaKurir,
		NoTelpKurir: noTelpKurir,
		Alamat:      alamat,
		Nik:         nik,
		Ttl:         ttl,
		File:        file,
	}

	result, err := h.kurirService.UpdateKurir(&kurir)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, result)

}

func (h *KurirHandler) HDeleteKurir(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	er := h.kurirService.HDeleteKurir(kurirId)
	if er != nil {
		response.ResponseError(c, er.Error(), http.StatusInternalServerError)
	}

	response.ResponseOk(c, "Data berhasil dihapus")

}

func (h *KurirHandler) SDeleteKurir(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	er := h.kurirService.SDeleteKurir(kurirId)
	if er != nil {
		response.ResponseError(c, er.Error(), http.StatusInternalServerError)
	}

	response.ResponseOk(c, "Data berhasil dihapus")

}

func (h *KurirHandler) SaveKurir(c *gin.Context) {
	namaKurir := c.PostForm("namaKurir")
	noTelpKurir := c.PostForm("noTelpKurir")
	alamat := c.PostForm("alamat")
	nik := c.PostForm("nik")
	ttl := c.PostForm("ttl")
	file := c.PostForm("file")

	kurir := entity.KurirModelView{
		NamaKurir:   namaKurir,
		NoTelpKurir: noTelpKurir,
		Alamat:      alamat,
		Nik:         nik,
		Ttl:         ttl,
		File:        file,
	}

	result, err := h.kurirService.SaveKurir(&kurir)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseCreated(c, result)
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
