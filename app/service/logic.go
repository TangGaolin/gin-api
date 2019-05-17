package service

import (
	"errors"
	models "gin-api/app/model"
	"gin-api/pkg/utils"
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

func (s *Service) Login(userName string, password string) (map[string]interface{}, error) {

	user, err := models.NewUser().GetUserByMobile(userName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户名密码不正确！")
	}
	if user.Password == utils.Md5(password+user.Salt) {
		return map[string]interface{}{
			"user_info": map[string]string{
				"user_name": user.Name,
				"mobile":    user.Mobile,
				"email":     user.Email,
				"status":    user.Status,
			},
			"token": uuid.New().String(),
		}, nil
	}
	return nil, errors.New("用户名密码不正确！")

}

func (s *Service) AuthInfo(userName string) (map[string]interface{}, error) {

	user, err := models.NewUser().GetUserByMobile(userName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("暂无用户信息！")
	}
	return map[string]interface{}{
		"user_info": map[string]string{
			"user_name": user.Name,
			"mobile":    user.Mobile,
			"email":     user.Email,
			"status":    user.Status,
		},
	}, nil

}
