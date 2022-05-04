package logger

import (
	"errors"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	ErrWrongLogFile = errors.New("wrong log file")
)

func NewLogFile(logFilePath string) (*os.File, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func NewJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(time.RFC1123))
	}
	encoderConfig.TimeKey = "time"
	return zapcore.NewJSONEncoder(encoderConfig)
}

func NewFileLogger(logFile *os.File) (logr.Logger, error) {
	log := logr.Logger{}
	if logFile == nil {
		return log, ErrWrongLogFile
	}
	uLogger, err := zap.NewProduction()
	if err != nil {
		return logr.Logger{}, err
	}
	log = zapr.NewLogger(uLogger)
	return log, nil

}

func NewLogger(scope string, toConsole bool) (logr.Logger, error) {
	log := logr.Logger{}
	if scope == "" {
		return log, errors.New("empty scope name")
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(time.RFC1123))
	}
	encoderConfig.TimeKey = "time"

	var config zap.Config
	if toConsole {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig = encoderConfig
	zapLogger, err := config.Build()
	if err != nil {
		return log, err
	}
	return zapr.NewLogger(zapLogger), nil

}
