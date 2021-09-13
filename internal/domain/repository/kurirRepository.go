package repository

import (
	"github.com/jinzhu/gorm"
	"yfa-golang/internal/domain/entity"
)

type KurirRepository struct {
	db *gorm.DB
}

type IKurirRepository interface {
	GetKurir() ([]entity.Kurir, error)
	GetOneKurir(int) (*entity.Kurir, error)
	SaveKurir(*entity.Kurir) (*entity.Kurir, error)
}

func NewKurirRepository(db *gorm.DB) *KurirRepository {
	var kurirRepo = KurirRepository{}
	kurirRepo.db = db
	return &kurirRepo
}

func (r *KurirRepository) SaveKurir(kurir *entity.Kurir) (*entity.Kurir, error) {
	err := r.db.Create(&kurir).Error
	if err != nil {
		return nil, err
	}

	return kurir, nil
}

func (r *KurirRepository) GetKurir() ([]entity.Kurir, error) {
	var kurir []entity.Kurir
	err := r.db.Order("id_kurir desc").Find(&kurir).Error
	if err != nil {
		return nil, err
	}
	return kurir, err
}

func (r *KurirRepository) GetOneKurir(id int) (*entity.Kurir, error) {
	var kurir entity.Kurir
	err := r.db.Where("id_kurir = ?", id).Find(&kurir).Error
	if err != nil {
		return nil, err
	}
	return &kurir, err
}
