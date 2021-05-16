package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		} else {
			log.Fatalf("unable to check for %s: %v", path, err)
		}
	}

	return true
}

func PathIsDirectory(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		} else {
			log.Fatalf("unable to check for %s: %v", path, err)
		}
	}

	return info.IsDir()
}
