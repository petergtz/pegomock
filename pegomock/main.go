package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/petergtz/pegomock/pegomock/mockgen"
	"github.com/petergtz/pegomock/pegomock/watch"
)

var (
	app = kingpin.New("pegomock", "Generates mocks based on interfaces.")
)

func main() {
	Run(os.Args, os.Stderr, app)
}

func Run(cliArgs []string, out io.Writer, app *kingpin.Application) {

	workingDir, err := os.Getwd()
	app.FatalIfError(err, "")

	var (
		generateCmd = app.Command("generate", "Generate mocks based on the args provided. ")
		destination = generateCmd.Flag("output", "Output file; defaults to stdout.").Short('o').String()
		packageOut  = generateCmd.Flag("package", "Package of the generated code; defaults to the package from which pegomock was executed suffixed with _test").Default(filepath.Base(workingDir) + "_test").String()
		selfPackage = generateCmd.Flag("self_package", "If set, the package this mock will be part of.").String()
		debugParser = generateCmd.Flag("debug_parser", "Print out parser results only.").Bool()
		args        = generateCmd.Arg("args", "Interfaces or go files").Required().Strings()

		watchCmd       = app.Command("watch", "Watch ")
		watchRecursive = watchCmd.Flag("recursive", "TODO").Short('r').Hidden().Bool()
		watchPackages  = watchCmd.Arg("packages", "TODO").Strings()
	)

	app.Writer(out)
	switch kingpin.MustParse(app.Parse(cliArgs[1:])) {

	case generateCmd.FullCommand():
		validateArgs(*args)
		if *destination == "" {
			*destination = filepath.Join(workingDir, "mock_"+strings.ToLower((*args)[len(*args)-1])+"_test.go")
		}
		if sourceMode(*args) {
			mockgen.Run((*args)[0], *destination, *packageOut, *selfPackage, *debugParser)
		} else {
			if len(*args) == 1 {
				mockgen.Run("", *destination, *packageOut, *selfPackage, *debugParser, packagePathFromWorkingDirectory(os.Getenv("GOPATH"), workingDir), (*args)[0])
			} else if len(*args) == 2 {
				mockgen.Run("", *destination, *packageOut, *selfPackage, *debugParser, (*args)[0], (*args)[1])
			} else {
				app.FatalUsage("Please provide exactly 1 interface or 1 package + 1 interface")
			}
		}

	case watchCmd.FullCommand():
		if len(*watchPackages) == 0 {
			watch.Watch(gopath(), []string{packagePathFromWorkingDirectory(gopath(), workingDir)}, *watchRecursive)
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

func packagePathFromWorkingDirectory(gopath string, workingDir string) string {
	relativePackagePath, err := filepath.Rel(filepath.Join(gopath, "src"), workingDir)
	app.FatalIfError(err, "")
	return relativePackagePath
}

func gopath() string {
	if os.Getenv("GOPATH") == "" {
		app.Fatalf("No GOPATH defined. Please define GOPATH as an environment variable.")
	}
	return os.Getenv("GOPATH")

}
