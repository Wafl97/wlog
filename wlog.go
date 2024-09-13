package wlog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type Level struct {
	order uint8
	name  string
	color string
}

const (
	reset  = "\u001B[0m"
	red    = "\u001B[31m"
	green  = "\u001B[32m"
	yellow = "\u001B[33m"
	blue   = "\u001B[34m"
)

var (
	levelOff   = Level{0, "OFF", ""}
	levelError = Level{1, "ERROR", red}
	levelWarn  = Level{2, "WARN", yellow}
	levelInfo  = Level{3, "INFO", green}
	levelDebug = Level{4, "DEBUG", blue}
)

var (
	LogToFile func(level Level, s string) error = func(level Level, s string) error {
		file, err := os.OpenFile(path.Join("logs", time.Now().Format(time.DateOnly)+".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("logger out: %w", err)
		}
		defer file.Close()
		if _, err = file.WriteString(fmt.Sprintf("[%-5s] %s", level.name, s)); err != nil {
			return fmt.Errorf("logger out: %w", err)
		}
		return nil
	}

	LogToConsole func(level Level, s string) error = func(level Level, s string) error {
		fmt.Printf("%s[%-5s] %s%s", level.color, level.name, s, reset)
		return nil
	}

	LogToConsoleAndFile func(level Level, s string) error = func(level Level, s string) error {
		LogToConsole(level, s)
		return LogToFile(level, s)
	}
)

type Logger struct {
	name  string
	level Level
	out   func(level Level, s string) error
}

func New(name string, out func(level Level, s string) error) Logger {
	return Logger{
		name:  name,
		level: levelInfo,
		out:   out,
	}
}

func (logger *Logger) Debug(message string) {
	if logger.level.order >= levelDebug.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Debugf(format string, args ...any) {
	if logger.level.order >= levelDebug.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Info(message string) {
	if logger.level.order >= levelInfo.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Infof(format string, args ...any) {
	if logger.level.order >= levelInfo.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Warn(message string) {
	if logger.level.order >= levelWarn.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Warnf(format string, args ...any) {
	if logger.level.order >= levelWarn.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, fmt.Sprintf(format, args...)))
	}
}

func (logger *Logger) Error(message string) {
	if logger.level.order >= levelError.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, message))
	}
}

func (logger *Logger) Errorf(format string, args ...any) {
	if logger.level.order >= levelError.order {
		logger.out(logger.level, fmt.Sprintf("[%s] %s\n", logger.name, fmt.Sprintf(format, args...)))
	}
}
