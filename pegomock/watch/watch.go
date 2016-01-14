package watch

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/petergtz/pegomock/pegomock/mockgen"
)

const wellKnownInterfaceListFile = "interfaces_to_mock"

// Watch watches the specified packagePaths and continuously
// generates mocks based on the interfaces.
func Watch(gopath string, packagePaths []string, recursive bool) {
	createWellKnownInterfaceListFilesIfNecessary(gopath, packagePaths)
	for {
		for _, packagePath := range packagePaths {
			check(packagePath)
		}
		time.Sleep(2 * time.Second)
	}
}

var lastErrors = make(map[string]string)

func check(packagePath string) {
	for _, ifaceName := range interfaceNamesFrom(wellKnownInterfaceListFile) {
		defer func() {
			err := recover()
			if err != nil {
				if lastErrors[ifaceName] != fmt.Sprint(err) {
					fmt.Println("Error while trying to generate mock for", ifaceName, ":", err)
					lastErrors[ifaceName] = fmt.Sprint(err)
				}
			}
		}()
		generated, mockFilePath := mockgen.GenerateMock(packagePath, ifaceName)
		if generated || lastErrors[ifaceName] != "" {
			fmt.Println("(Re)generated mock for", ifaceName, "in", mockFilePath)
		}
		delete(lastErrors, ifaceName)
	}
}

func createWellKnownInterfaceListFilesIfNecessary(gopath string, packagePaths []string) {
	for _, packagePath := range packagePaths {
		createWellKnownInterfaceListFileIfNecessary(gopath, packagePath)
	}
}

func createWellKnownInterfaceListFileIfNecessary(gopath, packagePath string) {
	file, err := os.OpenFile(filepath.Join(gopath, "src", packagePath, wellKnownInterfaceListFile), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()
	file.WriteString("### List here all interfaces you would like to mock. One per line.\n")
}

func interfaceNamesFrom(file string) (result []string) {
	content, err := ioutil.ReadFile(file)
	panicOnError(err)
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || line == "" {
			continue
		}
		result = append(result, line)
	}
	return
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
