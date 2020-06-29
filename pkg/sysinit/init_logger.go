package sysinit

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var LOG *zap.SugaredLogger

func InitLogger() {
	// 日志文件
	hook := lumberjack.Logger{
		Filename:   GCF.UString("logger.path", "./logs/log"), // 日志文件路径
		MaxSize:    5,                                        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                                       // 日志文件最多保存多少个备份
		MaxAge:     7,                                        // 文件最多保存多少天
		Compress:   true,                                     // 是否压缩
	}

	// 日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "line",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // level 大写编码器
		//EncodeTime:     zapcore.RFC3339TimeEncoder,     //  时间格式
		EncodeTime:     TimeEncoder, //  自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 路径编码器  FullCallerEncoder:可以点击进入日志行 ShortCallerEncoder：不可进入
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	var logLevel zapcore.Level
	switch strings.ToUpper(GCF.UString("logger.level", "INFO")) {
	case "DEBUG":
		logLevel = zap.DebugLevel
	case "INFO":
		logLevel = zap.InfoLevel
	case "WARNING", "WARN":
		logLevel = zap.WarnLevel
	case "ERROR":
		logLevel = zap.ErrorLevel
	case "CRITICAL", "FATAL":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)

	// 编码器配置 json/console
	var styleEncoder zapcore.Encoder
	if GCF.UString("logger.style", "console") == "console" {
		styleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		styleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		styleEncoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	//logger := zap.New(core, caller, development, )
	logger := zap.New(core, caller, development)

	//赋值给全局sugar
	LOG = logger.Sugar()
}

// 自定义日期样式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%v]", t.Format("2006-01-02 15:04:05")))
}
