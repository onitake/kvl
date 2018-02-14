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
	"time"
)

func TestAddTimeFilter(t *testing.T) {
	if StdTimeKey != "time" {
		t.Error("t00: time key should be named 'time'")
	}

	c01 := map[string]interface{}{}
	f01 := &AddTimeFilter{}
	f01.Printd(c01)
	if _, ok := c01[StdTimeKey].(time.Time); !ok {
		t.Errorf("t01: invalid type for time key")
	}

	c02 := map[string]interface{}{}
	f02 := &AddTimeFilter{
		TimeFormat: time.RFC3339,
	}
	f02.Printd(c02)
	if _, ok := c02[StdTimeKey].(string); !ok {
		t.Error("t02: invalid type for time key")
	}

	c03 := map[string]interface{}{
		"message": "test03",
	}
	f03 := &AddTimeFilter{}
	f03.Printd(c03)
	if _, ok := c03[StdTimeKey].(time.Time); !ok || c03["message"] != "test03" {
		t.Error("t03: invalid type for time key")
	}

	t04 := time.Date(2018, 1, 31, 8, 59, 2, 0, time.UTC)
	c04 := map[string]interface{}{
		StdTimeKey: t04,
	}
	f04 := &AddTimeFilter{}
	// wait for a time change
	time.Sleep(1)
	f04.Printd(c04)
	r04, ok := c04[StdTimeKey].(time.Time)
	if !ok {
		t.Error("t04: invalid type for time key")
	}
	if r04 != t04 {
		t.Error("t04: time does not match")
	}

	t05 := time.Date(2018, 1, 31, 8, 59, 2, 0, time.UTC)
	c05 := map[string]interface{}{
		StdTimeKey: t05,
	}
	f05 := &AddTimeFilter{
		TimeFormat: time.RFC3339,
	}
	// wait for a time change
	time.Sleep(1)
	f05.Printd(c05)
	r05, ok := c05[StdTimeKey].(string)
	if !ok {
		t.Error("t05: invalid type for time key")
	}
	if r05 != t05.Format(time.RFC3339) {
		t.Error("t05: time does not match")
	}
}

func TestMergeFilter(t *testing.T) {
	c01 := map[string]interface{}{}
	f01 := &MergeFilter{}
	f01.Printd(c01)
	if len(c01) != 0 {
		t.Errorf("t01: dict didn't stay empty")
	}

	c02 := map[string]interface{}{}
	f02 := &MergeFilter{
		Dict: map[string]interface{}{
			"key02a": "value02a",
		},
	}
	f02.Printd(c02)
	if len(c02) != 1 || c02["key02a"] != "value02a" {
		t.Errorf("t02: key not found")
	}

	c03 := map[string]interface{}{
		"key03a": "value03a",
	}
	f03 := &MergeFilter{}
	f03.Printd(c03)
	if len(c03) != 1 || c03["key03a"] != "value03a" {
		t.Errorf("t03: dict changed unexpectedly")
	}

	c04 := map[string]interface{}{
		"key04a": "value04a",
	}
	f04 := &MergeFilter{
		Dict: map[string]interface{}{
			"key04b": "value04b",
		},
	}
	f04.Printd(c04)
	if len(c04) != 2 || c04["key04a"] != "value04a" || c04["key04b"] != "value04b" {
		t.Errorf("t04: key not found")
	}

	c05 := map[string]interface{}{
		"key05a": "value05a",
		"key05b": "value05b",
	}
	f05 := &MergeFilter{}
	f05.Printd(c05)
	if len(c05) != 2 || c05["key05a"] != "value05a" || c05["key05b"] != "value05b" {
		t.Errorf("t05: dict changed unexpectedly")
	}

	c06 := map[string]interface{}{
		"key06a": "value06a",
		"key06b": "value06b",
	}
	f06 := &MergeFilter{
		Dict: map[string]interface{}{
			"key06c": "value06c",
		},
	}
	f06.Printd(c06)
	if len(c06) != 3 || c06["key06a"] != "value06a" || c06["key06b"] != "value06b" || c06["key06c"] != "value06c" {
		t.Errorf("t06: dict changed unexpectedly")
	}

	c07 := map[string]interface{}{
		"key07a": "value07a",
		"key07b": "value07b",
	}
	f07 := &MergeFilter{
		Dict: map[string]interface{}{
			"key07c": "value07c",
			"key07d": "value07d",
		},
	}
	f07.Printd(c07)
	if len(c07) != 4 || c07["key07a"] != "value07a" || c07["key07b"] != "value07b" || c07["key07c"] != "value07c" || c07["key07d"] != "value07d" {
		t.Errorf("t06: dict changed unexpectedly")
	}

	c08 := map[string]interface{}{
		"key08a": "value08a",
	}
	f08 := &MergeFilter{
		Dict: map[string]interface{}{
			"key08a": "value08b",
		},
	}
	f08.Printd(c08)
	if len(c08) != 1 || c08["key08a"] != "value08a" {
		t.Errorf("t08: value was replaced")
	}
}

func arrayContains(array []Filter, value Filter) bool {
	contains := false
	for _, v := range array {
		if v == value {
			contains = true
		}
	}
	return contains
}

func TestMultiFilter(t *testing.T) {
	c01 := map[string]interface{}{
		"key01a": "value01a",
	}
	f01 := &MultiFilter{
		Filters: []Filter{
			&MergeFilter{
				Dict: map[string]interface{}{
					"key01b": "value01b",
				},
			},
			&MergeFilter{
				Dict: map[string]interface{}{
					"key01c": "value01c",
					"key01d": "value01d",
				},
			},
		},
	}
	f01.Printd(c01)
	if len(c01) != 4 {
		t.Errorf("t01: dict wasn't combined correctly")
	}

	f02 := &MultiFilter{}
	f02a := &MergeFilter{}
	f02.AddFilter(f02a)
	f02b := &MergeFilter{}
	f02.AddFilter(f02b)
	if len(f02.Filters) != 2 || !arrayContains(f02.Filters, f02a) || !arrayContains(f02.Filters, f02b) {
		t.Errorf("t02: filters weren't added")
	}

	f03 := &MultiFilter{}
	f03a := &MergeFilter{}
	f03.AddFilter(f03a)
	f03b := &MergeFilter{}
	f03.AddFilter(f03b)
	f03.ClearFilters()
	f03c := &MergeFilter{}
	f03.AddFilter(f03c)
	if len(f03.Filters) != 1 || !arrayContains(f03.Filters, f03c) {
		t.Errorf("t03: list wasn't cleared before adding")
	}
}
