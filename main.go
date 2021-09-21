package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"yfa-golang/api"
	"yfa-golang/pkg/config"
	"yfa-golang/pkg/database"
)

func init() {
	config.GetConfig()
}

//func GetTransaksiii(db *gorm.DB) {
//
//	var trans []entity.Transaksi
//	//err := db.Table("transaksis").Joins("JOIN kurirs ON kurirs.id_kurir = transaksis.id_kurir").Find(&trans).Error
//	//err := db.Order("id_transaksi desc").Find(&trans).Error
//	err := db.Preload("Kurir").Find(&trans).Error
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("weww")
//
//	for _, item := range trans {
//		fmt.Println(item)
//	}
//
//}

type ExistId struct {
	Id int
}

func GetExistId(db *gorm.DB, id int) bool {
	existId := ExistId{}
	err := db.Table("kurirs").Select("id_kurir AS id").Where("id_kurir = ?", id).Scan(&existId).Error
	if err != nil {
		return false
	}

	return true
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return
	}

	defer db.Close()

	//fmt.Println("halo wey")
	//GetMaxId(db, "id_kurir", kurir)
	//Id := GetMax(db, "kurirs", "id_kurir")
	//fmt.Println(*Id)

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	app := api.SetupRouter(db)
	app.Run(port)

	//var kur repository.IKurirRepository
	//a := kur.GetExistId(202)
	//fmt.Println(a)

}
