package main

import (
	"flag"
	"io"
	"os"

	"github.com/petergtz/pegomock/mockgen"
)

var (
	source      = flag.String("source", "", "(source mode) Input Go source file; enables source mode.")
	destination = flag.String("destination", "", "Output file; defaults to stdout.")
	packageOut  = flag.String("package", "", "Package of the generated code; defaults to the package of the input with a 'mock_' prefix.")
	selfPackage = flag.String("self_package", "", "If set, the package this mock will be part of.")

	debugParser = flag.Bool("debug_parser", false, "Print out parser results only.")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	mockgen.Run(*source, *destination, *packageOut, *selfPackage, *debugParser, flag.Args()...)
}

func usage() {
	io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `mockgen (pegomock) has two modes of operation: source and reflect.

Source mode generates mock interfaces from a source file.
It is enabled by using the -source flag. Other flags that
may be useful in this mode are -imports and -aux_files.
Example:
	mockgen -source=foo.go [other options]

Reflect mode generates mock interfaces by building a program
that uses reflection to understand interfaces. It is enabled
by passing two non-flag arguments: an import path, and a
comma-separated list of symbols.
Example:
	mockgen database/sql/driver Conn,Driver

`
