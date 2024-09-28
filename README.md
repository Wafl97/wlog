# WLOG

Just a simple logging lib for go programs

## Planned changes

### 1. Make function to set level for all loggers

Have a way to use a single function call to set the level for all existing loggers and all future loggers.

> Sugestion 1:
>
> Make a new function to handle this
>
> ```go
> func SetLevel(logLevel level.Level) {
>   ...
> }
> ```

> Sugestion 2:
> 
> Reuse the existing `SetDefaultLevel`, but add a flag to specify the behaviour
>
> ```go 
> type FormatFlag int
>
> const (
>   ApplyToExisting FormatFlag = 0
>   ApplyToNew      FormatFlag = 1
>   ApplyToAll      FormatFlga = 2
>   // others
> )
> 
> func SetDefaultFormat(logFormat format.LogFormat, applyFlag FormatFlag) {
>   ...
> }
> ```

### 2. Make function to set format for all logger

Have a way to use a single function call to set the formatting for all existing loggers and all future loggers.

> Sugestion 1:
>
> Make a new function to handle this
>
> ```go
> func SetFormat(logFormat format.LogFormat) {
>   ...
> }
> ```

> Sugestion 2:
> 
> Reuse the existing `SetDefaultLevel`, but add a flag to specify the behaviour
>
> ```go 
> type LevelFlag int
>
> const (
>   ApplyToExisting LevelFlag = 0
>   ApplyToNew      LevelFlag = 1
>   ApplyToAll      FormatFlga = 2
>   // others
> )
> 
> func SetDefaultLevel(logLevel level.Level, applyFlag LevelFlag) {
>   ...
> }
> ```

## Usage

Add to go project with the following command:

`go get github.com/Wafl97/wlog`

### Create a simple console logger

By default the logger will only output to the console, if this is desired just pass a 'nil' as the 'output' (second) argument.

```go
import "github.com/Wafl97/wlog"

var logger = wlog.New("MyLogger", nil)
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

var logger = wlog.New("MyLogger", wlog.LogToFile("myLogFile.log"))

func thisWillLog() {
    logger.Info("This will be logged") // in 'myLogFile.log'
}
```

It is also possible to log both to console and file by using `wlog.LogToConsoleAndFile`. Like the one for only file logging, it also requires that you provide a filename.

```go
import "github.com/Wafl97/wlog"

var logger = wlog.New("MyLogger", wlog.LogToConsoleAndFile("myLogFile.log"))

func thisWillLog() {
    logger.Info("This will be logged") // to the console and 'myLogFile.log'
}
```

### Create a custom logger

If you need to send the log to some place that the three provided methods dont cover, you can create a custom handler for that. Just define a function as the 'output' (second) argument.

It is still possible to use the file and/or console logging functionally here by simply calling them in your handler function.

```go
import (
    "github.com/Wafl97/wlog"
    "github.com/Wafl97/wlog/level"
)

var logger = wlog.New("MyLogger", func(logLevel level.Level, message any) {
    // your custom handler logic here
    // wlog.LogToConsole(logLevel, message)
    // wlog.LogToFile("myLogFile.log")(logLevel, message)
    // wlog.LogToConsoleAndFile("myLogFile.log")(logLevel, message)
})

func thisWillLog() {
    logger.Info("This will be logged") // somewhere
}
```

### Level

This module suppers setting the log level for each of the loggers.
> Levels:
> - Off
> - Debug
> - Info
> - Warn
> - Error

Each level, with the exception of `Off`, have a coresponding method and format function, i.e. Info and Infof. By default the level is set to Info.

You can set the global default level with the `wlog.SetDefaultLevel` function.
This will only affect loggers created after calling this.

```go
import (
    "github.com/Wafl97/wlog"
    "github.com/Wafl97/wlog/level"
)

func main() {
    wlog.SetDefaultLevel(level.Debug)
}
```

Alternatively, you can change the level on each logger instance with the `SetLogLevel` method.

```go
import (
    "github.com/Wafl97/wlog"
    "github.com/Wafl97/wlog/level"
)

func main() {
    logger := wlog.New("MyLogger", nil)
    logger.SetLevel(level.Debug)
}
```

### Formatting

This library has support for using differnt formatting for the log outputs.
All of the available formatting options can be seen below:

> Format:
> - None         
> - Level        
> - LevelName    
> - LevelNameTime
> - LevelTime    
> - Name         
> - NameTime     
> - Time         

For setting the format for all new loggers, you can use `wlog.SetDefaultFormat`.
This will only affect any loggers created after calling the function.

```go
import (
    "github.com/Wafl97/wlog"
    "github.com/Wafl97/wlog/format"
)

func main() {
    wlog.SetDefaultFormat(format.NameTime)
}
```

Alternatively, you can set the format for the individual loggers with the `SetFormat` method on the logger.

```go
import (
    "github.com/Wafl97/wlog"
    "github.com/Wafl97/wlog/format"
)

func main() {
    logger := wlog.New("MyLogger", nil)
    logger.SetFormat(format.NameTime)
}
```
