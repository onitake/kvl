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
	"fmt"
	"testing"
	"time"
)

func consoleTest(t *testing.T, testno string, create func() *ConsoleFormatter, query map[string]interface{}, expected ...[]byte) {
	buffer := &bytes.Buffer{}
	testee := create()
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

func TestConsoleFormatter(t *testing.T) {
	c01 := func() *ConsoleFormatter {
		return &ConsoleFormatter{}
	}
	q01 := map[string]interface{}{
		"message": "test01 01test",
	}
	x01 := []byte(fmt.Sprintf("%s\n", q01["message"]))
	consoleTest(t, "t01", c01, q01, x01)

	c02 := func() *ConsoleFormatter {
		return &ConsoleFormatter{
			PrintTime: true,
			PrintKeys: true,
		}
	}
	q02 := map[string]interface{}{
		"message": "test02 02test",
	}
	x02 := []byte(fmt.Sprintf("%s\n", q02["message"]))
	consoleTest(t, "t02", c02, q02, x02)

	q03 := "test03 03test"
	q03k := "testkey03"
	q03v := "testvalue03"
	r03 := &bytes.Buffer{}
	t03 := &ConsoleFormatter{
		PrintTime: true,
		PrintKeys: true,
	}
	t03.Formatd(map[string]interface{}{
		"message": q03,
		q03k:      q03v,
	}, r03)
	v03 := q03 + " | " + q03k + ": " + q03v + "\n"
	if !bytes.Equal(r03.Bytes(), []byte(v03)) {
		t.Errorf("t03: log result and output are not equal: '%s' vs. '%s'", r03, v03)
	}

	q04 := "test04 04test"
	q04k := "testkey04"
	q04v := "testvalue04"
	q04k2 := "testkey04_2"
	q04v2 := 42
	r04 := &bytes.Buffer{}
	t04 := &ConsoleFormatter{
		PrintTime: true,
		PrintKeys: true,
	}
	t04.Formatd(map[string]interface{}{
		"message": q04,
		q04k:      q04v,
		q04k2:     q04v2,
	}, r04)
	// key order is not guaranteed - accept both versions
	v04a := fmt.Sprintf("%s | %s: %s | %s: %d\n", q04, q04k, q04v, q04k2, q04v2)
	v04b := fmt.Sprintf("%s | %s: %d | %s: %s\n", q04, q04k2, q04v2, q04k, q04v)
	if !bytes.Equal(r04.Bytes(), []byte(v04a)) && !bytes.Equal(r04.Bytes(), []byte(v04b)) {
		t.Errorf("t04: log result and output are not equal: '%s' vs. '%s' or '%s'", r04, v04a, v04b)
	}

	q05 := "test02 02test"
	q05k := "testkey05"
	q05v := "testvalue05"
	r05 := &bytes.Buffer{}
	t05 := &ConsoleFormatter{
		PrintTime: true,
	}
	t05.Formatd(map[string]interface{}{
		"message": q05,
		q05k:      q05v,
	}, r05)
	if !bytes.Equal(r05.Bytes(), []byte(q05+"\n")) {
		t.Error("t05: log result and output are not equal")
	}

	q06 := "test02 02test"
	q06t := "[2018-01-31 08:59:02] "
	q06v := time.Date(2018, 1, 31, 8, 59, 2, 0, time.Local)
	r06 := &bytes.Buffer{}
	t06 := &ConsoleFormatter{
		PrintTime: true,
	}
	t06.Formatd(map[string]interface{}{
		"message": q06,
		"time":    q06v,
	}, r06)
	if !bytes.Equal(r06.Bytes(), []byte(q06t+q06+"\n")) {
		t.Error("t06: log result and output are not equal")
	}
}
