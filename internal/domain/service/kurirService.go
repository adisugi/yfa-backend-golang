package service

import (
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/repository"
)

type KurirService struct {
	kurirRepo repository.IKurirRepository
}

type IKurirService interface {
	GetKurir() ([]entity.KurirModelView, error)
	GetOneKurir(int) (*entity.KurirModelView, error)
}

func NewKurirService(kurirRepo repository.IKurirRepository) *KurirService {
	var kurirService = KurirService{}
	kurirService.kurirRepo = kurirRepo
	return &kurirService
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

func (s * KurirService) GetOneKurir(id int) (*entity.KurirModelView, error) {
	result, err := s.kurirRepo.GetOneKurir(id)

	if err != nil {
		return nil, err
	}

	var kurirVM entity.KurirModelView

	if result != nil {
		kurirVM = entity.KurirModelView{
			IdKurir : result.IdKurir,
			NamaKurir : result.NamaKurir,
			NoTelpKurir : result.NoTelpKurir,
			Alamat: result.Alamat,
			Nik: result.Nik,
			Ttl: result.Ttl,
			File : result.File,
			IsDelete : result.IsDelete,
		}
	}

	return &kurirVM, nil
}