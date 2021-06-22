package repository

import (
	"github.com/jinzhu/gorm"
	"yfa-golang/internal/domain/entity"
)

type PengirimRepository struct {
	db *gorm.DB
}

type IPengirimRepository interface {
	GetPengirim() ([]entity.Pengirim, error)
}

func NewPengirimRepository(db *gorm.DB) *PengirimRepository {
	var pengirimRepo = PengirimRepository{}
	pengirimRepo.db = db
	return &pengirimRepo
}

func (r *PengirimRepository) GetPengirim() ([]entity.Pengirim, error) {
	var pengirim []entity.Pengirim
	err := r.db.Order("id_pengirim desc").Find(&pengirim).Error
	if err != nil {
		return nil, err
	}

	return pengirim, err
}
