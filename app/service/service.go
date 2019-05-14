package service

import (
	"context"
	"fmt"
	models "gin-api/app/model"
	"gin-api/pkg/logs"
	"github.com/spf13/viper"
)

type Service struct {
}

func New()(s *Service) {
	//配置管理
	viper.SetConfigName("dev")
	viper.AddConfigPath("../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error when load config file: %s \n", err))
	}
	logs.Init()
	models.Init()
	s = &Service{}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {

	return nil
}

// Close close the resource.
func (s *Service) Close() {
	//s.dao.Close()
}