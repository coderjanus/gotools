package log

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config for logger
type Config struct {
	AppName     string
	LogLevel    string
	OutFile     string
	ErrFile     string
	Console     bool
	Development bool
}

// Logger Logger
type Logger struct {
	logger *zap.Logger
}

// InitLogger init a logger
func InitLogger(config Config) (*Logger, error) {
	outFilepath := []string{"stdout"}
	errFilepath := []string{"stderr"}
	if len(config.OutFile) > 0 {
		if config.Console {
			outFilepath = append(outFilepath, config.OutFile)
		} else {
			outFilepath[0] = config.OutFile
		}
	}
	if len(config.ErrFile) > 0 {
		if config.Console {
			errFilepath = append(errFilepath, config.ErrFile)
		} else {
			errFilepath[0] = config.ErrFile
		}
	}

	logger, err := initZapLog(config.AppName, config.Development, config.LogLevel, outFilepath, errFilepath)
	if err != nil {
		return nil, err
	}

	gLogger = &Logger{logger}
	return gLogger, nil
}

// Debug log a debug level message
func (l *Logger) Debug(msg string) *Logger {
	l.logger.Debug(msg)
	return l
}

// Debugf log a debug level message
func (l *Logger) Debugf(msg string, args ...interface{}) *Logger {
	l.logger.Debug(fmt.Sprintf(msg, args...))
	return l
}

// Debugw log a debug level message
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().Debugw(msg, keysAndValues...)
	return l
}

// Info log a info level message
func (l *Logger) Info(msg string) *Logger {
	l.logger.Info(msg)
	return l
}

// Infof log a info level message
func (l *Logger) Infof(msg string, args ...interface{}) *Logger {
	l.logger.Info(fmt.Sprintf(msg, args...))
	return l
}

// Infow log a info level message
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().Infow(msg, keysAndValues...)
	return l
}

// Warn log a warn level message
func (l *Logger) Warn(msg string) *Logger {
	l.logger.Sugar().Warn(msg)
	return l
}

// Warnf log a warn level message
func (l *Logger) Warnf(msg string, args ...interface{}) *Logger {
	l.logger.Warn(fmt.Sprintf(msg, args...))
	return l
}

// Warnw log a warn level message
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().Warnw(msg, keysAndValues...)
	return l
}

// Error log a error level message
func (l *Logger) Error(msg string) *Logger {
	l.logger.Sugar().Error(msg)
	return l
}

// Errorf log a error level message
func (l *Logger) Errorf(msg string, args ...interface{}) *Logger {
	l.logger.Error(fmt.Sprintf(msg, args...))
	return l
}

// Errorw log a error level message
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().Errorw(msg, keysAndValues...)
	return l
}

// Fatal log a fatal level message, then call os.Exit(1)
func (l *Logger) Fatal(msg string) *Logger {
	l.logger.Sugar().Fatal(msg)
	return l
}

// Fatalf log a fatal level message, then call os.Exit(1)
func (l *Logger) Fatalf(msg string, args ...interface{}) *Logger {
	l.logger.Fatal(fmt.Sprintf(msg, args...))
	return l
}

// Fatalw log a fatal level message, then call os.Exit(1)
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().Fatalw(msg, keysAndValues...)
	return l
}

// Panic log a panic level message, then panics
func (l *Logger) Panic(msg string) *Logger {
	l.logger.Sugar().Panic(msg)
	return l
}

// Panicf log a panic level message, then panics
func (l *Logger) Panicf(msg string, args ...interface{}) *Logger {
	l.logger.Panic(fmt.Sprintf(msg, args...))
	return l
}

// Panicw log a panic level message, then panics
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) *Logger {
	l.logger.Sugar().DPanicw(msg, keysAndValues...)
	return l
}

// Sync sync the log buffers to the file
func (l *Logger) Sync() *Logger {
	l.logger.Sync()
	return l
}

// Helper methods
func initZapLog(app string, development bool, logLevel string, outFilepath []string, errFilepath []string) (*zap.Logger, error) {

	config := zap.Config{
		Level:             getLogLevel(logLevel),
		Development:       development,
		DisableCaller:     false,
		DisableStacktrace: false,
		// Sampling:          &zap.SamplingConfig{},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "t",
			NameKey:    "logger",
			CallerKey:  "calller",
			// FunctionKey:    "func",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEcodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			// EncodeCaller: zapcore.FullCallerEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeName:   zapcore.FullNameEncoder,
			// ConsoleSeparator: "\t",
		},
		OutputPaths:      outFilepath,
		ErrorOutputPaths: errFilepath,
		InitialFields: map[string]interface{}{
			"app": app,
		},
	}
	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(zapLogger)
	return zapLogger, nil
}

func getLogLevel(level string) (logLevel zap.AtomicLevel) {
	logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
		break
	case "INFO":
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
		break
	case "WARN":
		logLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
		break
	case "ERROR":
		logLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
		break
	case "FATAL":
		logLevel = zap.NewAtomicLevelAt(zap.FatalLevel)
		break
	case "PANIC":
		logLevel = zap.NewAtomicLevelAt(zap.PanicLevel)
		break
	case "DPANIC":
		logLevel = zap.NewAtomicLevelAt(zap.DPanicLevel)
		break
	}
	return
}

func formatEcodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

// func (l *Logger) formatLog(args []interface{}) *zap.Logger {
// 	log := l.logger.With(toFields(args))
// 	return log
// }

// func toFields(args []interface{}) zap.Field {
// 	dest := make([]string, 0)
// 	if len(args) > 0 {
// 		for _, v := range args {
// 			dest = append(dest, fmt.Sprintf("%+v", v))
// 		}
// 	}
// 	field := zap.Any("detail", dest)
// 	return field
// }
