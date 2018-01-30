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
	"testing"
)

func TestSliceToMap(t *testing.T) {
	c01 := []interface{}{}
	r01 := SliceToMap(c01)
	if len(r01) != 0 {
		t.Error("t01: result should be empty")
	}

	c02 := []interface{}{ "message" }
	r02 := SliceToMap(c02)
	if len(r02) != 0 {
		t.Error("t02: result should be empty")
	}

	c03 := []interface{}{ 23, 777 }
	r03 := SliceToMap(c03)
	if len(r03) != 0 {
		t.Error("t03: result should be empty")
	}

	c04 := []interface{}{ "message", "hello" }
	r04 := SliceToMap(c04)
	if len(r04) != 1 {
		t.Error("t04: result should have one key")
	}
	if _, ok := r04["message"]; !ok {
		t.Error("t04: message key not found")
	}
	if r04["message"] != "hello" {
		t.Error("t04: message has invalid value")
	}

	c05 := []interface{}{ "message", "hello", "theanswer", 42 }
	r05 := SliceToMap(c05)
	if len(r05) != 2 {
		t.Error("t05: result should have two keys")
	}
}
