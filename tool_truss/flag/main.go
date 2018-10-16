package main

import (
	"flag"
	"fmt"
	"os"
	"test/tool_truss/linux"
	"test/tool_truss/windows"
)

var (
	help        bool
	version     bool
	path        string
	environment string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&path, "p", "/", "set path value")
	flag.StringVar(&environment, "e", "w", "set environment:\n\tw means windows\n\tl means linux")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage = usage
		flag.Usage()
	} else if version {
		flag.Usage = trussVersion
		flag.Usage()
	} else {
		flag.Usage = trussRun
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
truss_tools version: truss_tools/1.0
Usage: truss_tools [-h help] [-v version] [-p "w" or "l" environment] [-p truss_filePath]
Options:
`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
eg1:truss_tools -e w -p "C:\it_coderepo\go\src\github.com\tuneinc\truss"
eg1:truss_tools -e l -p "/root/go/src/github.com/tuneinc/truss"
`)
}

func trussVersion() {
	fmt.Fprintln(os.Stderr, "version: truss_tools/1.0")
}

func trussPath() {
	fmt.Fprintf(os.Stderr, "path is %s\n", path)
}

func trussRun() {
	trussEnvironment()
}

func trussEnvironment() {
	if environment == "w" {
		fmt.Fprintln(os.Stderr, "windows environment")
		trussPath()
		windows.Option(path)
		windows.Windows(path)
	} else if environment == "l" {
		fmt.Fprintln(os.Stderr, "linux environment")
		trussPath()
		linux.Option(path)
		linux.Linux(path)
		//..
	} else {
		fmt.Fprintln(os.Stderr, "error environment")
	}
}
