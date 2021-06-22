package entity

type Penerima struct {
	IdPenerima             int        `gorm:"primary_key;column:id_penerima"`
	NamaPenerima           string     `gorm:"column:nama_penerima"`
	TelpPenerima           string     `gorm:"column:no_telp_penerima"`
	ProvinceNamePenerima   string     `gorm:"column:provinsi_penerima"`
	CityNamePenerima       string     `gorm:"column:kota_penerima"`
	AlamatPenerima         string     `gorm:"column:alamatpenerima"`
	KodePosPenerima        string     `gorm:"column:kode_pos_penerima"`
}

type PenerimaModelView struct {
	IdPenerima             int        `json:"idPenerima"`
	NamaPenerima           string     `json:"namaPenerima"`
	TelpPenerima           string     `json:"telpPenerima"`
	ProvinceNamePenerima   string     `json:"provinceNamePenerima"`
	CityNamePenerima       string     `json:"cityNamePenerima"`
	AlamatPenerima         string     `json:"alamatPenerima"`
	KodePosPenerima        string     `json:"kodePosPenerima"`
}