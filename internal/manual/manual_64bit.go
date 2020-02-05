// Copyright 2020 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build amd64 arm64 arm64be ppc64 ppc64le mips64 mips64le s390x sparc64

package manual

const (
	// maxArrayLen is a safe maximum length for slices on this architecture.
	maxArrayLen = 1<<50 - 1
)
