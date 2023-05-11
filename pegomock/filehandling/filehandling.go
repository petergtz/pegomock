package filehandling

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/petergtz/pegomock/v4/mockgen"
	"github.com/petergtz/pegomock/v4/modelgen/xtools_packages"
)

func GenerateMockFileInOutputDir(
	args []string,
	outputDirPath string,
	outputFilePathOverride string,
	nameOut string,
	packageOut string,
	selfPackage string,
	debugParser bool,
	out io.Writer) {

	// if a file path override is specified
	// ensure all directories in the path are created
	if outputFilePathOverride != "" {
		if err := os.MkdirAll(filepath.Dir(outputFilePathOverride), 0755); err != nil {
			panic(fmt.Errorf("failed to make output directory, error: %v", err))
		}
	}

	GenerateMockFile(
		args,
		OutputFilePath(args, outputDirPath, outputFilePathOverride),
		nameOut,
		packageOut,
		selfPackage,
		debugParser,
		out)
}

func OutputFilePath(args []string, outputDirPath string, outputFilePathOverride string) string {
	if outputFilePathOverride != "" {
		return outputFilePathOverride
	} else {
		return filepath.Join(outputDirPath, "mock_"+strings.ToLower(args[len(args)-1])+"_test.go")
	}
}

func GenerateMockFile(args []string, outputFilePath string, nameOut string, packageOut string, selfPackage string, debugParser bool, out io.Writer) {
	mockSourceCode := GenerateMockSourceCode(args, nameOut, packageOut, selfPackage, debugParser, out)

	err := os.WriteFile(outputFilePath, mockSourceCode, 0664)
	if err != nil {
		panic(fmt.Errorf("failed writing to destination: %v", err))
	}
}

func GenerateMockSourceCode(args []string, nameOut string, packageOut string, selfPackage string, debugParser bool, out io.Writer) []byte {
	if len(args) != 2 {
		log.Fatal("Expected exactly two arguments, but got " + fmt.Sprint(args))
	}
	ast, err := xtools_packages.GenerateModel(args[0], args[1])
	src := fmt.Sprintf("%v (interfaces: %v)", args[0], args[1])
	if err != nil {
		panic(fmt.Errorf("loading input failed: %v", err))
	}

	if debugParser {
		ast.Print(out)
	}
	return mockgen.GenerateOutput(ast, src, nameOut, packageOut, selfPackage)
}
