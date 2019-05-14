package logs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var (
	Logger *zap.Logger
)

func Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log_file"),
		MaxSize:    200, //最大200M
		MaxBackups: 10,  //最多保存7个日志文件
		MaxAge:     10,
	})

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.EpochTimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	)
	Logger = zap.New(core)
}


func LogErr(action string,  msg string) {
	Logger.Error(msg,
		zap.String("trace_id", traceId),
		zap.String("log_time", time.Now().Format("2006-01-02 15:04:05")))
}

func LogInfo(action string, traceId string, msg string) {
	Logger.Info(msg,
		zap.String("trace_id", traceId),
		zap.String("log_time", time.Now().Format("2006-01-02 15:04:05")))
}

func LogWarn(action string, traceId string, msg string) {
	Logger.Warn(msg,
		zap.String("trace_id", traceId),
		zap.String("log_time", time.Now().Format("2006-01-02 15:04:05")))
}