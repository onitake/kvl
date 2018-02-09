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
	"io"
	"time"
)

const (
	ConsoleTimeFormat  = "[2006-01-02 15:04:05]"
	invalidMessageType = "(message not printable)"
)

var (
	skipKeys = map[string]bool{
		StdMessageKey: true,
		StdTimeKey:    true,
	}
)

// ConsoleFormatter produces human-readable log lines and sends them to a Sink.
//
// If the PrintTime flag is set and StdTimeKey is present, the key's contents
// will be logged at the start of each line, followed by a space.
// If it is present but a time.Time object, it will be formatted according
// to ConsoleTimeFormat.
// If it is not present, nothing will be printed.
//
// If the PrintKeys flag is true and additional keys besides StdMessageKey and
// StdTimeKey are present, they will be logged after the message, separated
// by pipe characters: |
type ConsoleFormatter struct {
	// PrintTime determines if each log will be prepended with date and time.
	PrintTime bool
	// PrintKeys determines if key-value pairs will be printed after the message.
	PrintKeys bool
	// SortKeys determines if keys should be sorted alphabetically.
	SortKeys bool
}

func printKey(line *string, k string, v interface{}) {
	// skip if this is one of the keys with special meaning
	if _, ok := skipKeys[k]; !ok {
		*line += fmt.Sprintf(" | %s: %v", k, v)
	}
}

func (formatter *ConsoleFormatter) Formatd(dict map[string]interface{}, sink io.Writer) {
	var line string
	if formatter.PrintTime {
		switch t := dict[StdTimeKey].(type) {
		case string:
			line += t
			line += " "
		case time.Time:
			line += t.Format(ConsoleTimeFormat)
			line += " "
		default:
			// pass
			//line += time.Now().Format(ConsoleTimeFormat)
		}
	}
	switch m := dict[StdMessageKey].(type) {
	case string:
		line += m
	default:
		line += invalidMessageType
	}
	if formatter.PrintKeys {
		if formatter.SortKeys {
			for _, k := range OrderedStringKeys(dict) {
				printKey(&line, k, dict[k])
			}
		} else {
			for k, v := range dict {
				printKey(&line, k, v)
			}
		}
	}
	line += "\n"
	sink.Write([]byte(line))
}
