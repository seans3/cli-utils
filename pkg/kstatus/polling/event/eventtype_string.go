// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package event

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ResourceUpdateEvent-0]
	_ = x[CompletedEvent-1]
	_ = x[ErrorEvent-2]
}

const _EventType_name = "ResourceUpdateEventCompletedEventErrorEvent"

var _EventType_index = [...]uint8{0, 19, 33, 43}

func (i EventType) String() string {
	if i < 0 || i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
