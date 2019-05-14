package http

import (
	"fmt"
	"gin-api/app/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	svc *service.Service
)

func New(s *service.Service) (httpSrv *http.Server){

	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()
	initRouter(router)
	fmt.Println("new http:", viper.GetString("port"))
	httpSrv = &http.Server{
		Addr:    ":"+viper.GetString("port"),
		Handler: router,
	}
	svc = s
	go func() {
		//启动http服务
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return
}

func initRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome gin api Server")
	})
	router.GET("/ping", ping)
}


func ping(c *gin.Context) {
	if err := svc.Ping(c); err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
	c.JSON(http.StatusOK, "pong")
}