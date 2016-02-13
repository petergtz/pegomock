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

	"github.com/petergtz/pegomock/pegomock/mockgen"
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

var lastErrors = make(map[struct{ packageName, interfaceName string }]string)

func check(targetPath string) {
	for _, packageAndIfaceName := range packageAndInterfaceNamesFrom(wellKnownInterfaceListFile) {
		defer func() {
			err := recover()
			if err != nil {
				if lastErrors[packageAndIfaceName] != fmt.Sprint(err) {
					fmt.Println("Error while trying to generate mock for", packageAndIfaceName, ":", err)
					lastErrors[packageAndIfaceName] = fmt.Sprint(err)
				}
			}
		}()
		generated, mockFilePath := mockgen.GenerateMock(
			packageAndIfaceName.packageName,
			packageAndIfaceName.interfaceName,
			targetPath,
			filepath.Base(targetPath)+"_test")
		if generated || lastErrors[packageAndIfaceName] != "" {
			fmt.Println("(Re)generated mock for", packageAndIfaceName, "in", mockFilePath)
		}
		delete(lastErrors, packageAndIfaceName)
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

func packageAndInterfaceNamesFrom(file string) (result []struct{ packageName, interfaceName string }) {
	content, err := ioutil.ReadFile(file)
	panicOnError(err)
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || line == "" {
			continue
		}
		parts := regexp.MustCompile(`\W`).Split(line, -1)
		if len(parts) == 2 {
			result = append(result, struct{ packageName, interfaceName string }{parts[0], parts[1]})
		} else {
			result = append(result, struct{ packageName, interfaceName string }{packagePathFromWorkingDirectory(), parts[0]})
		}
	}
	return
}

func packagePathFromWorkingDirectory() string {
	workingDir, err := os.Getwd()
	panicOnError(err)

	relativePackagePath, err := filepath.Rel(filepath.Join(gopath(), "src"), workingDir)
	panicOnError(err)
	return relativePackagePath
}

func gopath() string {
	if os.Getenv("GOPATH") == "" {
		panic("No GOPATH defined. Please define GOPATH as an environment variable.")
	}
	return os.Getenv("GOPATH")

}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
