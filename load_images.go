package main

import (
	"os"
	"path/filepath"
	"strings"
)

const IMAGES_PATH = "./images"

func loadImages() ([]string, error) {
	var paths []string

	err := filepath.Walk(IMAGES_PATH, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && path != IMAGES_PATH && !strings.HasSuffix(path, "keep") && !strings.HasSuffix(path, "DS_Store") {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
