package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func SanitizePath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to fetch user home directory: %v", err)
		}
		path = filepath.Join(home, path[1:])
	}

	if runtime.GOOS == "windows" {
		path = os.ExpandEnv(path)
	}

	abs, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("failed to convert '%s' to absolute path: %v", path, err)
	}

	return abs, nil
}
