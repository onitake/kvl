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

// Logger is the standard interface for a log source.
// Specific implementations may add more methods.
type Logger interface {
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
	// Only string keys are supported; incompatible keys will be ignored.
	// The output format depends on the chosen logger chain.
	// For example, a JsonLogger outputs in JSON format, while
	// a StdLogger displays in a more human-readable fashion.
	// If available, the key 'message' will be used for the log message.
	// The same applies to 'time', which will override anything that uses the current time.
	Printkv(kv ...interface{})
}

// Filter is the standard interface for a data processor.
// It allows extending and enhancing a log chain with additional functionality.
// The interface is simply a combination of a source and a sink.
type Filter interface {
	Logger
	Sink
}

// Sink is the standard interface for a log sink.
// It is simply an alias to io.Writer.
type Sink interface {
	Write(p []byte) (n int, err error)
}
