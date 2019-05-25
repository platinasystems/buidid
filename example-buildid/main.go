// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/platinasystems/buildid"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"/proc/self/exe"}
	}
	for _, fn := range args {
		s, err := buildid.New(fn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Print(fn, ": ")
		}
		fmt.Println(s)
	}
}
