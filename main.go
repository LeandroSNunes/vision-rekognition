package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Classifier interface {
	Classify(file io.Reader) (map[string]float64, error)
}

func main() {
	config, err := InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	images, err := loadImages()
	if err != nil {
		fmt.Errorf("Error reading images: %s", err)
		os.Exit(1)
	}

	rekog, err := NewRekognitionClient(config.AwsRegion, config.AwsAccessKeyID, config.AwsSecretAccessKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vision, err := NewVisionClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, img := range images {
		wg.Add(1)

		go func(img string, wg *sync.WaitGroup) {
			report(img, vision, rekog)
			wg.Done()
		}(img, &wg)
	}

	wg.Wait()
}

func report(image string, vision, rekog Classifier) error {
	fileName := strings.Replace(image, "images/", "", 1)
	messages := make(chan string, 2)

	go func(filePath, fileName string) {
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer file.Close()

		cat, _ := rekog.Classify(file)
		messages <- fmt.Sprintf("%s -- REKOG: %v\n", fileName, formatResult(cat))
	}(image, fileName)

	go func(filePath, fileName string) {
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer file.Close()

		cat, _ := vision.Classify(file)
		messages <- fmt.Sprintf("%s -- Vision: %v\n", fileName, formatResult(cat))
	}(image, fileName)

	report := ""
	report += <-messages
	report += <-messages
	report += "------------------------"
	fmt.Println(report)
	return nil
}

func formatResult(r map[string]float64) string {
	var res string
	for k, v := range r {
		res += fmt.Sprintf("%s: %.2f ", k, v)
	}

	return res
}
