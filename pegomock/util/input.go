package util

import (
	"errors"
	"fmt"
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
	if len(args) == 1 {
		packagePath, err := packagePathFromWorkingDirectoryAndGoModule()
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

func packagePathFromWorkingDirectoryAndGoModule() (string, error) {
	dir, e := os.Getwd()
	if e != nil {
		return "", e
	}
	gomodDir := findModuleRoot(dir)
	subPackage, e := filepath.Rel(gomodDir, dir)
	if e != nil {
		return "", errors.New("Could not get a relative path for " + dir + " based on path " + gomodDir)
	}
	gomodFilepath := filepath.Join(gomodDir, "go.mod")

	content, e := os.ReadFile(gomodFilepath)
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

// copied from https://github.com/golang/go/blob/ab724d43efe7e1a7516c1d13e40b55dca26a61b4/src/cmd/go/internal/modload/init.go#L480-L495:
func findModuleRoot(dir string) (root string) {
	dir = filepath.Clean(dir)

	// Look for enclosing go.mod.
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
	}
	return ""
}
