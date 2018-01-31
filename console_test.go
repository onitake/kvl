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
	"bytes"
	"fmt"
	"time"
)

func TestConsoleFormatter(t *testing.T) {
	q01 := "test01 01test"
	r01 := &bytes.Buffer{}
	t01 := &ConsoleFormatter{
		Sink: r01,
	}
	t01.Logkv(map[string]interface{}{
		"message": q01,
	})
	if !bytes.Equal(r01.Bytes(), []byte(q01 + "\n")) {
		t.Error("t01: log result and output are not equal")
	}

	q02 := "test02 02test"
	r02 := &bytes.Buffer{}
	t02 := &ConsoleFormatter{
		Sink: r02,
		PrintTime: true,
		PrintKeys: true,
	}
	t02.Logkv(map[string]interface{}{
		"message": q02,
	})
	if !bytes.Equal(r02.Bytes(), []byte(q02 + "\n")) {
		t.Error("t02: log result and output are not equal")
	}

	q03 := "test03 03test"
	q03k := "testkey03"
	q03v := "testvalue03"
	r03 := &bytes.Buffer{}
	t03 := &ConsoleFormatter{
		Sink: r03,
		PrintTime: true,
		PrintKeys: true,
	}
	t03.Logkv(map[string]interface{}{
		"message": q03,
		q03k: q03v,
	})
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
		Sink: r04,
		PrintTime: true,
		PrintKeys: true,
	}
	t04.Logkv(map[string]interface{}{
		"message": q04,
		q04k: q04v,
		q04k2: q04v2,
	})
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
		Sink: r05,
		PrintTime: true,
	}
	t05.Logkv(map[string]interface{}{
		"message": q05,
		q05k: q05v,
	})
	if !bytes.Equal(r05.Bytes(), []byte(q05 + "\n")) {
		t.Error("t05: log result and output are not equal")
	}

	q06 := "test02 02test"
	q06t := "[2018-01-31 08:59:02] "
	q06v := time.Date(2018, 1, 31, 8, 59, 2, 0, time.Local)
	r06 := &bytes.Buffer{}
	t06 := &ConsoleFormatter{
		Sink: r06,
		PrintTime: true,
	}
	t06.Logkv(map[string]interface{}{
		"message": q06,
		"time": q06v,
	})
	if !bytes.Equal(r06.Bytes(), []byte(q06t + q06 + "\n")) {
		t.Error("t06: log result and output are not equal")
	}
}
