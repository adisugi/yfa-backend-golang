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
	UpdateKurir(*entity.Kurir) (*entity.Kurir, error)
	HDeleteKurir(int) error
	SDeleteKurir(*entity.Kurir) error
	GetExistId(id int) (bool, error)
	GetFile(id int) (*string, error)
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

func (r *KurirRepository) UpdateKurir(kurir *entity.Kurir) (*entity.Kurir, error) {
	err := r.db.Save(&kurir).Error
	if err != nil {
		return nil, err
	}

	return kurir, err
}

func (r *KurirRepository) HDeleteKurir(id int) error {
	var kurir entity.Kurir
	err := r.db.Where("id_kurir = ?", id).Delete(&kurir).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *KurirRepository) SDeleteKurir(kurir *entity.Kurir) error {
	err := r.db.Save(&kurir).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *KurirRepository) GetKurir() ([]entity.Kurir, error) {
	var kurir []entity.Kurir
	err := r.db.Where("is_delete = ?", 0).Order("id_kurir desc").Find(&kurir).Error
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

func (r *KurirRepository) GetExistId(id int) (bool, error) {
	type ExistId struct {
		Id int
	}

	existId := ExistId{}
	err := r.db.Table("kurirs").Select("id_kurir AS id").Where("id_kurir = ?", id).Scan(&existId).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *KurirRepository) GetFile(id int) (*string, error) {
	type kurirFile struct {
		File string
	}

	file := kurirFile{}
	err := r.db.Table("kurirs").Select("file").Where("id_kurir = ?", id).Scan(&file).Error
	if err != nil {
		return nil, err
	}

	return &file.File, err
}
