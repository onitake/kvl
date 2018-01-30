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
)

// FileLogger is pure sink will print log lines to a file or pseudo-file.
// Initialise with os.Stdout for example, to obtain a stdout logger.
// Beware that only string arguments are supported: Printf will ignore
// any extra arguments and other types will simply be skipped over.
// Printkv ignores keys and only prints all values, separated by a space.
type FileLogger struct {
	Output *os.File
}

func (logger *FileLogger) Write(p []byte) (n int, err error) {
	return logger.Output.Write(p)
}

func (logger *FileLogger) Print(v ...interface{}) {
	for _, line := range v {
		logger.Write([]byte(line))
	}
}
func (logger *FileLogger) Println(v ...interface{}) {
	logger.Print(v)
}
func (logger *FileLogger) Printf(format string, v ...interface{}) {
}
func (logger *FileLogger) Printkv(kv ...interface{}) {
}
func (logger *FileLogger) Backers() []Logger {
}
func (logger *FileLogger) SetBacker(backing Logger) {
}
func (*FileLogger) Head() *Logger {
}

