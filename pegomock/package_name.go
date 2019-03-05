package main

import (
	"path/filepath"
	"strings"
)

func DeterminePackageNameIn(dir string) (string, error) {
	return strings.Replace(filepath.Base(dir), "-", "_", -1) + "_test", nil
}
