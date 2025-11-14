package log

import (
	"github.com/SemgaTeam/blog/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"
)

var Log *zap.Logger

func InitLogger(logFile string) {
	conf := config.GetConfig()
	debug := conf.App.Debug

	rotate := &lumberjack.Logger{
		Filename: logFile,
		MaxSize: 10,
		MaxBackups: 5,
		MaxAge: 30,
		Compress: true,
	}

	fileWriter := zapcore.AddSync(rotate) // rotates files

	consoleWriter := zapcore.AddSync(os.Stdout) // writes to console

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "time"
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"

	var consoleEncoder zapcore.Encoder
	if debug {
		consoleEncoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		consoleEncoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	var logLevel zapcore.Level
	if debug {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.InfoLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleWriter, logLevel),
		zapcore.NewCore(fileEncoder, fileWriter, logLevel),
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
