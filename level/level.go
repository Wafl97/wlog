package level

import "github.com/Wafl97/wlog/colors"

type Level struct {
	Order uint8
	Name  string
	Color colors.Color
}

var (
	Off   = Level{0, "OFF", ""}
	Error = Level{1, "ERROR", colors.Red}
	Warn  = Level{2, "WARN", colors.Yellow}
	Info  = Level{3, "INFO", colors.Green}
	Debug = Level{4, "DEBUG", colors.Blue}
)
