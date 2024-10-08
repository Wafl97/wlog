package tests

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/Wafl97/wlog"
	"github.com/Wafl97/wlog/format"
	"github.com/Wafl97/wlog/level"
)

func TestLogger(t *testing.T) {
	logger := wlog.New("TESTING",
		wlog.LogToConsoleAndFile(path.Join("test", "logs", time.Now().Format(time.DateOnly)+".log")))
	logger.SetFormat(format.LevelNameTime)
	logger.SetLevel(level.Debug)

	logger.Debug("Debug")
	logger.Debugf("%s\n", "Debugf")
	logger.Info("Info")
	logger.Infof("%s\n", "Infof")
	logger.Warn("Warn")
	logger.Warnf("%s\n", "Warnf")
	logger.Error("Error")
	logger.Errorf("%s\n", "Errorf")
}

func TestFormatting(t *testing.T) {
	now := time.Now().Format(time.TimeOnly)
	expected := []string{
		"Is this working?\n",                                            //None
		"[INFO ] Is this working?\n",                                    //Level
		"[INFO ] [MY LOGGER] Is this working?\n",                        //LevelName
		fmt.Sprintf("[INFO ] [MY LOGGER] [%s] Is this working?\n", now), //LevelNameTime
		fmt.Sprintf("[INFO ] [%s] Is this working?\n", now),             //LevelTime
		"[MY LOGGER] Is this working?\n",                                //Name
		fmt.Sprintf("[MY LOGGER] [%s] Is this working?\n", now),         //NameTime
		fmt.Sprintf("[%s] Is this working?\n", now),                     //Time
	}

	logCatcher := make([]string, 0, len(expected))

	myLogger := wlog.New("MY LOGGER", func(logLevel level.Level, message any) {
		logCatcher = append(logCatcher, fmt.Sprintf("%v", message))
	})
	myLogger.SetFormat(format.None)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.Level)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.LevelName)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.LevelNameTime)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.LevelTime)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.Name)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.NameTime)
	myLogger.Info("Is this working?")
	myLogger.SetFormat(format.Time)
	myLogger.Info("Is this working?")

	for index := range expected {
		if logCatcher[index] != expected[index] {
			t.Errorf("expected `%s`, but got `%s`", expected[index], logCatcher[index])
		}
	}
}
