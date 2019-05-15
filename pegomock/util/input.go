package util

import (
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ValidateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("You must specify either exactly one source filename ending with .go, or at least one go interface name.")
	}
	if len(args) == 1 {
		return nil
	}
	if len(args) >= 2 {
		for _, arg := range args {
			if strings.HasSuffix(arg, ".go") {
				return errors.New("You can specify at most one go source file.")
			}
		}
	}
	return nil
}

func SourceArgs(args []string) ([]string, error) {
	if SourceMode(args) {
		return args[:], nil
	} else if len(args) == 1 {
		packagePath, err := packagePathFromWorkingDirectoryAndGoPathOrGoModule()
		if err != nil {
			return nil, fmt.Errorf("Couldn't determine package path from directory: %v", err)
		}
		return []string{packagePath, args[0]}, nil
	} else if len(args) == 2 {
		return args[:], nil
	} else {
		return nil, errors.New("Please provide exactly 1 interface or 1 package + 1 interface in the interfaces_to_mock file")
	}
}

func SourceMode(args []string) bool {
	if len(args) == 1 && strings.HasSuffix(args[0], ".go") {
		return true
	}
	return false
}

func packagePathFromWorkingDirectoryAndGoPathOrGoModule() (string, error) {
	dir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	if os.Getenv("GO111MODULE") == "off" ||
		((os.Getenv("GO111MODULE") == "auto" || os.Getenv("GO111MODULE") == "") &&
			os.Getenv("GOPATH") != "") {
		return packagePathFromDirectoryAndGoPath(dir, build.Default.GOPATH)
	}
	return packagePathFromDirectoryAndGoMod(dir)
}

func packagePathFromDirectoryAndGoPath(dir string, gopath string) (string, error) {
	relativePackagePath, err := filepath.Rel(filepath.Join(gopath, "src"), dir)
	if err != nil || strings.HasPrefix(relativePackagePath, "..") {
		return "", errors.New("Directory is not within a Go package path. GOPATH:" + gopath + "; dir: " + dir)
	}
	return relativePackagePath, nil
}

func packagePathFromDirectoryAndGoMod(dir string) (string, error) {
	gomodDir := dir
	subPackage := ""
	for {
		if _, err := os.Stat(filepath.Join(gomodDir, "go.mod")); !os.IsNotExist(err) {
			break
		}
		if gomodDir == "/" {
			return "", errors.New("Not within a file tree that contains go.mod file. Current directory: " + dir)
		}
		subPackage = filepath.Join(filepath.Base(gomodDir), subPackage)
		gomodDir = filepath.Dir(gomodDir)
	}
	gomodFilepath := filepath.Join(gomodDir, "go.mod")
	content, e := ioutil.ReadFile(gomodFilepath)
	if e != nil {
		return "", errors.New("Could not read file " + gomodFilepath)
	}
	modulePathMatches := regexp.MustCompile(`^module (.*)\n`).FindSubmatch(content)
	if len(modulePathMatches) != 2 {
		return "", errors.New("Cannot parse" + gomodFilepath + "file. File does not start with 'module'")
	}
	return filepath.Join(
			string(modulePathMatches[1]),
			subPackage),
		nil
}
