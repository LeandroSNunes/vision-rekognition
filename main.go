package main

import (
	"fmt"
	"os"
)

func main() {
	images, err := loadImages()

	if err != nil {
		fmt.Errorf("Error reading images: %s", err)
		os.Exit(1)
	}

	fmt.Println(images)
}
