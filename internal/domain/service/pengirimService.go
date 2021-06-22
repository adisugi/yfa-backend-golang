package service

import (
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/repository"
)

type PengirimService struct {
	pengirimRepo repository.IPengirimRepository
}

type IPengirimService interface {
	GetPengirim() ([]entity.PengirimModelView, error)
}

func NewPengirimService(pengirimRepo repository.IPengirimRepository) *PengirimService {
	var pengirimService = PengirimService{}
	pengirimService.pengirimRepo = pengirimRepo
	return &pengirimService
}

func (s *PengirimService) GetPengirim() ([]entity.PengirimModelView, error) {
	result, err := s.pengirimRepo.GetPengirim()
	if err != nil {
		return nil, err
	}

	var pengirims []entity.PengirimModelView
	for _, items := range result {
		var pengirim entity.PengirimModelView
		pengirim.IdPengirim = items.IdPengirim
		pengirim.NamaPengirim = items.NamaPengirim
		pengirim.TelpPengirim = items.TelpPengirim
		pengirim.ProvinceName = items.ProvinceName
		pengirim.CityName = items.CityName
		pengirim.AlamatPengirim = items.AlamatPengirim
		pengirim.KodePosPengirim = items.KodePosPengirim
		pengirims = append(pengirims, pengirim)
	}

	return pengirims, nil
}
