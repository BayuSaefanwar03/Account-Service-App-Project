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

func HistoryTopup(connection *gorm.DB, user data.Users) (bool, error) {
	err := connection.Where("hp = ?", user.HP).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, err
		}
		return false, err
	}
	var count int64
	err = connection.Model(&data.Topup{}).Where("hp = ?", user.HP).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
