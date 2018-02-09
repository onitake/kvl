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
	"bytes"
	"testing"
)

func jsonTest(t *testing.T, testno string, query map[string]interface{}, expected ...[]byte) {
	buffer := &bytes.Buffer{}
	testee := &JsonFormatter{}
	testee.Formatd(query, buffer)
	result := buffer.Bytes()
	ok := false
	for _, ex1 := range expected {
		if bytes.Equal(ex1, result) {
			ok = true
		} else {
			t.Logf("%s: no match. expected: '%s' got: '%s'", testno, ex1, result)
		}
	}
	if !ok {
		t.Errorf("%s: no matching result", testno)
	}
}

func TestJsonFormatter(t *testing.T) {
	q01 := map[string]interface{}{}
	x01 := []byte("{}\n")
	jsonTest(t, "t01", q01, x01)

	q02 := map[string]interface{}{
		"message": "test 02",
	}
	x02 := []byte("{\"message\":\"test 02\"}\n")
	jsonTest(t, "t02", q02, x02)

	q03 := map[string]interface{}{
		"message": "test 03",
		"value":   1234567,
	}
	x03a := []byte("{\"value\":1234567,\"message\":\"test 03\"}\n")
	x03b := []byte("{\"message\":\"test 03\",\"value\":1234567}\n")
	jsonTest(t, "t03", q03, x03a, x03b)
}
