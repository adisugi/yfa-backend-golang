package repository

import (
	"github.com/jinzhu/gorm"
	"yfa-golang/internal/domain/entity"
)

type TransaksiRepository struct {
	db *gorm.DB
}

type ITransaksiRepository interface {
	GetTransaksi() (*[]entity.Transaksi, error)
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	var transRepo = TransaksiRepository{}
	transRepo.db = db
	return &transRepo
}

func (r *TransaksiRepository) GetTransaksi() (*[]entity.Transaksi, error) {
	var trans []entity.Transaksi
	//err := r.db.Order("id_transaksi desc").Find(&trans).Error
	//err := r.db.Joins("JOIN kurirs ON kurirs.id_kurir = transaksis.id_kurir").Find(&trans).Error
	//err := r.db.Table("transaksis").Joins("JOIN kurirs ON kurirs.id_kurir = transaksis.id_kurir").Find(&trans).Error
	err := r.db.Preload("Kurir").Find(&trans).Error
	if err != nil {
		return nil, err
	}

	return &trans, err
}