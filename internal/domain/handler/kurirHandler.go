package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"path/filepath"
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

func (h *KurirHandler) SaveKurir(c *gin.Context) {
	namaKurir := c.PostForm("namaKurir")
	noTelpKurir := c.PostForm("noTelpKurir")
	alamat := c.PostForm("alamat")
	nik := c.PostForm("nik")
	ttl := c.PostForm("ttl")
	file, err := c.FormFile("file")

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	path := viper.GetString("Files.KurirPath")
	fileName := filepath.Base(file.Filename)

	if er := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, fileName)); er != nil {
		response.ResponseError(c, er.Error(), http.StatusBadRequest)
		return
	}

	kurir := entity.KurirModelView{
		NamaKurir:   namaKurir,
		NoTelpKurir: noTelpKurir,
		Alamat:      alamat,
		Nik:         nik,
		Ttl:         ttl,
		File:        fileName,
	}

	result, err := h.kurirService.SaveKurir(&kurir)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseCreated(c, result)
}

func (h *KurirHandler) UpdateKurir(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	existId := h.kurirService.GetExistId(kurirId)
	if existId == false {
		response.ResponseError(c, "Id tidak ditemukan", http.StatusNotFound)
		return
	}

	namaKurir := c.PostForm("namaKurir")
	noTelpKurir := c.PostForm("noTelpKurir")
	alamat := c.PostForm("alamat")
	nik := c.PostForm("nik")
	ttl := c.PostForm("ttl")
	file, err := c.FormFile("file")

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	path := viper.GetString("Files.KurirPath")
	fileName := filepath.Base(file.Filename)

	if er := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, fileName)); er != nil {
		response.ResponseError(c, er.Error(), http.StatusBadRequest)
		return
	}

	kurir := entity.KurirModelView{
		IdKurir:     kurirId,
		NamaKurir:   namaKurir,
		NoTelpKurir: noTelpKurir,
		Alamat:      alamat,
		Nik:         nik,
		Ttl:         ttl,
		File:        fileName,
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

	existId := h.kurirService.GetExistId(kurirId)
	if existId == false {
		response.ResponseError(c, "Id tidak ditemukan", http.StatusNotFound)
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

	existId := h.kurirService.GetExistId(kurirId)
	if existId == false {
		response.ResponseError(c, "Id tidak ditemukan", http.StatusNotFound)
		return
	}

	er := h.kurirService.SDeleteKurir(kurirId)
	if er != nil {
		response.ResponseError(c, er.Error(), http.StatusInternalServerError)
	}

	response.ResponseOk(c, "Data berhasil dihapus")

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

func (h *KurirHandler) GetFile(c *gin.Context) {
	kurirId, err := strconv.Atoi(c.Param("id_kurir"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

	result, err := h.kurirService.GetFile(kurirId)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseOkWithData(c, result)
}
