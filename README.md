# KeyValueLogger

A light-weight Go logging package.

## Key Features

* Light-weight
* Extensible
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
logger.Printm(map[string]interface{}{"message": "Take one down", "pass": "around"})
```
```console
[2006-01-02 15:04:05] Welcome to KeyValueLogger!
[2006-01-02 15:04:06] A float value: 1.0
[2006-01-02 15:04:08] Bottles on the wall | count: 99 | content: beer
[2006-01-02 15:04:08] Take one down | pass: around
```

For machine processing, JSON is more useful:
```go
jsonlogger := kvl.NewStdJsonLog()
logger.Print("Welcome to KeyValueLogger!")
logger.Printf("A float value: %f", 1.0)
logger.Printkv("message", "Bottles on the wall", "count", 99, "content", "beer")
logger.Printm(map[string]interface{}{"message": "Take one down", "pass": "around"})
```
```json
{"time":"2006-01-02T15:04:05Z07:00","message":"Welcome to KeyValueLogger!"}
{"time":"2006-01-02T15:04:06Z07:00","message":"A float value: 1.0"}
{"time":"2006-01-02T15:04:08Z07:00","message":"Bottles on the wall","count":99,"content":"beer"}
{"time":"2006-01-02T15:04:09Z07:00","message":"Take one down","pass":"around"}
```

## Extend

The core of a logger serves as a skeleton for Frontends, Filters, Formatters
and Sinks. Everything is extensible.

For example, the Std logger family provides convenient Print* functions
as a Frontend for application programs.
Filters can alter data or add additional information.
Formatters turn dictionaries into byte strings, and Sinks send the result to
a socket, file or other output device.

See [convenience.go](convenience.go) for examples, or use the Std loggers as-is.

The core logger and the interfaces provided by Filters and Formatters
is defined in [logger.go](logger.go).

### Frontends

A Frontend provides a logging interface to applications.

There are no specific requirements on what the API should look like - it can be
anything from a simple passthrough implementation of the Filter interface to
a drop-in replacement for another logger.

Frontend classes should extend the core Logger class and use the Printd
function to send a dictionary to the log.

### Filters

Filters serve as intermediate elements in a more complex chain.
They can be used to enhance log message with additional information, filter
out certain message or format values in specific ways.

For example, one Filter could interpret a key 'level' as the
log level and drop messages below a certain threshold.

Filters should implement the Filter interface and modify information in-place.

### Formatters

A Formatter processes log messages from a Filter or Frontend into a certain
output format.

It serves as the glue point between a Logger and the environment, so it should
act like a Filter, but but send its output to a sink instead.

### Sinks

Once the Formatter has processed data, the result will be sent through an I/O
channel, to a remote server or simply be written to Stdout.

This makes `io.Writer` interface a natural candidate for such a sink.

## Copyright + License

KeyValueLogger is Copyright © 2018 by Gregor Riepl

You may use this software under the terms of the MIT license.
See [LICENSE](LICENSE) for details.
