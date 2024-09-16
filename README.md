# WLOG

Just a simple logging lib for go programs

## Usage

Add to go project with following the command:
`go get github.com/Wafl97/wlog`

### Create a simple console logger

By default the logger will only output to the console, if this is desired just pass a 'nil' as the 'output' (second) argument.

```go
import "github.com/Wafl97/wlog"

var logger := wlog.New("MyLogger", nil)
// If you want to be explicit, use the line below
// var logger := wlog.New("MyLogger", wlog.LogToConsole)

func thisWillLog() {
    logger.Info("This will be logged") // to the comsole
}
```

### Create a simple file logger

You can make the logger write to file by using `wlog.LogToFile`.
This requires you give it a filename to use for the file.

```go
import "github.com/Wafl97/wlog"

var logger := wlog.New("MyLogger", wlog.LogToFile("myLogFile.log"))

func thisWillLog() {
    logger.Info("This will be logged") // in 'myLogFile.log'
}
```

It is also possible to log both to console and file by using `wlog.LogToConsoleAndFile`. Like the one for only file logging, it also requires that you provide a filename.

```go
import "github.com/Wafl97/wlog"

var logger := wlog.New("MyLogger", wlog.LogToConsoleAndFile("myLogFile.log"))

func thisWillLog() {
    logger.Info("This will be logged") // to the console and 'myLogFile.log'
}
```

### Create a custom logger

If you need to send the log to some place that the three provided methods dont cover, you can create a custom handler for that. Just define a function as the 'output' (second) argument.

It is still possible to use the file and/or console logging functionally here by simply calling them in your handler function.

```go
import "github.com/Wafl97/wlog"

var logger := wlog.New("MyLogger", func(logLevel level.Level, message any) {
    // your custom handler logic here
    // wlog.LogToConsole(logLevel, message)
    // wlog.LogToFile("myLogFile.log")(logLevel, message)
    // wlog.LogToConsoleAndFile("myLogFile.log")(logLevel, message)
})

func thisWillLog() {
    logger.Info("This will be logged") // somewhere
}
```