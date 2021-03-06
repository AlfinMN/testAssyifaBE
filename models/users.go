package models

type Users struct {
	ID           string `json:"ID"`
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	NamaLengkap  string `json:"NamaLengkap"`
	TanggalLahir string `json:"TanggalLahir"`
	Alamat       string `json:"Alamat"`
	Photo        string `json:"Photo"`
}
