package service

import (
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/repository"
)

type PenerimaService struct {
	penerimaRepo repository.IPenerimaRepository
}

type IPenerimaService interface {
	GetPenerima() ([]entity.PenerimaModelView, error)
}

func NewPenerimaService(penerimaRepo repository.IPenerimaRepository) *PenerimaService {
	var penerimaService = PenerimaService{}
	penerimaService.penerimaRepo = penerimaRepo
	return &penerimaService
}

func (s *PenerimaService) GetPenerima() ([]entity.PenerimaModelView, error) {
	result, err := s.penerimaRepo.GetPenerima()
	if err != nil {
		return nil, err
	}

	var penerimas []entity.PenerimaModelView
	for _, items := range result {
		var penerima entity.PenerimaModelView
		penerima.IdPenerima = items.IdPenerima
		penerima.NamaPenerima = items.NamaPenerima
		penerima.TelpPenerima = items.TelpPenerima
		penerima.ProvinceNamePenerima = items.ProvinceNamePenerima
		penerima.CityNamePenerima = items.CityNamePenerima
		penerima.AlamatPenerima = items.AlamatPenerima
		penerima.KodePosPenerima = items.KodePosPenerima
		penerimas = append(penerimas, penerima)
	}

	return penerimas, nil
}
