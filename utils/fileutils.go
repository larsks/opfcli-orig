package utils

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		log.Fatalf("unable to check for %s: %v", path, err)
	}

	return true
}

func PathIsDirectory(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		log.Fatalf("unable to check for %s: %v", path, err)
	}

	return info.IsDir()
}
