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
)

// SimpleLogger is a very basic log frontend that packs strings into a map
// and passes them to a Filter.
//
// The Print* family of log methods are modeled after log.Logger.
type StdLogger struct {
	// Yes, we extend Logger.
	Logger
}

// Print outputs each argument on a separate log line.
func (logger *StdLogger) Print(v ...interface{}) {
	for _, line := range v {
		logger.Printd(map[string]interface{}{
			StdMessageKey: line,
		})
	}
}

// Println is simply an alias for Print, as log messages are always terminated with a newline.
// Provided for compatibility with log.Logger.
func (logger *StdLogger) Println(v ...interface{}) {
	logger.Print(v)
}

// Printf formats a string like log.Printf does, then logs it as a single
// log line.
func (logger *StdLogger) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	logger.Print(message)
}

// Printkv packs a list of key-value pairs and sends it to the Filter chain.
// Only string keys are supported; incompatible keys or a missing value
// at the end will be silently ignored.
func (logger *StdLogger) Printkv(kv ...interface{}) {
	mkv := SliceToMap(kv)
	logger.Printd(mkv)
}
