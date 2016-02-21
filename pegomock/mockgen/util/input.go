package util

import (
	"errors"
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

func SourceMode(args []string) bool {
	if len(args) == 1 && strings.HasSuffix(args[0], ".go") {
		return true
	}
	return false
}

func PackagePathFromDirectory(gopath, dir string) (string, error) {
	relativePackagePath, err := filepath.Rel(filepath.Join(gopath, "src"), dir)
	if err != nil {
		return "", errors.New("Directory is not within a Go package path. GOPATH:" + gopath + "; dir: " + dir)
	}
	return relativePackagePath, nil
}
