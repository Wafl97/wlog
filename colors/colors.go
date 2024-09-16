package colors

type Color string

const (
	Reset  Color = "\u001B[0m"
	Red    Color = "\u001B[31m"
	Green  Color = "\u001B[32m"
	Yellow Color = "\u001B[33m"
	Blue   Color = "\u001B[34m"
)
