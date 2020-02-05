// Copyright 2020 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build invariants tracing

package cache

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
)

func checkValue(obj interface{}) {
	v := obj.(*Value)
	if v.buf != nil {
		fmt.Fprintf(os.Stderr, "%p: cache value was not freed: refs=%d\n%s",
			v.buf, atomic.LoadInt32(&v.refs), v.traces())
		os.Exit(1)
	}
}

// newManualValue creates a Value with a manually managed buffer of size n.
//
// This definition of newManualValue is used when either the "invariants" or
// "tracing" build tags are specified. It hooks up a finalizer to the returned
// Value that checks for memory leaks when the GC determines the Value is no
// longer reachable.
func newManualValue(n int) *Value {
	if n == 0 {
		return nil
	}
	b := allocNew(n)
	v := &Value{buf: b, refs: 1}
	v.trace("alloc")
	runtime.SetFinalizer(v, checkValue)
	return v
}
