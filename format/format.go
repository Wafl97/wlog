package format

type LogFormat string

const (
	None          LogFormat = "MESSAGE"                       // Example: Some Message
	Level         LogFormat = "[LEVEL] MESSAGE"               // Example: [INFO ] Some Message
	LevelName     LogFormat = "[LEVEL] [NAME] MESSAGE"        // Example: [INFO ] [SERVER] Some Message
	NameLevel     LogFormat = "[NAME] [LEVEL] MESSAGE"        // Example: [SERVER] [INFO ] Some Message
	LevelNameTime LogFormat = "[LEVEL] [NAME] [TIME] MESSAGE" // Example: [INFO ] [SERVER] [15:05:03] Some Message
	LevelTimeName LogFormat = "[LEVEL] [TIME] [NAME] MESSAGE" // Example: [INFO ] [15:05:03] [SERVER] Some Message
	LevelTime     LogFormat = "[LEVEL] [TIME] MESSAGE"        // Example: [INFO ] [15:05:03] Some Message
	TimeLevel     LogFormat = "[TIME] [LEVEL] MESSAGE"        // Example: [15:05:03] [LEVEL] Some Message
	Name          LogFormat = "[NAME] MESSAGE"                // Example: [SERVER] Some Message
	NameTime      LogFormat = "[NAME] [TIME] MESSAGE"         // Example: [SERVER] [15:05:03] Some Message
	TimeName      LogFormat = "[TIME] [NAME] MESSAGE"         // Example: [15:05:03] [SERVER] Some Message
	Time          LogFormat = "[TIME] MESSAGE"                // Example: [15:05:03] Some Message
)
