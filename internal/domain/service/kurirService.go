package service

import (
	"github.com/jinzhu/gorm"
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/repository"
)

type KurirService struct {
	kurirRepo repository.IKurirRepository
	db        *gorm.DB
}

type IKurirService interface {
	GetKurir() ([]entity.KurirModelView, error)
	GetOneKurir(int) (*entity.KurirModelView, error)
	SaveKurir(*entity.KurirModelView) (*entity.KurirModelView, error)
}

func NewKurirService(kurirRepo repository.IKurirRepository, db *gorm.DB) *KurirService {
	var kurirService = KurirService{}
	kurirService.kurirRepo = kurirRepo
	kurirService.db = db
	return &kurirService
}

func (s *KurirService) SaveKurir(kurirVM *entity.KurirModelView) (*entity.KurirModelView, error) {

	NewId := GetMax(s.db, "kurirs", "id_kurir")

	kurir := entity.Kurir{
		IdKurir:     *NewId,
		NamaKurir:   kurirVM.NamaKurir,
		NoTelpKurir: kurirVM.NoTelpKurir,
		Alamat:      kurirVM.Alamat,
		Nik:         kurirVM.Nik,
		Ttl:         kurirVM.Ttl,
		File:        kurirVM.File,
		IsDelete:    kurirVM.IsDelete,
	}

	result, err := s.kurirRepo.SaveKurir(&kurir)

	if err != nil {
		return nil, err
	}

	if result != nil {
		kurirVM = &entity.KurirModelView{
			IdKurir:     result.IdKurir,
			NamaKurir:   result.NamaKurir,
			NoTelpKurir: result.NoTelpKurir,
			Alamat:      result.Alamat,
			Nik:         result.Nik,
			Ttl:         result.Ttl,
			File:        result.File,
			IsDelete:    result.IsDelete,
		}
	}

	return kurirVM, nil

}

func (s *KurirService) GetKurir() ([]entity.KurirModelView, error) {
	result, err := s.kurirRepo.GetKurir()
	if err != nil {
		return nil, err
	}

	var kurirs []entity.KurirModelView
	for _, items := range result {
		var kurir entity.KurirModelView
		kurir.IdKurir = items.IdKurir
		kurir.NamaKurir = items.NamaKurir
		kurir.NoTelpKurir = items.NoTelpKurir
		kurir.File = items.File
		kurir.IsDelete = items.IsDelete
		kurirs = append(kurirs, kurir)
	}

	return kurirs, nil
}

func (s *KurirService) GetOneKurir(id int) (*entity.KurirModelView, error) {
	result, err := s.kurirRepo.GetOneKurir(id)

	if err != nil {
		return nil, err
	}

	var kurirVM entity.KurirModelView

	if result != nil {
		kurirVM = entity.KurirModelView{
			IdKurir:     result.IdKurir,
			NamaKurir:   result.NamaKurir,
			NoTelpKurir: result.NoTelpKurir,
			Alamat:      result.Alamat,
			Nik:         result.Nik,
			Ttl:         result.Ttl,
			File:        result.File,
			IsDelete:    result.IsDelete,
		}
	}

	return &kurirVM, nil
}
