/*
Copyright (c) 2019 Haines Chan <zhinhai@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/hainesc/pagurus/cmd/crab/app"
)

var (
	version bool
	help    bool
	flags   app.Flags
)

func main() {
	rand.Seed(time.Now().UnixNano())
	defer glog.Flush()
	// glog.MaxSize = 1024 * 1024
	flag.Parse()

	if version {
		fmt.Fprintf(os.Stdout, "crab version: %s\n", app.Version)
		return
	}
	if help {
		usage()
		return
	}
	// ----  Load config file         Complete and validate          ----
	// Flags ---------------> Options ---------------------> App ----> Run
	command := app.NewCrabCommand(&flags)
	if err := command.Execute(); err != nil {
		glog.Fatalf("TODO")
		// TODO: make sure we need this if we call Fatalf.
		os.Exit(1)
	}
}

func init() {
	flag.BoolVar(&version, "version", false, "show version information")
	flag.BoolVar(&help, "help", false, "show help message")
	flag.StringVar(&flags.ConfigFile, "config", "", "the config file used for crab")
	flag.StringVar(&flags.FunctionName, "function", "", "the function that crab served")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stdout, `crab version: %s
Usage: crab
`, app.Version)
	flag.PrintDefaults()
}
