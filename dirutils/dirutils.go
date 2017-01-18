// Package util provides auxiliary functions for other MTree packages
package dirutils

import (
	"path"
	"runtime"
)

// ThisDir returns directory of the source file it's called from
func ThisDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

//CallerDir returns the directory of the caller of this function's caller
func CallerDir() string {
	_, filename, _, _ := runtime.Caller(2)
	return path.Dir(filename)
}
