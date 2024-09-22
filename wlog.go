package wlog

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Wafl97/wlog/colors"
	"github.com/Wafl97/wlog/format"
	"github.com/Wafl97/wlog/level"
)

var defaultErrorLogger = New("WLOG", LogToConsole)

// Output functions
var (
	LogToFile = func(filename string) func(logLevel level.Level, message any) {
		dir, _ := path.Split(filename)
		os.MkdirAll(dir, 0644)
		return func(logLevel level.Level, message any) {
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				defaultErrorLogger.Error(err)
				return
			}
			defer file.Close()
			if _, err = file.WriteString(fmt.Sprintf("%v", message)); err != nil {
				defaultErrorLogger.Error(err)
				return
			}
		}
	}

	LogToConsole = func(logLevel level.Level, message any) {
		fmt.Printf("%s%s%s", logLevel.Color, message, colors.Reset)
	}

	LogToConsoleAndFile = func(filename string) func(logLevel level.Level, message any) {
		return func(logLevel level.Level, message any) {
			LogToFile(filename)(logLevel, message)
			LogToConsole(logLevel, message)
		}
	}
)

type logFormat func(logger *Logger, logLevel level.Level, message any) string

// logFormats
var (
	_None logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("%v", message)
	}
	_Level logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%-5s] %v", logLevel.Name, message)
	}
	_LevelName logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%-5s] [%s] %v", logLevel.Name, logger.name, message)
	}
	_LevelNameTime logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%-5s] [%s] [%s] %v",
			logLevel.Name, logger.name, time.Now().Format(time.TimeOnly), message)
	}
	_LevelTime logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%-5s] [%s] %v",
			logLevel.Name, time.Now().Format(time.TimeOnly), message)
	}
	_Name logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%s] %v", logger.name, message)
	}
	_NameTime logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%s] [%s] %v",
			logger.name, time.Now().Format(time.TimeOnly), message)
	}
	_Time logFormat = func(logger *Logger, logLevel level.Level, message any) string {
		return fmt.Sprintf("[%s] %v", time.Now().Format(time.TimeOnly), message)
	}
)

// Defaults
var (
	_DefaultFormat logFormat   = _LevelName
	_DefaultLevel  level.Level = level.Info
)

// SetDefaultFormat sets the default format used when creating a new logger.
// The chosen format is only applied to loggers created after calling this.
// Use the SetLogFormat method on existing loggers to change their formats.
func SetDefaultFormat(logFormat format.LogFormat) {
	switch logFormat {
	case format.None:
		_DefaultFormat = _None
	case format.Level:
		_DefaultFormat = _Level
	case format.LevelName:
		_DefaultFormat = _LevelName
	case format.LevelTime:
		_DefaultFormat = _LevelTime
	case format.LevelNameTime:
		_DefaultFormat = _LevelNameTime
	case format.Name:
		_DefaultFormat = _Name
	case format.NameTime:
		_DefaultFormat = _NameTime
	case format.Time:
		_DefaultFormat = _Time
	default:
		_DefaultFormat = _None
	}
}

// SetDefaultLevel sets the default level used when creating a new logger.
// The chosen level is only applied to loggers created after calling this.
// Use the SetLevel method on existing loggers to change their level.
func SetDefaultLevel(logLevel level.Level) {
	_DefaultLevel = logLevel
}

type Logger struct {
	name      string
	level     level.Level
	logFormat logFormat
	output    func(logLevel level.Level, message any)
}

func New(name string, output func(logLevel level.Level, message any)) *Logger {
	if output == nil {
		output = LogToConsole
	}
	return &Logger{
		name:      name,
		level:     _DefaultLevel,
		logFormat: _DefaultFormat,
		output:    output,
	}
}

func (logger *Logger) SetFormat(logFormat format.LogFormat) {
	switch logFormat {
	case format.None:
		logger.logFormat = _None
	case format.Level:
		logger.logFormat = _Level
	case format.LevelName:
		logger.logFormat = _LevelName
	case format.LevelTime:
		logger.logFormat = _LevelTime
	case format.LevelNameTime:
		logger.logFormat = _LevelNameTime
	case format.Name:
		logger.logFormat = _Name
	case format.NameTime:
		logger.logFormat = _NameTime
	case format.Time:
		logger.logFormat = _Time
	default:
		logger.logFormat = _None
	}
}

func (logger *Logger) applyFormat(logLevel level.Level, message any) string {
	return logger.logFormat(logger, logLevel, message)
}

func (logger *Logger) SetLevel(level level.Level) {
	logger.level = level
}

func (logger *Logger) Debug(message any) {
	if logger.level.Order >= level.Debug.Order {
		logger.output(level.Debug, logger.applyFormat(level.Debug, message)+"\n")
	}
}

func (logger *Logger) Debugf(format string, args ...any) {
	if logger.level.Order >= level.Debug.Order {
		logger.output(level.Debug, logger.applyFormat(level.Debug, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Info(message any) {
	if logger.level.Order >= level.Info.Order {
		logger.output(level.Info, logger.applyFormat(level.Info, message)+"\n")
	}
}

func (logger *Logger) Infof(format string, args ...any) {
	if logger.level.Order >= level.Info.Order {
		logger.output(level.Info, logger.applyFormat(level.Info, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Warn(message any) {
	if logger.level.Order >= level.Warn.Order {
		logger.output(level.Warn, logger.applyFormat(level.Warn, message)+"\n")
	}
}

func (logger *Logger) Warnf(format string, args ...any) {
	if logger.level.Order >= level.Warn.Order {
		logger.output(level.Warn, logger.applyFormat(level.Warn, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Error(message any) {
	if logger.level.Order >= level.Error.Order {
		logger.output(level.Error, logger.applyFormat(level.Error, message)+"\n")
	}
}

func (logger *Logger) Errorf(format string, args ...any) {
	if logger.level.Order >= level.Error.Order {
		logger.output(level.Error, logger.applyFormat(level.Error, fmt.Sprintf(format, args...)))
	}
}
