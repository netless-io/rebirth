package utils

import (
	"os"
	"path/filepath"
)

// Exists Determine whether the file(directory) exists
func Exists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir Whether the path is a directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile Whether the path is a fail
func IsFile(path string) bool {
	return Exists(path) && !IsDir(path)
}

// ExecDir Absolute directory when rebirth is currently executed
func ExecDir() (string, bool) {
	ex, err := os.Executable()
	if err == nil {
		d := filepath.Dir(ex)
		return d, true
	}

	if exReal, err := filepath.EvalSymlinks(ex); err == nil {
		return exReal, true
	}
	return "", false
}
