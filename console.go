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
	"fmt"
	"time"
)

const (
	ConsoleTimeFormat = "[2006-01-02 15:04:05] "
)

var (
	skipKeys = map[string]bool {
		StdMessageKey: true,
		StdTimeKey: true,
	}
)

// ConsoleFormatter produces human-readable log lines and sends them to a Sink.
// If Sink is not set, nothing will be logged.
// If the PrintTime flag is true and a "time" key is present, it will be logged
// at the start of each log line, in the format given by ConsoleTimeFormat.
// If the PrintKeys flag is true and additional keys besides "message" and
// "time" are present, the will be logged after the message, separated by pipe characters: |
type ConsoleFormatter struct {
	// Sink is the output sink for this formatter.
	Sink io.Writer
	// PrintTime determines if each log will be prepended with the current date and time.
	PrintTime bool
	// PrintKeys determines if key-value pairs will be printed after the message.
	PrintKeys bool
}

func (formatter *ConsoleFormatter) Logkv(kv map[string]interface{}) {
	if formatter.Sink != nil {
		var line string
		if formatter.PrintTime {
			kt, ok := kv[StdTimeKey]
			if ok {
				t, ok := kt.(time.Time)
				if ok {
					line += t.Format(ConsoleTimeFormat)
				}
			}
		}
		km, ok := kv[StdMessageKey]
		if ok {
			m, ok := km.(string)
			if ok {
				line += m
			}
		}
		if formatter.PrintKeys {
			// TODO sort keys first?
			for k, v := range kv {
				// skip if this is one of the keys with special meaning
				if _, ok := skipKeys[k]; !ok {
					line += fmt.Sprintf(" | %s: %v", k, v)
				}
			}
		}
		line += "\n"
		formatter.Sink.Write([]byte(line))
	}
}
