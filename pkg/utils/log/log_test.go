package log

import (
	"os"
	"testing"
	"time"
)

func TestDefaultLog(t *testing.T) {
	Debug(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Debugf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Debugw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Info(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Infof("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Infow("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Error(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Errorf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Errorw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())

	ZapInfo("message", Int("code", 200), String("Src", "192.168.10.64"), String("Method", "GET"),
		String("Url", "/api/v1/health"), Time("Time", time.Now()))
}

func TestCustomizedLog(t *testing.T) {
	logger := New(os.Stderr, JsonFormat, DebugLevel, WithCaller(true))
	ResetDefault(logger)

	Debug(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Debugf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Debugw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Info(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Infof("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Infow("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Error(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Errorf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Errorw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())

	ZapInfo("message", Int("code", 200), String("Src", "192.168.10.64"), String("Method", "GET"),
		String("Url", "/api/v1/health"), Time("Time", time.Now()))
}

func TestCustomizedTeeLog(t *testing.T) {
	file, err := os.OpenFile("/tmp/go_proj_boot.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	var tops = []TeeOption{
		{
			W: os.Stderr,
			Lef: func(lvl Level) bool {
				return lvl <= InfoLevel
			},
			F: ConsoleFormat,
		},
		{
			W: file,
			Lef: func(lvl Level) bool {
				return lvl > InfoLevel
			},
			F: JsonFormat,
		},
	}

	logger := NewTee(tops, WithCaller(true))
	ResetDefault(logger)

	Debug(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Debugf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Debugw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Info(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Infof("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Infow("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Error(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Errorf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Errorw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())

	ZapInfo("message", Int("code", 200), String("Src", "192.168.10.64"), String("Method", "GET"),
		String("Url", "/api/v1/health"), Time("Time", time.Now()))
}

func TestCustomizedTeeWithRotateLog(t *testing.T) {
	var tops = []TeeWithRotateOption{
		{
			Filename:   "/tmp/go_proj_boot/go_proj_boot.log",
			MaxSize:    10,
			MaxAge:     16,
			MaxBackups: 16,
			Compress:   true,
			Lef: func(lvl Level) bool {
				return lvl >= InfoLevel
			},
			F: ConsoleFormat,
		},
	}

	logger := NewTeeWithRotate(tops, WithCaller(true))
	ResetDefault(logger)

	Debug(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Debugf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Debugw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Info(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Infof("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Infow("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())
	Error(200, ", 192.168.10.64", ", GET", ", /api/v1/health, ", time.Now())
	Errorf("%d, %s, %s, %s, %s", 200, "192.168.10.64", "GET", "/api/v1/health", time.Now())
	Errorw("message", "code", 200, "Src", "192.168.10.64", "Method", "GET", "Url", "/api/v1/health", "Time", time.Now())

	ZapInfo("message", Int("code", 200), String("Src", "192.168.10.64"), String("Method", "GET"),
		String("Url", "/api/v1/health"), Time("Time", time.Now()))
}