## 概述
基于Go mod和gin web 框架搭建的上手即用api服务

### 使用的库

* [ ] gin:  web开发中比较常用的框架
* [ ] viper: 方便好用的配置库
* [ ] zap: 高性能日志库zap 
* [ ] gorm: 开发人员友好的全功能ORM
* [ ] redigo: 用池管理redis连接


### 利用这些库实现了什么
* [ ] 信号管理
* [ ] 路由、中间件、数据响应 - gin
* [ ] 配置文件 - viper & yml
* [ ] 日志管理 - zap
* [ ] 数据库管理 - gorm
* [ ] 请求参数处理、业务代码、数据持久化等逻辑分层

### 代码结构

```
├── README.md               
├── app                     应用主体
│   ├── http                
│   │   ├── controller.go   //请求参数处理
│   │   └── server.go       //http server主体
│   ├── model               //数据库操作
│   │   └── db.go           
│   └── service
│       ├── logic.go        //业务逻辑代码
│       └── service.go      //app服务主体
├── cmd
│   └── main.go             //应用入口
├── configs                 //各个环境配置文件
│   ├── local.yml
│   ├── qa.yml
│   └── prod.yml
├── go.mod
├── go.sum
└── pkg                     //公共代码库
    ├── logs
    │   └── log.go
    ├── redis
    │   └── redis.go
    └── utils
        └── func.go
```

