package topup

import (
	"project1/data"

	"gorm.io/gorm"
)

func Newtopup(connection *gorm.DB, user data.Users, nominal int) (bool, error) {
	err := connection.Create(&data.Topup{HP: user.HP, Nominal: nominal}).Error
	if err != nil {
		return false, err
	} else {
		query := connection.Table("users").Where("hp = ?", user.HP).Update("saldo", user.Saldo+nominal)
		if err := query.Error; err != nil {
			return false, err
		}

		return query.RowsAffected > 0, nil
	}

}

func HistoryTopup(connection *gorm.DB, user data.Users) []data.Topup {
	var result []data.Topup

	connection.Where("hp = ?", user.HP).Find(&result)
	return result

}
