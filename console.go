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
	"fmt"
	"time"
)

const (
	timeFormat = "[2006-01-02 15:04:05] "
)

// ConsoleFilter produces human-readable log lines and sends them to a Sink.
// If the Sink is not set, nothing will be logged.
// If the PrintTime flag is true, each line will be prepended with the current date and time.
// If the PrintKeys flag is true, all key-value pairs will be printed after the message,
// separated by pipe characters: |
type ConsoleFilter struct {
	// Sink is the output sink for this filter.
	Sink Sink
	// PrintTime determines if each log will be prepended with the current date and time.
	PrintTime bool
	// PrintKeys determines if key-value pairs will be printed after the message.
	PrintKeys bool
}

func (filter *ConsoleFilter) Write(p []byte) (n int, err error) {
	if filter.Sink != nil {
		return filter.Sink.Write(p)
	}
	return 0, nil
}

func (filter *ConsoleFilter) printmkv(kv map[string]interface{}) {
	var line string
	if filter.PrintTime {
		var t time.Time
		kt, ok := kv["time"]
		if ok {
			t, ok = kt.(time.Time)
		}
		if !ok {
			t = time.Now()
		}
		line += t.Format(timeFormat)
	}
	var m string
	km, ok := kv["message"]
	if ok {
		m, ok = km.(string)
	}
	line += m
	if filter.PrintKeys {
		for k, v := range kv {
			if k != "message" {
				line += fmt.Sprintf(" | %s: %s", k, v)
			}
		}
	}
	line += "\n"
	filter.Write([]byte(line))
}

func (filter *ConsoleFilter) Print(v ...interface{}) {
	for _, line := range v {
		filter.printmkv(map[string]interface{}{
			"message": line,
		})
	}
}

func (filter *ConsoleFilter) Println(v ...interface{}) {
	filter.Print(v)
}

func (filter *ConsoleFilter) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	filter.Print(message)
}

func (filter *ConsoleFilter) Printkv(kv ...interface{}) {
	mkv := SliceToMap(kv)
	filter.printmkv(mkv)
}
