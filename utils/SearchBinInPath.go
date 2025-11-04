package utils

import (
	"os"
	"strings"
)

func SearchBinInPath(bin string) (string, bool) {
	pathEnv := os.Getenv("PATH")

	dirs := strings.Split(pathEnv, ":")
	for _, dir := range dirs {
		fullPath := dir + "/" + bin

		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}
