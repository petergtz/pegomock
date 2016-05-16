package util

import (
	"errors"
	"os"
	"path/filepath"
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

func SourceArgs(args []string, targetPath string) ([]string, error) {
	if sourceMode(args) {
		return args[:], nil
	} else if len(args) == 1 {
		packagePath, err := packagePathFromDirectory(os.Getenv("GOPATH"), targetPath)
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

func sourceMode(args []string) bool {
	if len(args) == 1 && strings.HasSuffix(args[0], ".go") {
		return true
	}
	return false
}

func packagePathFromDirectory(gopath, dir string) (string, error) {
	relativePackagePath, err := filepath.Rel(filepath.Join(gopath, "src"), dir)
	if err != nil {
		return "", errors.New("Directory is not within a Go package path. GOPATH:" + gopath + "; dir: " + dir)
	}
	return relativePackagePath, nil
}
