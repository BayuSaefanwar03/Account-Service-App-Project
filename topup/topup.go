package topup

import (
	"gorm.io/gorm"
)

type Topup struct {
	gorm.Model
	HP      string
	Nominal int
}

type Users struct {
	HP    string
	Saldo int
}

func Newtopup(connection *gorm.DB, user Users, nominal int) (bool, error) {
	err := connection.Create(&Topup{HP: user.HP, Nominal: nominal}).Error
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
