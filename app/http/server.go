package http

import (
	"gin-api/app/service"
	"gin-api/pkg/logs"
	"gin-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var (
	svc *service.Service
)

//-------------为 gin.Context 新增方法---------------------------------------
type Context struct {
	*gin.Context
}
type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

func (c *Context) success(data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	c.JSON(http.StatusOK, gin.H{
		"rid":  c.MustGet("rid").(string),
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}
func (c *Context) error(code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"rid":  c.MustGet("rid").(string),
		"code": code,
		"msg":  msg,
	})
}

//----------------------------------------------------

func New(s *service.Service) (httpSrv *http.Server) {
	gin.SetMode(viper.GetString("mode"))
	router := gin.New()
	initRouter(router)
	httpSrv = &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}
	svc = s
	go func() {
		//启动http服务
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("http server start listen: %s\n", viper.GetString("port"))
	return
}

func initRouter(router *gin.Engine) {
	router.Use(Logger)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome gin api Server")
	})
	router.GET("/ping", Handle(ping)) //服务检测

	api := router.Group("/api")
	api.GET("/login", Handle(login))
	api.Use(Auth)
	{
		api.GET("/auth_info", Handle(authInfo))
	}
}

//中间件 Logger
func Logger(c *gin.Context) {
	t := time.Now() //记录开始时间
	rid := utils.GenRid()
	c.Set("rid", rid)
	c.Next()
	logs.Logger.Info("",
		zap.String("router", c.Request.URL.Path),
		zap.String("trace_id", rid),
		zap.Int64("cost", int64(time.Since(t).Nanoseconds()/1000000)),
		zap.String("log_time", t.Format("2006-01-02 15:04:05")),
	)
}

//中间件 Auth
func Auth(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"rid":  c.MustGet("rid").(string),
			"code": 403,
			"msg":  "无访问权限",
		})
		return
	}
	c.Next()
}
