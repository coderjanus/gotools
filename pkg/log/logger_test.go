package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	logfile = "/tmp/zap.log"
)

type testInfo struct {
	level       string
	expectLines int
}

func TestLogLevel(t *testing.T) {
	infos := []testInfo{
		{level: "debug", expectLines: 12},
		{level: "info", expectLines: 9},
		{level: "warn", expectLines: 6},
		{level: "error", expectLines: 3},
	}
	for _, info := range infos {
		os.Remove(logfile)
		l, _ := InitLogger(Config{
			AppName:     "LoggerTest",
			LogLevel:    info.level,
			OutFile:     logfile,
			ErrFile:     logfile,
			Console:     false,
			Development: true,
		})

		writeLogMsg(l)

		content, err := ioutil.ReadFile(logfile)
		if err != nil {
			t.FailNow()
		}
		str := string(content)
		lines := strings.Split(strings.TrimSpace(str), "\n")
		if len(lines) != info.expectLines {
			t.Errorf("expect %d lines, actural %d line", info.expectLines, len(lines))
		}
	}
}

func TestPublic(t *testing.T) {
	infos := []testInfo{
		{level: "debug", expectLines: 12},
		{level: "info", expectLines: 9},
		{level: "warn", expectLines: 6},
		{level: "error", expectLines: 3},
	}
	for _, info := range infos {
		os.Remove(logfile)
		_, _ = InitLogger(Config{
			AppName:     "LoggerTest",
			LogLevel:    info.level,
			OutFile:     logfile,
			ErrFile:     logfile,
			Console:     false,
			Development: true,
		})

		writePublicLogMsg()

		content, err := ioutil.ReadFile(logfile)
		if err != nil {
			t.FailNow()
		}
		str := string(content)
		lines := strings.Split(strings.TrimSpace(str), "\n")
		if len(lines) != info.expectLines {
			t.Errorf("expect %d lines, actural %d line", info.expectLines, len(lines))
		}
		// time.Sleep(5 * time.Second)
	}
}

func writeLogMsg(l *Logger) {
	l.Debug("log Debug msg")
	l.Debugf("log Debug msg with %s", "args")
	l.Debugw("log Debug msg with fields", "key", "value")
	l.Info("log Info msg")
	l.Infof("log Info msg with %s", "args")
	l.Infow("log Info msg with fields", "key", "value")
	l.Warn("log Warn msg")
	l.Warnf("log Warn msg with %s", "args")
	l.Warnw("log Warn msg with fields", "key", "value")
	l.Error("log Error msg")
	l.Errorf("log Error msg with %s", "args")
	l.Errorw("log Error msg with fields", "key", "value")
	l.Sync()
}

func writePublicLogMsg() {
	Debug("log Debug msg")
	Debugf("log Debug msg with %s", "args")
	Debugw("log Debug msg with fields", "key", "value")
	Info("log Info msg")
	Infof("log Info msg with %s", "args")
	Infow("log Info msg with fields", "key", "value")
	Warn("log Warn msg")
	Warnf("log Warn msg with %s", "args")
	Warnw("log Warn msg with fields", "key", "value")
	Error("log Error msg")
	Errorf("log Error msg with %s", "args")
	Errorw("log Error msg with fields", "key", "value")
	Sync()
}
