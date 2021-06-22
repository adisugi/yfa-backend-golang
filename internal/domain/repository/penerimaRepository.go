package repository

import (
	"github.com/jinzhu/gorm"
	"yfa-golang/internal/domain/entity"
)

type PenerimaRepository struct {
	db *gorm.DB
}

type IPenerimaRepository interface {
	GetPenerima() ([]entity.Penerima, error)
}

func NewPenerimaRepository(db *gorm.DB) *PenerimaRepository {
	var penerimaRepo = PenerimaRepository{}
	penerimaRepo.db = db
	return &penerimaRepo
}

func (r *PenerimaRepository) GetPenerima() ([]entity.Penerima, error) {
	var penerima []entity.Penerima
	err := r.db.Order("id_penerima desc").Find(&penerima).Error
	if err != nil {
		return nil, err
	}

	return penerima, err
}
