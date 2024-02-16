package data

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Nama     string
	HP       string
	Password string
	Alamat   string
	Saldo    int
	// Transfer_Pengirim []Transfer //`gorm:"foreignKey:HP_Pengirim"`
	// Transfer_Penerima []Transfer //`gorm:"foreignKey:HP_Penerima"`
	// Topup             []Transfer `gorm:"foreignKey:HP"`
}

// type Transfer struct {
// 	gorm.Model
// 	HP_Pengirim string
// 	HP_Penerima string
// 	Nominal     int
// }

type Topup struct {
	gorm.Model
	HP      string
	Nominal int
}
