package wlog

import (
	"fmt"
	"os"
	"path"
	"time"
	"wlog/level"
	"wlog/util"
)

var (
	LogToFile func(level level.Level, s string) error = func(level level.Level, s string) error {
		file, err := os.OpenFile(path.Join("logs", time.Now().Format(time.DateOnly)+".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("logger out: %w", err)
		}
		defer file.Close()
		if _, err = file.WriteString(fmt.Sprintf("[%-5s] %s", level.Name, s)); err != nil {
			return fmt.Errorf("logger out: %w", err)
		}
		return nil
	}

	LogToConsole func(level level.Level, s string) error = func(level level.Level, s string) error {
		fmt.Printf("%s[%-5s] %s%s", level.Color, level.Name, s, util.Reset)
		return nil
	}

	LogToConsoleAndFile func(level level.Level, s string) error = func(level level.Level, s string) error {
		LogToConsole(level, s)
		return LogToFile(level, s)
	}
)

type Logger struct {
	name  string
	level level.Level
	out   func(level level.Level, s string) error
}

func New(name string, out func(level level.Level, s string) error) *Logger {
	return &Logger{
		name:  name,
		level: level.Info,
		out:   out,
	}
}

func (logger *Logger) SetLevel(level level.Level) {
	logger.level = level
}

func (logger *Logger) Debug(message string) {
	if logger.level.Order >= level.Debug.Order {
		logger.out(level.Debug, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Debugf(format string, args ...any) {
	if logger.level.Order >= level.Debug.Order {
		logger.out(level.Debug, fmt.Sprintf("[%s] %s", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Info(message string) {
	if logger.level.Order >= level.Info.Order {
		logger.out(level.Info, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Infof(format string, args ...any) {
	if logger.level.Order >= level.Info.Order {
		logger.out(level.Info, fmt.Sprintf("[%s] %s", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Warn(message string) {
	if logger.level.Order >= level.Warn.Order {
		logger.out(level.Warn, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Warnf(format string, args ...any) {
	if logger.level.Order >= level.Warn.Order {
		logger.out(level.Warn, fmt.Sprintf("[%s] %s", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Error(message string) {
	if logger.level.Order >= level.Error.Order {
		logger.out(level.Error, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Errorf(format string, args ...any) {
	if logger.level.Order >= level.Error.Order {
		logger.out(level.Error, fmt.Sprintf("[%s] %s", logger.name, fmt.Sprintf(format, args...)))
	}
}
