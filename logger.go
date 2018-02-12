// Copyright (c) 2018, Gregor Riepl <onitake@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//    1. Redistributions of source code must retain the above copyright notice, this list of
//       conditions and the following disclaimer.
//
//    2. Redistributions in binary form must reproduce the above copyright notice, this list
//       of conditions and the following disclaimer in the documentation and/or other materials
//       provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package kvl

import (
	"io"
	"os"
)

const (
	// StdMessageKey is the default key to use for log messages.
	// Type: string
	StdMessageKey = "message"
	// StdTimeKey is the default key to use for timestamps.
	// Type: time.Time
	StdTimeKey = "time"
)

// Filter is the standard interface for a data processor.
type Filter interface {
	// Printd accepts dictionaries and processes them.
	// This function should modify the map in-place.
	Printd(dict map[string]interface{})
}

// Formatter is the standard interface for an output generator.
type Formatter interface {
	// Formatd processes the log dictionary into a byte stream and
	// send it to the Sink.
	Formatd(dict map[string]interface{}, sink io.Writer)
}

// Logger is the kvl logging core.
// Its interface is very basic, you should extend it or use one of the Std
// loggers instead.
type Logger struct {
	// Formatter is a Formatter that writes processed output into a Sink.
	// If unset, it simply converts the value StdMessageKey to a byte
	// string and writes it to the Sink.
	Formatter Formatter
	// Sink is an output sink.
	// Defaults to os.Stdout if unset.
	Sink io.Writer
}

// Printd sends a dictionary to the log.
func (logger *Logger) Printd(dict map[string]interface{}) {
	formatter := logger.Formatter
	if formatter == nil {
		formatter = &dummyFormatter{}
	}
	sink := logger.Sink
	if sink == nil {
		sink = os.Stdout
	}
	formatter.Formatd(dict, sink)
}

type dummyFormatter struct{}

func (formatter *dummyFormatter) Formatd(dict map[string]interface{}, sink io.Writer) {
	switch m := dict[StdMessageKey].(type) {
	case string:
		sink.Write([]byte(m))
	case []byte:
		sink.Write(m)
	default:
		sink.Write([]byte("Invalid type for log message"))
	}
}
