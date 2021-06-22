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
	GetResi(string) (*entity.TransaksiModelView, error)
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
		transaksi.PenerimaPaket = item.PenerimaPaket
		transaksi.FotoPenerima = item.FotoPenerima
		transaksi.TanggalSampai = item.TanggalTransaksi
		transaksi.IsDelete = item.IsDelete
		//kurir
		transaksi.IdKurir = item.IdKurir
		transaksi.NamaKurir = item.Kurir.NamaKurir
		transaksi.NoTelpKurir = item.Kurir.NoTelpKurir
		transaksi.Alamat = item.Kurir.Alamat
		transaksi.Nik = item.Kurir.Nik
		transaksi.Ttl = item.Kurir.Ttl
		transaksi.File = item.Kurir.File
		//penerima
		transaksi.IdPenerima = item.IdPenerima
		transaksi.NamaPenerima = item.Penerima.NamaPenerima
		transaksi.TelpPenerima = item.Penerima.TelpPenerima
		transaksi.ProvinceNamePenerima = item.Penerima.ProvinceNamePenerima
		transaksi.CityNamePenerima = item.Penerima.CityNamePenerima
		transaksi.AlamatPenerima = item.Penerima.AlamatPenerima
		transaksi.KodePosPenerima = item.Penerima.KodePosPenerima
		//pengirirm
		transaksi.IdPengirim = item.IdPengirim
		transaksi.NamaPengirim = item.Pengirim.NamaPengirim
		transaksi.TelpPengirim = item.Pengirim.TelpPengirim
		transaksi.ProvinceName = item.Pengirim.ProvinceName
		transaksi.CityName = item.Pengirim.CityName
		transaksi.AlamatPengirim = item.Pengirim.AlamatPengirim
		transaksi.KodePosPengirim = item.Pengirim.KodePosPengirim
		transaksis = append(transaksis, transaksi)
	}

	return &transaksis, nil

}

func (s *TransaksiService) GetResi(resi string) (*entity.TransaksiModelView, error) {
	result, err := s.transRepo.GetResi(resi)

	if err != nil {
		return nil, err
	}

	var transVM entity.TransaksiModelView

	if result != nil {
		transVM = entity.TransaksiModelView{
			IdTransaksi:      result.IdTransaksi,
			NamaBarang:       result.NamaBarang,
			JumlahBarang:     result.JumlahBarang,
			KategoriLayanan:  result.KategoriLayanan,
			OngkosKirim:      result.OngkosKirim,
			Estimasi:         result.Estimasi,
			BeratBarang:      result.BeratBarang,
			StatusDelivery:   result.StatusDelivery,
			TanggalTransaksi: result.TanggalTransaksi,
			Resi:             result.Resi,
			PenerimaPaket:    result.PenerimaPaket,
			FotoPenerima:     result.FotoPenerima,
			TanggalSampai:    result.TanggalTransaksi,
			IsDelete:         result.IsDelete,
			//kurir
			IdKurir:     result.IdKurir,
			NamaKurir:   result.Kurir.NamaKurir,
			NoTelpKurir: result.Kurir.NoTelpKurir,
			Alamat:      result.Kurir.Alamat,
			Nik:         result.Kurir.Nik,
			Ttl:         result.Kurir.Ttl,
			File:        result.Kurir.File,
			//penerima
			IdPenerima:           result.IdPenerima,
			NamaPenerima:         result.Penerima.NamaPenerima,
			TelpPenerima:         result.Penerima.TelpPenerima,
			ProvinceNamePenerima: result.Penerima.ProvinceNamePenerima,
			CityNamePenerima:     result.Penerima.CityNamePenerima,
			AlamatPenerima:       result.Penerima.AlamatPenerima,
			KodePosPenerima:      result.Penerima.KodePosPenerima,
			//pengirirm
			IdPengirim:      result.IdPengirim,
			NamaPengirim:    result.Pengirim.NamaPengirim,
			TelpPengirim:    result.Pengirim.TelpPengirim,
			ProvinceName:    result.Pengirim.ProvinceName,
			CityName:        result.Pengirim.CityName,
			AlamatPengirim:  result.Pengirim.AlamatPengirim,
			KodePosPengirim: result.Pengirim.KodePosPengirim,
		}
	}

	return &transVM, nil
}
