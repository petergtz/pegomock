package util

import "os"

// WithinWorkingDir changes the current working directory temporarily and
// executes cb within that context.
func WithinWorkingDir(targetPath string, cb func(workingDir string)) {
	origWorkingDir, e := os.Getwd()
	PanicOnError(e)
	e = os.Chdir(targetPath)
	PanicOnError(e)
	defer func() { os.Chdir(origWorkingDir) }()
	cb(targetPath)
}
