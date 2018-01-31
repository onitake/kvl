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

Each logger is more than just a single class handling all the work:
Instead, several components are tied together to form a processing chain.

These components can be divided into four categories:
Frontends, Filters, Formatters and Sinks.

Several ready-to-use loggers are defined in [convenience.go](convenience.go),
they can be used as-is or serve as a starting point for your own loggers.

The various interfaces you need to implement are declared in [logger.go](logger.go).

### Frontends

A Frontend provides a logging interface to applications.

There are no specific requirements on what the API should look like - it can be
anything from a simple passthrough implementation of the Filter interface to
a drop-in replacement for another logger.

After preprocessing, the Frontend should pack its output into a map
and send it to a Filter. See [simple.go](simple.go) for an example.

### Filters

Filters serve as intermediate elements in a more complex chain.
They can be used to enhance log message with additional information, filter
out certain message or format values in specific ways.

For example, one Filter could interpret a key 'level' as the
log level and drop messages below a certain threshold.

Filters should implement the Filter interface and send their output to another
Filter.

### Formatters

A Formatter processes log messages from a Filter or Frontend into a certain
output format.

It serves as the glue point between a Logger and the environment, so it should
implement the Filter interface and send its output to a Sink.

### Sinks

Ultimately, all data needs to end up somewhere (or be discarded).
This is where Sinks come in.

Once the Formatter has processed its data, it will send it through an I/O
channel, to a remote server or simply write to Stdout.

For most purposes this is best abstracted by the `io.Writer` interface.
It is used by many standard APIs and can be implemented easily for a custom
log destination.

If `io.Writer` does not serve your purpose sufficiently, you can also
write a custom Formatter that sends its data to a different type of Sink.

## Copyright + License

KeyValueLogger is Copyright © 2018 by Gregor Riepl

You may use this software under the terms of the MIT license.
See [LICENSE](LICENSE) for details.
