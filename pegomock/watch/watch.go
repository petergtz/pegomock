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

package watch

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/petergtz/pegomock/pegomock/mockgen"
	"github.com/petergtz/pegomock/pegomock/mockgen/util"
)

const wellKnownInterfaceListFile = "interfaces_to_mock"

type Callable interface {
	Call()
}

type Checker struct {
	recursive   bool
	targetPaths []string
}

func NewChecker(targetPaths []string, recursive bool) *Checker {
	return &Checker{targetPaths: targetPaths, recursive: recursive}
}

func (checker *Checker) Call() {
	for _, targetPath := range checker.targetPaths {
		check(targetPath)
	}
}

// Watch watches the specified packagePaths and continuously
// generates mocks based on the interfaces.
func Watch(callable Callable, done chan bool) {
	for {
		select {
		case <-done:
			return

		default:
			callable.Call()
			time.Sleep(2 * time.Second)
		}
	}
}

var lastErrors = make(map[string]string)

func check(targetPath string) {
	origWorkingDir, e := os.Getwd()
	panicOnError(e)
	e = os.Chdir(targetPath)
	panicOnError(e)
	defer func() { os.Chdir(origWorkingDir) }()

	if _, err := os.Stat(wellKnownInterfaceListFile); os.IsNotExist(err) {
		return
	}

	for _, lineParts := range linesIn(wellKnownInterfaceListFile) {
		lineCmd := kingpin.New("What should go in here", "And what should go in here")
		destination := lineCmd.Flag("output", "Output file; defaults to mock_<interface>_test.go.").Short('o').String()
		packageOut := lineCmd.Flag("package", "Package of the generated code; defaults to the package from which pegomock was executed suffixed with _test").Default(filepath.Base(targetPath) + "_test").String()
		selfPackage := lineCmd.Flag("self_package", "If set, the package this mock will be part of.").String()
		lineArgs := lineCmd.Arg("args", "A (optional) Go package path + space-separated interface or a .go file").Required().Strings()

		_, parseErr := lineCmd.Parse(lineParts)
		if parseErr != nil {
			fmt.Println("Error while trying to generate mock for line", strings.Join(lineParts, " "), ":", parseErr)
			continue
		}
		defer func() {
			err := recover()
			if err != nil {
				if lastErrors[strings.Join(*lineArgs, "_")] != fmt.Sprint(err) {
					fmt.Println("Error while trying to generate mock for", strings.Join(lineParts, " "), ":", err)
					lastErrors[strings.Join(*lineArgs, "_")] = fmt.Sprint(err)
				}
			}
		}()

		panicOnError(util.ValidateArgs(*lineArgs))
		sourceArgs, err := util.SourceArgs(*lineArgs)
		panicOnError(err)

		generatedMockSourceCode := mockgen.GenerateMockSourceCode(sourceArgs, *packageOut, *selfPackage, false, os.Stdout)
		mockFilePath := mockgen.OutputFilePath(sourceArgs, ".", *destination)
		hasChanged := writeFileIfChanged(mockFilePath, generatedMockSourceCode)

		if hasChanged || lastErrors[strings.Join(*lineArgs, "_")] != "" {
			fmt.Println("(Re)generated mock for", strings.Join(*lineArgs, "_"), "in", mockFilePath)
		}
		delete(lastErrors, strings.Join(*lineArgs, "_"))
	}
}

func writeFileIfChanged(outputFilepath string, output []byte) bool {
	existingFileContent, err := ioutil.ReadFile(outputFilepath)
	if err != nil {
		if os.IsNotExist(err) {
			err = ioutil.WriteFile(outputFilepath, output, 0664)
			panicOnError(err)
			return true
		} else {
			panic(err)
		}
	}
	if string(existingFileContent) == string(output) {
		return false
	} else {
		err = ioutil.WriteFile(outputFilepath, output, 0664)
		panicOnError(err)
		return true
	}
}

func CreateWellKnownInterfaceListFilesIfNecessary(targetPaths []string) {
	for _, targetPath := range targetPaths {
		CreateWellKnownInterfaceListFileIfNecessary(targetPath)
	}
}

func CreateWellKnownInterfaceListFileIfNecessary(targetPath string) {
	file, err := os.OpenFile(filepath.Join(targetPath, wellKnownInterfaceListFile), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()
	file.WriteString("### List here all interfaces you would like to mock. One per line.\n")
}

func linesIn(file string) (result [][]string) {
	content, err := ioutil.ReadFile(file)
	panicOnError(err)
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || line == "" {
			continue
		}
		parts := regexp.MustCompile(`\s`).Split(line, -1)
		// TODO: do validation here like in main
		result = append(result, parts)
	}
	return
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
