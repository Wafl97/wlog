package tests

import (
	"testing"

	"github.com/Wafl97/wlog"
	"github.com/Wafl97/wlog/level"
)

func TestLogger(t *testing.T) {
	logger := wlog.New("TESTING", wlog.LogToConsole)
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
