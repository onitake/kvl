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
	"os"
	"time"
)

// NewStdLog creates a simple StdOut logger suitable for human consumption.
// Key-Value pairs are separated by a pipe character: |
// Each log line is prepended with the current date and time.
func NewStdLog() *StdLogger {
	return &StdLogger{
		Logger: &MultiFilter{
			Filters: []Filter{
				&AddTimeFilter{},
			},
			Logger: &Logger{
				Formatter: &ConsoleFormatter{
					PrintTime: true,
					PrintKeys: true,
					SortKeys:  true,
				},
				Sink: os.Stdout,
			},
		},
	}
}

// NewJsonLog creates a JSON logger that places each message into the
// StdMessageKey key and adds StdMessageKey with the current time in RFC3339 format.
func NewJsonLog() *StdLogger {
	return &StdLogger{
		Logger: &MultiFilter{
			Filters: []Filter{
				&AddTimeFilter{
					TimeFormat: time.RFC3339,
				},
			},
			Logger: &Logger{
				Formatter: &JsonFormatter{},
				Sink:      os.Stdout,
			},
		},
	}
}
