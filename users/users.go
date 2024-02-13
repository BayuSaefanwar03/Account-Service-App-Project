package users

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nama     string
	HP       string
	Password string
	Alamat   string
	Saldo    int
}

func (u *Users) GantiPassword(connection *gorm.DB, newPassword string) (bool, error) {
	query := connection.Table("users").Where("hp = ?", u.HP).Update("password", newPassword)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func Register(connection *gorm.DB, newUser Users) (bool, error) {
	var user Users
	query := connection.Where("hp = ?", newUser.HP).Limit(1).Find(&user)
	if err := query.Error; err != nil {
		return false, err
	} else if user.HP != "" {
		return false, nil
	} else {
		err2 := connection.Create(&newUser).Error
		if err2 != nil {
			return false, err2
		}
		return true, nil
	}
}

func Login(connection *gorm.DB, hp string, password string) (Users, bool, error) {
	var user Users
	err := connection.Where("hp = ? AND password = ?", hp, password).Limit(1).Find(&user).Error
	if err != nil {
		return Users{}, false, err
	} else if user.HP == "" {
		return Users{}, false, nil
	} else {
		return user, true, nil
	}
}
