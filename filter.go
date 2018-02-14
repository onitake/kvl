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
	"time"
)

const (
	// filterListAllocation is the default length of the filter list, and
	// its growth.
	filterListAllocation = 10
)

// AddTimeFilter adds StdTimeKey, containing the current local time as
// a time.Time object.
// If TimeFormat is set, the current time will be formatted according to
// this format.
// If StdTimeKey is already present and a string, it will not be modified.
// If it is present but a time.Time object and TimeFormat is set, it
// will be formatted according to this format.
// If it is present but of a different type, it will be treated like it
// wasn't present.
type AddTimeFilter struct {
	TimeFormat string
}

func (filter *AddTimeFilter) Printd(kv map[string]interface{}) {
	switch t := kv[StdTimeKey].(type) {
	case string:
		// pass
	case time.Time:
		if filter.TimeFormat != "" {
			kv[StdTimeKey] = t.Format(filter.TimeFormat)
		}
	default:
		if filter.TimeFormat != "" {
			kv[StdTimeKey] = time.Now().Format(filter.TimeFormat)
		} else {
			kv[StdTimeKey] = time.Now()
		}
	}
}

// MergeFilter merges a constant dictionary with anything that is being logged.
// Existing values are not replaced.
//
// Note: Go does not actually support constant dictionaries or generic
// value types. Do NOT do deep modifications of values.
type MergeFilter struct {
	Dict map[string]interface{}
}

func (filter *MergeFilter) Printd(kv map[string]interface{}) {
	for k, v := range filter.Dict {
		if _, ok := kv[k]; !ok {
			kv[k] = v
		}
	}
}

// MultiFilter applies a list of filters in sequence, then sends the output
// to a Logger (or another filter).
type MultiFilter struct {
	Filters []Filter
	Logger  Filter
}

func (filter *MultiFilter) Printd(kv map[string]interface{}) {
	for _, f := range filter.Filters {
		f.Printd(kv)
	}
}

// AddFilter appends a filter to the end of the list.
func (filter *MultiFilter) AddFilter(f Filter) {
	// grow if necessary
	if len(filter.Filters)+1 >= cap(filter.Filters) {
		filters := make([]Filter, len(filter.Filters), len(filter.Filters)+filterListAllocation)
		copy(filters, filter.Filters)
		filter.Filters = filters
	}
	filter.Filters = append(filter.Filters, f)
}

// ClearFilters clears the filter list.
func (filter *MultiFilter) ClearFilters() {
	// allocate with default size
	filter.Filters = make([]Filter, 0, filterListAllocation)
}
