package log

var gLogger *Logger

// Public functions for default logger

// Debug log a debug msg
func Debug(msg string) *Logger { gLogger.Debug(msg); return gLogger }

// Debugf log a debug msg
func Debugf(msg string, args ...interface{}) *Logger { gLogger.Debugf(msg, args...); return gLogger }

// Debugw log a debug msg
func Debugw(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Debugw(msg, keysAndValues...)
	return gLogger
}

// Info log a debug msg
func Info(msg string) *Logger { gLogger.Info(msg); return gLogger }

// Infof log a debug msg
func Infof(msg string, args ...interface{}) *Logger { gLogger.Infof(msg, args...); return gLogger }

// Infow log a debug msg
func Infow(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Infow(msg, keysAndValues...)
	return gLogger
}

// Warn log a debug msg
func Warn(msg string) *Logger { gLogger.Warn(msg); return gLogger }

// Warnf log a debug msg
func Warnf(msg string, args ...interface{}) *Logger { gLogger.Warnf(msg, args...); return gLogger }

// Warnw log a debug msg
func Warnw(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Warnw(msg, keysAndValues...)
	return gLogger
}

// Error log a debug msg
func Error(msg string) *Logger { gLogger.Error(msg); return gLogger }

// Errorf log a debug msg
func Errorf(msg string, args ...interface{}) *Logger { gLogger.Errorf(msg, args...); return gLogger }

// Errorw log a debug msg
func Errorw(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Errorw(msg, keysAndValues...)
	return gLogger
}

// Fatal log a debug msg
func Fatal(msg string) *Logger { gLogger.Fatal(msg); return gLogger }

// Fatalf log a debug msg
func Fatalf(msg string, args ...interface{}) *Logger { gLogger.Fatalf(msg, args...); return gLogger }

// Fatalw log a debug msg
func Fatalw(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Fatalw(msg, keysAndValues...)
	return gLogger
}

// Panic log a debug msg
func Panic(msg string) *Logger { gLogger.Panic(msg); return gLogger }

// Panicf log a debug msg
func Panicf(msg string, args ...interface{}) *Logger { gLogger.Panicf(msg, args...); return gLogger }

// Panicw log a debug msg
func Panicw(msg string, keysAndValues ...interface{}) *Logger {
	gLogger.Panicw(msg, keysAndValues...)
	return gLogger
}

// Sync sync the log buffers to the file
func Sync() *Logger {
	gLogger.Sync()
	return gLogger
}
