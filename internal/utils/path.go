package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func SanitizePath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}

	if runtime.GOOS == "windows" {
		path = os.ExpandEnv(path)
	}

	cleaned := filepath.Clean(path)

	abs, err := filepath.Abs(cleaned)
	if err != nil {
		return "", err
	}

	return abs, nil
}
