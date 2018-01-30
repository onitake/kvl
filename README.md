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

Loggers can be connected to form a processing chain.
These chains are highly customisable through standard interfaces:

Sources, Filters and Sinks.

Individual loggers can act as one or more of these.

For the lazy, a number of ready-to-use chains are provided. Take a look at
[convenience.go](convenience.go) to get started.

If you would rather write your own loggers and processing chains, please look at the
interface declarations in [logger.go](logger.go). Note that the
interfaces closely follow established standards to allow reuse of
existing APIs.

### Sources

A source provides a logging interface to applications.

Its API is based on the `log` standard package and ideads from [logxi](https://github.com/mgutz/logxi),
but in contrast to them, no log levels or side-effects are provided.
If such functionality is desired, it can be added easily through a custom source.

### Filters

Filters are intended as intermediate elements in a more complex chain.
They can be used to enhance log message with additional information or
format messages in specific ways.

For example, one filter could interpret a key 'level' as the
log level and filter messages below a threshold.
Or, a filter could process messages into coloured text, depending
on flags, keys or a log level.

### Sinks

Ultimately, all data needs to end up somewhere (or be discarded).

A sink is nothing more than an `io.Writer` - which means you can
attach open files, stdout, network sockets and other things.

## Copyright + License

KeyValueLogger is Copyright © 2018 by Gregor Riepl

You may use this software under the terms of the MIT license.
See [LICENSE](LICENSE) for details.
