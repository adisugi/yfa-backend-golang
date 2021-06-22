package entity

type Transaksi struct {
	IdTransaksi      int      `gorm:"primarykey:IdTransaksi"`
	NamaBarang       string   `json:"namaBarang"`
	JumlahBarang     int      `json:"jumlahBarang"`
	KategoriLayanan  string   `json:"kategoriLayanan"`
	OngkosKirim      int      `json:"ongkosKirim"`
	Estimasi         string   `json:"estimasi"`
	BeratBarang      int      `json:"beratBarang"`
	StatusDelivery   string   `json:"statusDelivery"`
	TanggalTransaksi string   `json:"tanggalTransaksi"`
	Resi             string   `json:"resi"`
	PenerimaPaket    string   `json:"penerimaPaket"`
	FotoPenerima     string   `json:"fotoPenerima"`
	TanggalSampai    string   `json:"tanggalSampai"`
	IsDelete         int      `json:"isDelete"`
	IdKurir          int      `json:"idKurir"`
	Kurir            Kurir    `gorm:"foreignkey:IdKurir"`
	IdPenerima       int      `json:"idPenerima"`
	Penerima         Penerima `gorm:"foreignkey:IdPenerima"`
	IdPengirim       int      `json:"idPengirim"`
	Pengirim         Pengirim `gorm:"foreignkey:IdPengirim"`
}

type TransaksiModelView struct {
	IdTransaksi          int    `json:"idTransaksi"`
	NamaBarang           string `json:"namaBarang"`
	JumlahBarang         int    `json:"jumlahBarang"`
	KategoriLayanan      string `json:"kategoriLayanan"`
	OngkosKirim          int    `json:"ongkosKirim"`
	Estimasi             string `json:"estimasi"`
	BeratBarang          int    `json:"beratBarang"`
	StatusDelivery       string `json:"statusDelivery"`
	TanggalTransaksi     string `json:"tanggalTransaksi"`
	Resi                 string `json:"resi"`
	PenerimaPaket        string `json:"penerimaPaket"`
	FotoPenerima         string `json:"fotoPenerima"`
	TanggalSampai        string `json:"tanggalSampai"`
	IsDelete             int    `json:"isDelete"`
	IdKurir              int    `json:"idKurir"`
	NamaKurir            string `json:"namaKurir"`
	NoTelpKurir          string `json:"noTelpKurir"`
	Alamat               string `json:"alamat"`
	Nik                  string `json:"nik"`
	Ttl                  string `json:"ttl"`
	File                 string `json:"file"`
	IdPenerima           int    `json:"idPenerima"`
	NamaPenerima         string `json:"namaPenerima"`
	TelpPenerima         string `json:"telpPenerima"`
	ProvinceNamePenerima string `json:"provinceNamePenerima"`
	CityNamePenerima     string `json:"cityNamePenerima"`
	AlamatPenerima       string `json:"alamatPenerima"`
	KodePosPenerima      string `json:"kodePosPenerima"`
	IdPengirim           int    `json:"idPengirim"`
	NamaPengirim         string `json:"namaPengirim"`
	TelpPengirim         string `json:"telpPengirim"`
	ProvinceName         string `json:"provinceName"`
	CityName             string `json:"cityName"`
	AlamatPengirim       string `json:"alamatPengirim"`
	KodePosPengirim      string `json:"kodePosPengirim"`
	//Kurir             Kurir    `json:"kurir"`
}
