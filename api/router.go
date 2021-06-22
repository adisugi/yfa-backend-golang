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

	//transaksi
	transrepo := repository.NewTransaksiRepository(db)
	transService := service.NewTransaksiService(transrepo)
	transHandler := handler.NewTransaksiHandler(transService)
	trans := router.Group("/transaksi")
	{
		trans.GET("/", middleware.CORSMiddleware(), transHandler.GetTransaksi)
		trans.GET("/resi/:resi", middleware.CORSMiddleware(), transHandler.GetResi)
	}

	//kurir
	kurirrepo := repository.NewKurirRepository(db)
	kurirService := service.NewKurirService(kurirrepo)
	kurirHandler := handler.NewKurirHandler(kurirService)
	kurirs := router.Group("/kurir")
	{
		kurirs.GET("/", middleware.CORSMiddleware(), kurirHandler.GetKurir)
		kurirs.GET("/id/:id_kurir", middleware.CORSMiddleware(), kurirHandler.GetOneKurir)
	}

	//penerima
	penerimaRepo := repository.NewPenerimaRepository(db)
	penerimaService := service.NewPenerimaService(penerimaRepo)
	penerimaHandler := handler.NewPenerimaHandler(penerimaService)
	penerimas := router.Group("/penerima")
	{
		penerimas.GET("/", middleware.CORSMiddleware(), penerimaHandler.GetPenerima)
	}

	//pengirim
	pengirimRepo := repository.NewPengirimRepository(db)
	pengirimService := service.NewPengirimService(pengirimRepo)
	pengirimHandler := handler.NewPengirimHandler(pengirimService)
	pengirims := router.Group("/pengirim")
	{
		pengirims.GET("/", middleware.CORSMiddleware(), pengirimHandler.GetPengirim)
	}

	return router
}
