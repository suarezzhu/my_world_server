package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//type User struct {
//
//	ID           int    `gorm:"primary_key" json:"id"`
//	Username     string `json:"username"`
//	Password     string `json:"password"`
//	roles        string `json:"roles"`
//	avatar       string `json:"avatar"`
//	introduction string `json:"introduction"`
//	email        string `json:"email"`
//}

func GetUserInfoId(id int) (*User, error) {
	var user User
	//err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	fmt.Println(nil, err)
	//	return nil, err
	//}
	//return &user, nil

	err := db.Raw("select * from m_user where id=?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(&user, "user123")
	return &user, nil

}
