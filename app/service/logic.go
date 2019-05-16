package service

import (
	"errors"
	"fmt"
	models "gin-api/app/model"
	"gin-api/pkg/logs"
	"github.com/google/uuid"
)

// Ping ping the resource.
func (s *Service) Ping() error {
	db, err := models.GetConnect()
	defer func() {
		_ = db.Close()
	}()
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

	user, err := models.NewUser().GetUserByMobile(userName)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("user GetUserByMobile err: %s", err.Error()))
	}
	if user != nil {

	}
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
