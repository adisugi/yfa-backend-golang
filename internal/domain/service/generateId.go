package service

import "github.com/jinzhu/gorm"

type Result struct {
	Id int
}

func GetMax(db *gorm.DB, tableName string, idName string) *int {

	result := Result{}
	//find max id
	err := db.Table(tableName).Select("MAX(" + idName + ") AS id").Scan(&result).Error

	if err != nil {
		return nil
	}

	//Make a next id
	NewId := result.Id + 1

	return &NewId
}
