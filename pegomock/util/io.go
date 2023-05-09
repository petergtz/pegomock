package util

import "os"

// WithinWorkingDir changes the current working directory temporarily and
// executes cb within that context.
func WithinWorkingDir(targetPath string, cb func(workingDir string)) {
	origWorkingDir, e := os.Getwd()
	PanicOnError(e)
	e = os.Chdir(targetPath)
	PanicOnError(e)
	defer func() { _ = os.Chdir(origWorkingDir) }()
	cb(targetPath)
}

func WriteFileIfChanged(outputFilepath string, output []byte) bool {
	existingFileContent, e := os.ReadFile(outputFilepath)
	if os.IsNotExist(e) {
		e = os.WriteFile(outputFilepath, output, 0664)
		PanicOnError(e)
		return true
	}
	PanicOnError(e)

	if string(existingFileContent) == string(output) {
		return false
	}
	e = os.WriteFile(outputFilepath, output, 0664)
	PanicOnError(e)
	return true
}
