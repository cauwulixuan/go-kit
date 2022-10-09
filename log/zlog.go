/*
Copyright 2022 The Inspur AIStation Group Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.

package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	logger  *zap.Logger
	Slogger *zap.SugaredLogger
)

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func init() {
	multi := viper.GetBool("log.multi_staging")
	if multi {
		InitWithMultiLevelOutPut()
	} else {
		InitWithSingleLevelOutput()
	}

}

func InitWithSingleLevelOutput() {
	level := getLogLevel(viper.GetString("log.level"))
	atom := zap.NewAtomicLevelAt(level)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewCustomEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getAllLogWriter())),
		atom,
	)

	// 1. AddCaller with file name and line number.
	// 2. AddStacktrace record a stack trace for all messages at or above WARN level.
	// 3. Add serviceName field.
	field := zap.Fields(zap.String("serviceName", viper.GetString("svc_name")))
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.LevelEnablerFunc(warnLevel)), field)
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	Slogger = logger.Sugar()
	Slogger.Info("Setting logger successfully.")
}

// return log with level above INFO
func warnLevel(l zapcore.Level) bool {
	return l > zapcore.InfoLevel && l > getLogLevel(viper.GetString("log.level"))
}

func infoLevel(l zapcore.Level) bool {
	return l <= zapcore.InfoLevel && l > getLogLevel(viper.GetString("log.level"))
}

func InitWithMultiLevelOutPut() {
	atom := zap.NewAtomicLevelAt(getLogLevel(viper.GetString("log.level")))
	// define LevelEnablerFunc
	infoLvl := zap.LevelEnablerFunc(infoLevel)
	warnLvl := zap.LevelEnablerFunc(warnLevel)

	// get writer
	infoWriter := getInfoLogWriter()
	warnWriter := getWarnLogWriter()

	// with multiple output
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(NewCustomEncoderConfig()), zapcore.AddSync(infoWriter), infoLvl),
		zapcore.NewCore(zapcore.NewJSONEncoder(NewCustomEncoderConfig()), zapcore.AddSync(warnWriter), warnLvl),
		zapcore.NewCore(zapcore.NewConsoleEncoder(NewCustomEncoderConfig()), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), atom),
	)
	// 1. AddCaller with file name and line number.
	// 2. AddStacktrace record a stack trace for all messages at or above WARN level.
	// 3. Add serviceName field.
	field := zap.Fields(zap.String("serviceName", viper.GetString("svc_name")))
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(warnLvl), field)
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	Slogger = logger.Sugar()
	Slogger.Info("Setting logger successfully.")

}

func NewCustomEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func getEncodeTime() zapcore.TimeEncoder {
	return zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.0000000")
}

func getWarnLogWriter() io.Writer {
	return getLogWriter(viper.GetString("log.rotate.warn_log_path"))
}

func getInfoLogWriter() io.Writer {
	return getLogWriter(viper.GetString("log.rotate.info_log_path"))
}

func getAllLogWriter() io.Writer {
	return getLogWriter(viper.GetString("log.rotate.all_log_path"))
}

func getLogWriter(path string) io.Writer {
	return &lumberjack.Logger{
		Filename: path,
		// unit: megabytes
		MaxSize: viper.GetInt("log.rotate.max_size"),
		// max number of backup files
		MaxBackups: viper.GetInt("log.rotate.max_backups"),
		// max age of keepping log files
		MaxAge: viper.GetInt("log.rotate.max_age"),
		// Compress backup log files or not, default false
		Compress: viper.GetBool("log.rotate.compress"),
	}
}
