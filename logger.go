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

// Logger is a generic type for all loggers.
// You should provide a reasonable implementation for all functions, to allow
// full interchangeability of individual loggers.
// If a particular function makes no sense, provide
// a minimal implementation that preserves as much information as possible.
type Logger interface {
	// Write outputs an encoded string.
	// Sinks should handle this method, other loggers should simply pass it
	// on to their children.
	Write(p []byte) (n int, err error)
	// Print outputs each argument on a separate log line.
	Print(v ...interface{})
	// Println is simply an alias for Print,
	// as log messages are always terminated with a newline.
	// Provided for compatibility with log.Logger.
	Println(v ...interface{})
	// Printf works the same as log.Printf, but the end result
	// is sent through kvl.Print.
	Printf(format string, v ...interface{})
	// Printkv expects a list of alternating key-value pairs.
	// The output format depends on the chosen logger chain.
	// For example, a JsonLogger outputs in JSON format, while
	// a StdLogger displays in a more human-readable fashion.
	Printkv(kv ...interface{})
	// Backers returns a list of all backing loggers attached to this logger.
	// How they are used depends on the specific logger.
	// Most loggers will have only a single Next entry, while pure sinks
	// have none.
	Backers() []Logger
	// SetBacker assigns the backing logger.
	// If multiple backing loggers are supported, this assigns the 'main'
	// logger. Additional functions should be provided by the logger to configure others.
	SetBacker(backing Logger)
	// Head returns a pointer to the last 'main' logger in a chain.
	// This is usually a pure sink and Head can be used to update the final destination.
	// Secondary loggers can not be reached with this function.
	Head() *Logger
}
