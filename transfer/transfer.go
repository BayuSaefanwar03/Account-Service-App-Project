package transfer

import (
	"fmt"
	"project1/data"

	"gorm.io/gorm"
)

func CheckUser(connection *gorm.DB, HP string) bool {
	var result data.Users
	connection.Where("hp = ?", HP).Limit(1).Find(&result)
	return result.HP != ""
}

func Send(connection *gorm.DB, pengirim data.Users, penerima string, nominal int) error {

	if pengirim.HP == penerima {
		return fmt.Errorf("anda memasukkan nomer anda sendiri")
	}

	if nominal <= 0 {
		return fmt.Errorf("nominal kurang dari 1")
	}

	if pengirim.Saldo < nominal {
		return fmt.Errorf("nominal yang di input kurang dari saldo")
	}

	if !CheckUser(connection, penerima) {
		return fmt.Errorf("nomor tujuan tidak di temukan")
	}

	err := connection.Create(&data.Transfer{HP_Pengirim: pengirim.HP, HP_Penerima: penerima, Nominal: nominal}).Error
	if err != nil {
		return err
	}

	var saldo_penerima data.Users
	connection.Where("hp = ?", penerima).Limit(1).Find(&saldo_penerima)
	query1 := connection.Table("users").Where("hp = ?", penerima).Update("saldo", saldo_penerima.Saldo+nominal)
	if err := query1.Error; err != nil || query1.RowsAffected == 0 {
		return fmt.Errorf("gagal mengubah saldo pengirim")
	}

	query2 := connection.Table("users").Where("hp = ?", penerima).Update("saldo", pengirim.Saldo+nominal)
	if err := query2.Error; err != nil || query2.RowsAffected == 0 {
		return fmt.Errorf("gagal menambah ke saldo tujuan")
	}

	return nil
}

func HistoryTransfer(connection *gorm.DB, user data.Users) []data.Transfer {
	var result []data.Transfer
	connection.Where("hp_pengirim = ?", user.HP).Find(&result)
	return result
}
