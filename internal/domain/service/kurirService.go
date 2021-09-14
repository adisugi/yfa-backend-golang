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
	UpdateKurir(*entity.KurirModelView) (*entity.KurirModelView, error)
	HDeleteKurir(int) error
	SDeleteKurir(int) error
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
		IsDelete:    0,
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

func (s *KurirService) UpdateKurir(kurirVm *entity.KurirModelView) (*entity.KurirModelView, error) {
	kurir := entity.Kurir{
		IdKurir:     kurirVm.IdKurir,
		NamaKurir:   kurirVm.NamaKurir,
		NoTelpKurir: kurirVm.NoTelpKurir,
		Alamat:      kurirVm.Alamat,
		Nik:         kurirVm.Nik,
		Ttl:         kurirVm.Ttl,
		File:        kurirVm.File,
	}

	_, err := s.kurirRepo.UpdateKurir(&kurir)
	if err != nil {
		return nil, err
	}

	return kurirVm, err
}

func (s *KurirService) HDeleteKurir(id int) error {
	err := s.kurirRepo.HDeleteKurir(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *KurirService) SDeleteKurir(id int) error {
	result, err := s.kurirRepo.GetOneKurir(id)

	if err != nil {
		return err
	}

	var kurir entity.Kurir

	if result != nil {
		kurir = entity.Kurir{
			IdKurir:     id,
			NamaKurir:   result.NamaKurir,
			NoTelpKurir: result.NoTelpKurir,
			Alamat:      result.Alamat,
			Nik:         result.Nik,
			Ttl:         result.Ttl,
			File:        result.File,
			IsDelete:    1,
		}
	}

	er := s.kurirRepo.SDeleteKurir(&kurir)
	if er != nil {
		return er
	}

	return nil
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
