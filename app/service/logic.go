package service

import (
	"errors"
	models "gin-api/app/model"
	"github.com/google/uuid"
)

// Ping ping the resource.
func (s *Service) Ping() error {
	db, err := models.GetConnect()
	defer db.Close()
	if err != nil {
		return err
	}
	return nil
}

//模拟用户数据
var userData = map[string]string{
	"admin":    "123456",
	"adminPro": "adminPro",
}

func (s *Service) Login(userName string, password string) (map[string]interface{}, error) {

	if _, ok := userData[userName]; ok {
		//存在
		if userData[userName] == password {
			return map[string]interface{}{
				"user_name": userName,
				"token":     uuid.New().String(),
			}, nil
		}
	}
	return nil, errors.New("用户名密码不正确！")
}

func (s *Service) AuthInfo(userName string) (map[string]interface{}, error) {

	if _, ok := userData[userName]; ok {
		//存在
		return map[string]interface{}{
			"user_name": userData[userName],
			"token":     uuid.New().String(),
		}, nil
	}
	return nil, errors.New("暂无用户信息！")
}
