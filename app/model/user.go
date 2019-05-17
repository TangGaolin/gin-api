package models

import (
	"fmt"
	"gin-api/pkg/logs"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `gorm:"column:passwd"  json:"passwd"`
	Salt     string `json:"salt"`
	Status   string `json:"status"`
	Ctime    string `json:"ctime"`
	Utime    string `json:"utime"`
}

func NewUser() *User {
	return &User{}
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "users"
}

func (u *User) GetUserByMobile(userName string) (*User, error) {

	var user User
	err := db.Where("name = ?", userName).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return &user, nil
		}
		logs.Logger.Error(fmt.Sprintf("user GetUserByMobile err: %s", err.Error()))
	}
	return &user, err

}
