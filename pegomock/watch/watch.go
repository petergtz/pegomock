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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/petergtz/pegomock/pegomock/mockgen"
	"github.com/petergtz/pegomock/pegomock/mockgen/util"
)

const wellKnownInterfaceListFile = "interfaces_to_mock"

// Watch watches the specified packagePaths and continuously
// generates mocks based on the interfaces.
func Watch(targetPaths []string, recursive bool, done chan bool) {
	createWellKnownInterfaceListFilesIfNecessary(targetPaths)
	for {
		select {
		case <-done:
			return

		default:
			for _, targetPath := range targetPaths {
				check(targetPath)
			}
			time.Sleep(2 * time.Second)
		}
	}
}

var lastErrors = make(map[string]string)

func check(targetPath string) {
	// TODO: currently this returns also all CLI options. In the end args should be parsed by kingpin again to properly use it.
	for _, args := range linesIn(wellKnownInterfaceListFile) {
		defer func() {
			err := recover()
			if err != nil {
				// TODO: this mechanism doesnt work correctly now, because all possible CLI options would go into the joined string as well
				if lastErrors[strings.Join(args, "_")] != fmt.Sprint(err) {
					fmt.Println("Error while trying to generate mock for", strings.Join(args, "_"), ":", err)
					lastErrors[strings.Join(args, "_")] = fmt.Sprint(err)
				}
			}
		}()
		panicOnError(util.ValidateArgs(args))
		sourceArgs, err := sourceArgs(args, targetPath)
		panicOnError(err)

		generatedMockSource := mockgen.GenerateMockSourceCode(sourceArgs, filepath.Base(targetPath)+"_test", "", false, os.Stdout)
		mockFilePath := mockgen.OutputFilePath(sourceArgs, targetPath, "") // <- adjust last param
		hasChanged := writeFileIfChanged(mockFilePath, generatedMockSource)

		if hasChanged || lastErrors[strings.Join(args, "_")] != "" {
			fmt.Println("(Re)generated mock for", strings.Join(args, "_"), "in", mockFilePath)
		}
		delete(lastErrors, strings.Join(args, "_"))
	}
}

func sourceArgs(args []string, targetPath string) ([]string, error) {
	if util.SourceMode(args) {
		return args[:], nil
	} else if len(args) == 1 {
		packagePath, err := util.PackagePathFromDirectory(os.Getenv("GOPATH"), targetPath)
		if err != nil {
			return nil, errors.New("Couldn't determine package path from directory")
		}
		return []string{packagePath, args[0]}, nil
	} else if len(args) == 2 {
		return args[:], nil
	} else {
		return nil, errors.New("Please provide exactly 1 interface or 1 package + 1 interface in the interfaces_to_mock file")
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

func createWellKnownInterfaceListFilesIfNecessary(targetPaths []string) {
	for _, targetPath := range targetPaths {
		createWellKnownInterfaceListFileIfNecessary(targetPath)
	}
}

func createWellKnownInterfaceListFileIfNecessary(targetPath string) {
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
