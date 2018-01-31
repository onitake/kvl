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

// Filter is the standard interface for a data processor.
// This interface is also used in Formatters.
//
// Note: The log maps sent through this interface may be modified by
// anything in the chain. Clone them if you need to keep them.
// Data will not be modified, however, so a deep copy is not necessary.
type Filter interface {
	// Logkv accepts key-value maps, processes them and sends them to the next
	// Filters or Formatters in the chain.
	// Or, in the case of a Formatter, to a log Sink.
	// This function may modify the map, but not its values.
	Logkv(kv map[string]interface{})
}

// AddTimeFilter adds the key "time" with the current time
// and sends the dictionary to the next filter in the chain.
// If a "time" is already present, it will not be modified.
type AddTimeFilter struct {
	Chain Filter
}

func (filter *AddTimeFilter) Logkv(kv map[string]interface{}) {
	if filter.Chain != nil {
		// add the curent time
		if _, ok := kv[StdTimeKey]; !ok {
			kv[StdMessageKey] = time.Now()
		}
		// pass the map on
		filter.Chain.Logkv(kv)
	}
}
