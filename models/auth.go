package models

import "github.com/jinzhu/gorm"

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Roles        string `json:"roles"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Email        string `json:"email"`
	Model
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
