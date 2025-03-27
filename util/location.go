package util

import (
	"os"
	"path/filepath"
	"strings"
)

func Location() string {
	var dir string
	executable, _ := os.Executable()
	if strings.Contains(executable, os.TempDir()) {
		dir, _ = os.Getwd()
	} else {
		dir = filepath.Dir(executable)
	}
	return dir
}
