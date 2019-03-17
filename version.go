// Copyright 2013 Joe Walnes and the websocketd team.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"strings"
)

// This value can be set for releases at build time using:
//   go {build|run} -ldflags "-X main.version 1.2.3 -X main.buildinfo timestamp-@githubuser-platform".
// If unset, Version() shall return "DEVBUILD".
var version = "DEVBUILD"
var buildinfo = "--"

func getVersionString() string {
	printable := strings.Replace(buildinfo, "^", " ", -1)
	return fmt.Sprintf("%s (%s %s-%s) %s", version, runtime.Version(), runtime.GOOS, runtime.GOARCH, printable)
}
