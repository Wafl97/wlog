package tests

import (
	"testing"

	"github.com/Wafl97/wlog"
)

func TestLogger(t *testing.T) {
	logger := wlog.New("TESTING", wlog.LogToConsoleAndFile)

	logger.Debug("Debug")
	logger.Debugf("%s\n", "Debugf")
	logger.Info("Info")
	logger.Infof("%s\n", "Infof")
	logger.Warn("Warn")
	logger.Warnf("%s\n", "Warnf")
	logger.Error("Error")
	logger.Errorf("%s\n", "Errorf")
}
