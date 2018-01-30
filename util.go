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

// SliceToMap converts a list of interleaved key-value pairs to a map.
func SliceToMap(kv []interface{}) map[string]interface{} {
	values := make(map[string]interface{})
	// take care that we don't overrun - there must be two next arguments
	for i := 0; i + 1 < len(kv); i += 2 {
		key, ok := kv[i].(string)
		// ignore non-string keys
		if ok {
			values[key] = kv[i + 1]
		}
	}
	return values
}
