package entity

type Pengirim struct {
	IdPengirim      int    `gorm:"primary_key;column:id_pengirim"`
	NamaPengirim    string `json:"namaPengirim"`
	TelpPengirim    string `json:"telpPengirim"`
	ProvinceName    string `gorm:"column:provinsi_pengirim"`
	CityName        string `gorm:"column:kota_pengirim"`
	AlamatPengirim  string `json:"alamatPengirim"`
	KodePosPengirim string `json:"kodePosPengirim"`
}

type PengirimModelView struct {
	IdPengirim      int    `json:"idPengirim"`
	NamaPengirim    string `json:"namaPengirim"`
	TelpPengirim    string `json:"telpPengirim"`
	ProvinceName    string `json:"provinceName"`
	CityName        string `json:"cityName"`
	AlamatPengirim  string `json:"alamatPengirim"`
	KodePosPengirim string `json:"kodePosPengirim"`
}
