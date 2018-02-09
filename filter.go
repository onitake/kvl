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

func (filter *AddTimeFilter) Filterd(kv map[string]interface{}) {
	switch t := kv[StdTimeKey].(type) {
	case string:
		// pass
	case time.Time:
		if filter.TimeFormat != "" {
			kv[StdMessageKey] = t.Format(filter.TimeFormat)
		}
	default:
		if filter.TimeFormat != "" {
			kv[StdMessageKey] = time.Now().Format(filter.TimeFormat)
		} else {
			kv[StdMessageKey] = time.Now()
		}
	}
}
