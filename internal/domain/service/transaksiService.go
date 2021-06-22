package service

import (
	"yfa-golang/internal/domain/entity"
	"yfa-golang/internal/domain/repository"
)

type TransaksiService struct {
	transRepo repository.ITransaksiRepository
}

type ITransaksiService interface {
	GetTransaksi() (*[]entity.TransaksiModelView, error)
}

func NewTransaksiService(transRepo repository.ITransaksiRepository) *TransaksiService {
	var transService = TransaksiService{}
	transService.transRepo = transRepo
	return &transService
}

func (s *TransaksiService) GetTransaksi() (*[]entity.TransaksiModelView, error) {
	result, err := s.transRepo.GetTransaksi()
	if err != nil {
		return nil, err
	}

	var transaksis []entity.TransaksiModelView

	for _, item := range *result {
		var transaksi entity.TransaksiModelView
		//resultKurir, _ := repository.KurirRepository.GetOneKurir(item.IdKurir)
		transaksi.IdTransaksi = item.IdTransaksi
		transaksi.NamaBarang = item.NamaBarang
		transaksi.JumlahBarang = item.JumlahBarang
		transaksi.KategoriLayanan = item.KategoriLayanan
		transaksi.OngkosKirim = item.OngkosKirim
		transaksi.Estimasi = item.Estimasi
		transaksi.BeratBarang = item.BeratBarang
		transaksi.StatusDelivery = item.StatusDelivery
		transaksi.TanggalTransaksi = item.TanggalTransaksi
		transaksi.Resi = item.Resi
		transaksi.PenerimaPaket = item.Resi
		transaksi.FotoPenerima = item.FotoPenerima
		transaksi.TanggalSampai = item.TanggalTransaksi
		transaksi.IsDelete = item.IsDelete
		transaksi.IdKurir = item.IdKurir
		transaksi.NamaKurir = item.Kurir.NamaKurir
		transaksi.NoTelpKurir = item.Kurir.NoTelpKurir
		transaksi.Alamat = item.Kurir.Alamat
		transaksi.Nik = item.Kurir.Nik
		transaksi.Ttl = item.Kurir.Ttl
		transaksi.File = item.Kurir.File
		transaksis = append(transaksis, transaksi)
	}

	return &transaksis, nil

}
