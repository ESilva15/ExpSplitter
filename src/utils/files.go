package utils

import (
  "os"
  "log"
)

func LoadFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
    log.Printf("Unable to open file: %s\n", path)
	}
	return file
}
