package users

import (
	"Account-Service/src/data"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, hp string, password string) (data.Users, error) {
	var result data.Users
	err := db.Find(&result, data.Users{HP: hp, Password: password}).Error
	if err != nil {
		return data.Users{}, err
	}

	return result, nil
}

func Register(connection *gorm.DB, newUser data.Users) (bool, error) {
	var user data.Users
	query := connection.Where("hp = ?", newUser.HP).Find(&user)
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

func Update(db *gorm.DB, user_old data.Users, user_new data.Users) (bool, error) {
	query := db.Model(&user_old).Updates(&user_new)
	if err := query.Error; err != nil {
		return false, err
	} else if !(query.RowsAffected > 0) {
		return false, nil
	} else {
		return true, nil
	}
}

func Delete(db *gorm.DB, user data.Users) (bool, error) {
	query := db.Delete(user)
	if err := query.Error; err != nil {
		return false, err
	} else if !(query.RowsAffected > 0) {
		return false, nil
	} else {
		return true, nil
	}
}

func View(db *gorm.DB, hp string) (data.Users, error) {
	var result data.Users
	err := db.Select("hp", "nama", "alamat", "saldo").Find(&result, data.Users{HP: hp}).Error
	if err != nil {
		return data.Users{}, err
	}

	return result, nil
}
