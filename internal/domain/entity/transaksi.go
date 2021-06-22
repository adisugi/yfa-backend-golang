package entity

type Transaksi struct {
	IdTransaksi		  int      `gorm:"primarykey:IdTransaksi"`
	NamaBarang        string   `json:"namaBarang"`
	JumlahBarang      int      `json:"jumlahBarang"`
	KategoriLayanan   string   `json:"kategoriLayanan"`
	OngkosKirim       int      `json:"ongkosKirim"`
	Estimasi		  string   `json:"estimasi"`
	BeratBarang		  int      `json:"beratBarang"`
	StatusDelivery    string   `json:"statusDelivery"`
	TanggalTransaksi  string   `json:"tanggalTransaksi"`
	Resi              string   `json:"resi"`
	PenerimaPaket     string   `json:"penerimaPaket"`
	FotoPenerima      string   `json:"fotoPenerima"`
	TanggalSampai     string   `json:"tanggalSampai"`
	IsDelete          int      `json:"isDelete"`
	IdKurir           int      `json:"idKurir"`
	Kurir             Kurir    `gorm:"foreignkey:IdKurir"`
}

type TransaksiModelView struct {
	IdTransaksi		  int      `json:"idTransaksi"`
	NamaBarang        string   `json:"namaBarang"`
	JumlahBarang      int      `json:"jumlahBarang"`
	KategoriLayanan   string   `json:"kategoriLayanan"`
	OngkosKirim       int      `json:"ongkosKirim"`
	Estimasi		  string   `json:"estimasi"`
	BeratBarang		  int      `json:"beratBarang"`
	StatusDelivery    string   `json:"statusDelivery"`
	TanggalTransaksi  string   `json:"tanggalTransaksi"`
	Resi              string   `json:"resi"`
	PenerimaPaket     string   `json:"penerimaPaket"`
	FotoPenerima      string   `json:"fotoPenerima"`
	TanggalSampai     string   `json:"tanggalSampai"`
	IsDelete          int      `json:"isDelete"`
	IdKurir           int      `json:"idKurir"`
	NamaKurir         string   `json:"namaKurir"`
	NoTelpKurir       string   `json:"noTelpKurir"`
	Alamat            string   `json:"alamat"`
	Nik               string   `json:"nik"`
	Ttl               string   `json:"ttl"`
	File              string   `json:"file"`
	//Kurir             Kurir    `json:"kurir"`
}

