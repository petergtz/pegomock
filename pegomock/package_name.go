package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/loader"
)

func DeterminePackageNameIn(dir string) (string, error) {
	fileInfos, e := ioutil.ReadDir(dir)
	if e != nil {
		return "", fmt.Errorf("Could not get files in directory %v: %v", dir, e.Error())
	}
	var filenames []string
	var testFilenames []string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() || !strings.HasSuffix(fileInfo.Name(), ".go") {
			continue
		}
		if strings.HasSuffix(fileInfo.Name(), "_test.go") {
			testFilenames = append(testFilenames, fileInfo.Name())
		} else {
			filenames = append(filenames, fileInfo.Name())
		}
	}
	packageName, e := packageNameFrom(dir, testFilenames)
	if e != nil {
		return "", e
	}
	if packageName != "" {
		return packageName, nil
	}
	packageName, e = packageNameFrom(dir, filenames)
	if e != nil {
		return "", e
	}
	if packageName != "" {
		return packageName + "_test", nil
	}
	packageName = strings.Replace(filepath.Base(dir), "-", "_", -1)
	if packageName != "" {
		return packageName + "_test", nil
	}
	panic("Unexpected error when determining the package name of the mock file.")
}

func packageNameFrom(dir string, filenames []string) (string, error) {
	conf := loader.Config{
		Cwd:         dir,
		AllowErrors: true,
	}
	conf.CreateFromFilenames("", filenames...)
	program, e := conf.Load()
	if e != nil {
		panic(e)
	}
	if len(program.Created[0].Errors) != 0 {
		return "", fmt.Errorf("Error while determining Go package name from Go file in %v: %v", dir, program.Created[0].Errors)
	}
	return program.Created[0].Pkg.Name(), nil
}
