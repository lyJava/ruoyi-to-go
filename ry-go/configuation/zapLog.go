package configuation

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLog *zap.Logger

// init zap日志初始化
func init() {

	dateTimeFormat := "2006-01-02 15:04:05.000"

	// 构建日志目录路径
	logDir := filepath.Join("logs")
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 构建日志文件路径
	logFileName := filepath.Join(logDir, "ry-go-echo.log")
	abs, _ := filepath.Abs(logDir)
	zap.L().Sugar().Infof("当前zap日志===%s", abs)

	// 文件日志的编码器（无颜色）
	fileConfig := zap.NewProductionEncoderConfig()
	fileConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileConfig.EncodeTime = parseTime(dateTimeFormat)
	fileConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 控制台日志的编码器（有颜色）
	consoleConfig := zap.NewDevelopmentEncoderConfig()
	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleConfig.EncodeTime = parseTime(dateTimeFormat)
	consoleConfig.EncodeCaller = zapcore.ShortCallerEncoder

	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    5,
		MaxAge:     30,
		MaxBackups: 30,
		Compress:   false,
	}
	fileSync := zapcore.AddSync(lumberjackLogger)
	consoleSync := zapcore.Lock(os.Stdout)

	fileEncoder := zapcore.NewConsoleEncoder(fileConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileSync, zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, consoleSync, zapcore.DebugLevel),
	)
	ZapLog = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(ZapLog)
}

// parseTime 进行时间格式处理
func parseTime(layout string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(layout))
	}
}

// customerLogFileName 自定义日志文件名
func customerLogFileName(baseName, dateTimeFormat string) string {
	return baseName + "_" + time.Now().Format(dateTimeFormat) + ".log"
}
