package service

import (
	models "gin-api/app/model"
	"gin-api/pkg/logs"
	"gin-api/pkg/redis"
	"github.com/spf13/viper"
	"log"
)

type Service struct {
}

func New() (s *Service) {
	viper.SetConfigName(getEnv())
	viper.AddConfigPath("../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error when load config file: %s \n", err)
	}
	//模块启动
	logs.Init()
	models.Init()
	redis.Init()
	s = &Service{}
	return s
}

const ENV = "GIN_API"

func getEnv() string {
	_ = viper.BindEnv(ENV)
	env := viper.Get(ENV)
	if nil == env {
		return "local"
	}
	return env.(string)
}

// Close close the resource.
func (s *Service) Close() {
	models.Close()
}
