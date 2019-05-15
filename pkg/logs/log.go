package logs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

func logErrorF(msg string) string {
	Logger.Error(msg)
	return msg
}