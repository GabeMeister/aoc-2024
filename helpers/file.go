package helpers

import (
	"os"
	"strings"
)

func ReadFile(path string) []string {
	bytes, _ := os.ReadFile(path)

	return strings.Split(string(bytes), "\n")
}
