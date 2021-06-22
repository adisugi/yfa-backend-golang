package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"yfa-golang/api/middleware"
	"yfa-golang/internal/domain/handler"
	"yfa-golang/internal/domain/repository"
	"yfa-golang/internal/domain/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	transrepo := repository.NewTransaksiRepository(db)
	transService := service.NewTransaksiService(transrepo)
	transHandler := handler.NewTransaksiHandler(transService)

	kurirrepo := repository.NewKurirRepository(db)
	kurirService := service.NewKurirService(kurirrepo)
	kurirHandler := handler.NewKurirHandler(kurirService)

	trans := router.Group("/transaksi")
	{
		trans.GET("/", middleware.CORSMiddleware(), transHandler.GetTransaksi)
	}

	kurirs := router.Group("/kurir")
	{
		kurirs.GET("/", middleware.CORSMiddleware(), kurirHandler.GetKurir)
		kurirs.GET("/:id_kurir", middleware.CORSMiddleware(), kurirHandler.GetOneKurir)

	}

	return router

}
