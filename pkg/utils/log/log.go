package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var std = New(os.Stderr, ConsoleFormat, InfoLevel, WithCaller(true))

var (
	Debug     = std.Debug
	Debugf    = std.Debugf
	Debugw    = std.Debugw
	Info      = std.Info
	Infof     = std.Infof
	Infow     = std.Infow
	Warn      = std.Warn
	Warnf     = std.Warnf
	Warnw     = std.Warnw
	Error     = std.Error
	Errorf    = std.Errorf
	Errorw    = std.Errorw
	DPanic    = std.DPanic
	DPanicf   = std.DPanicf
	DPanicw   = std.DPanicw
	Panic     = std.Panic
	Panicf    = std.Panicf
	Panicw    = std.Panicw
	Fatal     = std.Fatal
	Fatalf    = std.Fatalf
	Fatalw    = std.Fatalw
	ZapDebug  = std.ZapDebug
	ZapInfo   = std.ZapInfo
	ZapWarn   = std.ZapWarn
	ZapError  = std.ZapError
	ZapDPanic = std.ZapDPanic
	ZapPanic  = std.ZapPanic
	ZapFatal  = std.ZapFatal
	Sync      = std.Sync
)

func Default() *Logger {
	return std
}

func ResetDefault(l *Logger) {
	std = l
	Debug = std.Debug
	Debugf = std.Debugf
	Debugw = std.Debugw
	Info = std.Info
	Infof = std.Infof
	Infow = std.Infow
	Warn = std.Warn
	Warnf = std.Warnf
	Warnw = std.Warnw
	Error = std.Error
	Errorf = std.Errorf
	Errorw = std.Errorw
	DPanic = std.DPanic
	DPanicf = std.DPanicf
	DPanicw = std.DPanicw
	Panic = std.Panic
	Panicf = std.Panicf
	Panicw = std.Panicw
	Fatal = std.Fatal
	Fatalf = std.Fatalf
	Fatalw = std.Fatalw
	ZapDebug = std.ZapDebug
	ZapInfo = std.ZapInfo
	ZapWarn = std.ZapWarn
	ZapError = std.ZapError
	ZapDPanic = std.ZapDPanic
	ZapPanic = std.ZapPanic
	ZapFatal = std.ZapFatal
	Sync = std.Sync
}

type Format int8

const (
	ConsoleFormat Format = iota
	JsonFormat
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel
	InfoLevel   Level = zap.InfoLevel
	WarnLevel   Level = zap.WarnLevel
	ErrorLevel  Level = zap.ErrorLevel
	DPanicLevel Level = zap.DPanicLevel
	PanicLevel  Level = zap.PanicLevel
	FatalLevel  Level = zap.FatalLevel
)

type Field = zapcore.Field

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	AddStacktrace = zap.AddStacktrace
)

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
)

type Logger struct {
	l  *zap.Logger
	sl *zap.SugaredLogger
}

func (l *Logger) Debug(args ...interface{}) {
	l.sl.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sl.Debugf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sl.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sl.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sl.Infof(template, args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sl.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sl.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sl.Warnf(template, args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sl.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sl.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sl.Errorf(template, args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sl.Errorw(msg, keysAndValues...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.sl.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.sl.DPanicf(template, args...)
}

func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sl.DPanicw(msg, keysAndValues...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.sl.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.sl.Panicf(template, args...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sl.Panicw(msg, keysAndValues...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sl.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sl.Fatalf(template, args...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sl.Fatalw(msg, keysAndValues...)
}

func (l *Logger) ZapDebug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) ZapInfo(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) ZapWarn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) ZapError(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) ZapDPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func (l *Logger) ZapPanic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) ZapFatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func New(writer io.Writer, format Format, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core zapcore.Core
	if format == ConsoleFormat {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(writer),
			zapcore.Level(level),
		)
	} else if format == JsonFormat {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(writer),
			zapcore.Level(level),
		)
	} else {
		panic("invalid format")
	}

	logger := &Logger{
		l: zap.New(core, opts...),
	}
	logger.sl = logger.l.Sugar()

	return logger
}

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	W   io.Writer
	Lef LevelEnablerFunc
	F   Format
}

func NewTee(tops []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	for _, top := range tops {
		top := top
		if top.W == nil {
			panic("the writer is nil")
		}

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		})

		var core zapcore.Core
		if top.F == ConsoleFormat {
			core = zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.AddSync(top.W),
				lv,
			)
		} else if top.F == JsonFormat {
			core = zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(top.W),
				lv,
			)
		} else {
			panic("invalid format")
		}

		cores = append(cores, core)
	}

	logger := &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
	logger.sl = logger.l.Sugar()

	return logger
}

type TeeWithRotateOption struct {
	Filename   string
	MaxSize    int  // the maximum size in megabytes of the log file before it gets rotated
	MaxAge     int  // the maximum number of days to retain old log files based on the timestamp encoded in their filename
	MaxBackups int  // the maximum number of old log files to retain
	Compress   bool // the rotated log files should be compressed or not
	Lef        LevelEnablerFunc
	F          Format
}

func NewTeeWithRotate(tops []TeeWithRotateOption, opts ...Option) *Logger {
	var cores []zapcore.Core

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	for _, top := range tops {
		top := top
		if top.Filename == "" {
			panic("the log file name is emtpy")
		}

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		})

		var core zapcore.Core

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,
			MaxSize:    top.MaxSize,
			MaxBackups: top.MaxBackups,
			MaxAge:     top.MaxAge,
			Compress:   top.Compress,
		})

		if top.F == ConsoleFormat {
			core = zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.AddSync(w),
				lv,
			)
		} else if top.F == JsonFormat {
			core = zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(w),
				lv,
			)
		} else {
			panic("invalid format")
		}

		cores = append(cores, core)
	}

	logger := &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
	logger.sl = logger.l.Sugar()

	return logger
}
