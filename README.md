# KeyValueLogger

A light-weight Go logging package.

## Key Features

* Light-weight
* Highly extensible
* No dependencies on other packages
* Produces machine- and human-readable logs

## Synposis

Import the package first:
```go
import "github.com/onitake/kvl"

```
Try a simple console log:
```go
logger := kvl.NewStdLog()
logger.Print("Welcome to KeyValueLogger!")
logger.Printf("A float value: %f", 1.0)
logger.Printkv("message", "Bottles on the wall", "count", 99, "content", "beer")
```
```console
[2006-01-02 15:04:05] Welcome to KeyValueLogger!
[2006-01-02 15:04:06] A float value: 1.0
[2006-01-02 15:04:08] bottles on the wall | count: 99 | content: beer
```

For log-shipping, JSON is more useful:
```go
jsonlogger := kvl.NewJsonLog()
logger.Print("Welcome to KeyValueLogger!")
logger.Printf("A float value: %f", 1.0)
logger.Printkv("message", "Bottles on the wall", "count", 99, "content", "beer")
```
```json
{"time":"2006-01-02T15:04:05Z07:00","message":"Welcome to KeyValueLogger!"}
{"time":"2006-01-02T15:04:06Z07:00","message":"A float value: 1.0"}
{"time":"2006-01-02T15:04:08Z07:00","message":"bottles on the wall","count":99,"content":"beer"}
```

## Mix, Match + Extend

Each logger can act as both source, sink and filter for other loggers.

Chain them together to create a logger customised to *your* application!

For the lazy, a number of ready-to-use chains are provided. Take a look at
[convenience.go](convenience.go) to get started.

If you would prefer to write your own loggers and sinks, please look at the
interface declaration in [logger.go](logger.go).

## Copyright + License

KeyValueLogger is Copyright © 2018 by Gregor Riepl

You may use this software under the terms of the MIT license.
See [LICENSE](LICENSE) for details.
