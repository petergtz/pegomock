// Copyright 2016 Peter Goetz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/petergtz/pegomock/pegomock/mockgen"
	"github.com/petergtz/pegomock/pegomock/mockgen/util"
	"github.com/petergtz/pegomock/pegomock/watch"
)

var (
	app = kingpin.New("pegomock", "Generates mocks based on interfaces.")
)

func main() {
	Run(os.Args, os.Stderr, app, make(chan bool))
}

func Run(cliArgs []string, out io.Writer, app *kingpin.Application, done chan bool) {

	workingDir, err := os.Getwd()
	app.FatalIfError(err, "")

	var (
		generateCmd     = app.Command("generate", "Generate mocks based on the args provided. ")
		destination     = generateCmd.Flag("output", "Output file; defaults to mock_<interface>_test.go.").Short('o').String()
		packageOut      = generateCmd.Flag("package", "Package of the generated code; defaults to the package from which pegomock was executed suffixed with _test").Default(filepath.Base(workingDir) + "_test").String()
		selfPackage     = generateCmd.Flag("self_package", "If set, the package this mock will be part of.").String()
		debugParser     = generateCmd.Flag("debug", "Print debug information.").Short('d').Bool()
		generateCmdArgs = generateCmd.Arg("args", "A (optional) Go package path + space-separated interface or a .go file").Required().Strings()

		watchCmd       = app.Command("watch", "Watch ")
		watchRecursive = watchCmd.Flag("recursive", "TODO").Short('r').Hidden().Bool()
		watchPackages  = watchCmd.Arg("packages", "TODO").Strings()
	)

	app.Writer(out)
	switch kingpin.MustParse(app.Parse(cliArgs[1:])) {

	case generateCmd.FullCommand():
		if err := util.ValidateArgs(*generateCmdArgs); err != nil {
			app.FatalUsage(err.Error())
		}
		sourceArgs, err := util.SourceArgs(*generateCmdArgs, workingDir)
		if err != nil {
			app.FatalUsage(err.Error())
		}

		mockgen.GenerateMockFileInOutputDir(
			sourceArgs,
			workingDir,
			*destination,
			*packageOut,
			*selfPackage,
			*debugParser,
			out)

	case watchCmd.FullCommand():
		if len(*watchPackages) == 0 {
			watch.Watch([]string{workingDir}, *watchRecursive, done)
		} else {
			watch.Watch(*watchPackages, *watchRecursive, done)
		}
	}
}
