// Copyright 2018 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nflog implements a garbage-collected and snapshottable append-only log of
// active/resolved notifications. Each log entry stores the active/resolved state,
// the notified receiver, and a hash digest of the notification's identifying contents.
// The log can be queried along different parameters.
package hash

func Uint64(uints []uint64) uint64 {
	if len(uints) == 1 {
		return uints[0]
	}
	if len(uints) == 0 {
		return 0
	}
	var hash uint64
	for _, n := range uints {
		hash ^= n
	}

	return hash
}
