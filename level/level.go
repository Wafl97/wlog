package level

import "github.com/Wafl97/wlog/util"

type Level struct {
	Order uint8
	Name  string
	Color string
}

var (
	Off   = Level{0, "OFF", ""}
	Error = Level{1, "ERROR", util.Red}
	Warn  = Level{2, "WARN", util.Yellow}
	Info  = Level{3, "INFO", util.Green}
	Debug = Level{4, "DEBUG", util.Blue}
)
