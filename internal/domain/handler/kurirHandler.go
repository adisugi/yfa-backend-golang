package handler

import (
	"fmt"
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

func (h *KurirHandler) SaveKurir(c *gin.Context) {
	//IdKurir: kurirVM.IdKurir,
	//	NamaKurir: kurirVM.NamaKurir,
	//		NoTelpKurir: kurirVM.NoTelpKurir,
	//		Alamat: kurirVM.Alamat,
	//		Nik: kurirVM.Nik,
	//		Ttl: kurirVM.Ttl,
	//		File: kurirVM.File,
	//		IsDelete: kurirVM.IsDelete,
	//idKurir := c.PostForm("idKurir")
	namaKurir := c.PostForm("namaKurir")
	noTelpKurir := c.PostForm("noTelpKurir")
	alamat := c.PostForm("alamat")
	nik := c.PostForm("nik")
	ttl := c.PostForm("ttl")
	file := c.PostForm("file")
	isDelete := c.DefaultPostForm("isDelete", "0")
	fmt.Print(isDelete)

	kurir := entity.KurirModelView{
		NamaKurir:   namaKurir,
		NoTelpKurir: noTelpKurir,
		Alamat:      alamat,
		Nik:         nik,
		Ttl:         ttl,
		File:        file,
		//IsDelete: isDelete,
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
