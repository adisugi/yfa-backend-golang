package entity

type Kurir struct {
	IdKurir     int    `gorm:"primary_key;column:id_kurir"`
	NamaKurir   string `json:"namaKurir"`
	NoTelpKurir string `json:"noTelpKurir"`
	Alamat      string `json:"alamat"`
	Nik         string `json:"nik"`
	Ttl         string `json:"ttl"`
	File        string `json:"file"`
	IsDelete    int    `json:"isDelete"`
}

type KurirModelView struct {
	IdKurir     int    `json:"idKurir"`
	NamaKurir   string `json:"namaKurir"`
	NoTelpKurir string `json:"noTelpKurir"`
	Alamat      string `json:"alamat"`
	Nik         string `json:"nik"`
	Ttl         string `json:"ttl"`
	File        string `json:"file"`
	IsDelete    int    `json:"isDelete"`
}
