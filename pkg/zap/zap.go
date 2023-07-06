package zap

import (
	"os"
	"sword-project/pkg/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	errorPath = "tmp/error.log"
	logPath   = "tmp/info.log"
)

var s *zap.SugaredLogger

func NewZapSugaredLogger() *zap.SugaredLogger {
	if s == nil {

		var zapConfig zap.Config
		switch configs.ApplicationCfg.Env {
		case configs.Development:
			cfg := zap.NewDevelopmentConfig()
			cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, errorPath)
			cfg.OutputPaths = append(cfg.OutputPaths, logPath)

			zapConfig = cfg
		default:

			config := zap.NewProductionEncoderConfig()
			config.EncodeTime = zapcore.ISO8601TimeEncoder
			fileEncoder := zapcore.NewJSONEncoder(config)
			logFile, _ := os.OpenFile(errorPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			writer := zapcore.AddSync(logFile)
			defaultLogLevel := zapcore.ErrorLevel
			core := zapcore.NewTee(
				zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			)
			return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		}

		logger, err := zapConfig.Build()
		if err != nil {
			panic(err)
		}
		s = logger.Sugar()
	}
	return s
}
