package main

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/petergtz/pegomock/pegomock/mockgen"
	"github.com/petergtz/pegomock/pegomock/watch"
)

var (
	app = kingpin.New("pegomock", "Generates mocks based on interfaces.")

	generateCmd = app.Command("generate", "Generate mocks based on the args provided. ")
	destination = generateCmd.Flag("output", "Output file; defaults to stdout.").Short('o').String()
	packageOut  = generateCmd.Flag("package", "Package of the generated code; defaults to the package of the input with a 'mock_' prefix.").String()
	selfPackage = generateCmd.Flag("self_package", "If set, the package this mock will be part of.").String()
	debugParser = generateCmd.Flag("debug_parser", "Print out parser results only.").Bool()
	args        = generateCmd.Arg("args", "Interfaces or go files").Required().Strings()

	watchCmd       = app.Command("watch", "Watch ")
	watchRecursive = watchCmd.Flag("recursive", "TODO").Short('r').Hidden().Bool()
	watchPackages  = watchCmd.Arg("packages", "TODO").Strings()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case generateCmd.FullCommand():
		validateArgs(*args)
		if sourceMode(*args) {
			mockgen.Run((*args)[0], *destination, *packageOut, *selfPackage, *debugParser)
		} else {
			if len(*args) == 1 {
				mockgen.Run("", *destination, *packageOut, *selfPackage, *debugParser, packagePathFromWorkingDirectory(gopath()), (*args)[0])
			} else {
				mockgen.Run("", *destination, *packageOut, *selfPackage, *debugParser, *args...)
			}
		}

	case watchCmd.FullCommand():
		if len(*watchPackages) == 0 {
			watch.Watch(gopath(), []string{packagePathFromWorkingDirectory(gopath())}, *watchRecursive)
		} else {
			watch.Watch(gopath(), *watchPackages, *watchRecursive)
		}
	}
}

func validateArgs(args []string) {
	if len(args) == 0 {
		app.FatalUsage("You must specify either exactly one source filename ending with .go, or at least one go interface name.")
	}
	if len(args) == 1 {
		return
	}
	if len(args) >= 2 {
		for _, arg := range args {
			if strings.HasSuffix(arg, ".go") {
				app.FatalUsage("You can specify at most one go source file.")
			}
		}
	}
}

func sourceMode(args []string) bool {
	if len(args) == 1 && strings.HasSuffix(args[0], ".go") {
		return true
	}
	return false
}

// PackagePathFromWorkingDirectory TODO
func packagePathFromWorkingDirectory(gopath string) string {
	absolutePackagePath, err := os.Getwd()
	app.FatalIfError(err, "")
	relativePackagePath, err := filepath.Rel(filepath.Join(gopath, "src"), absolutePackagePath)
	app.FatalIfError(err, "")
	return relativePackagePath
}

func gopath() string {
	if os.Getenv("GOPATH") == "" {
		app.Fatalf("No GOPATH defined. Please define GOPATH as an environment variable.")
	}
	return os.Getenv("GOPATH")

}
