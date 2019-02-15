package tools

import (
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		LogErr.Printf("Get Root Path Error: %v", err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}
