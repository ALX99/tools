// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.16
// +build go1.16

package main

import (
	"testing"

	"golang.org/x/tools/internal/testenv"
)

func TestGenerated(t *testing.T) {
	testenv.NeedsGoBuild(t) // This is a lie. We actually need the source code.

	// This test fails on 1.18 Kokoro for unknown reasons; in any case, it
	// suffices to run this test on any builder.
	testenv.NeedsGo1Point(t, 19)

	ok, err := doMain(false)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("documentation needs updating. run: `go run doc/generate.go` from the gopls module.")
	}
}
