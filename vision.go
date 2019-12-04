package main

import (
	"context"
	"fmt"
	"io"

	gv "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

const CREDENTIAL_FILE_PATH = "vision-credentials.json"

// Client allow classification
type VisionClient struct {
	engine *gv.ImageAnnotatorClient
	ctx    context.Context
}

func (c *VisionClient) detectSafeSearch(file io.Reader) (*pb.SafeSearchAnnotation, error) {
	image, err := gv.NewImageFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("vision: NewImageFromReader -- %s", err)
	}

	props, err := c.engine.DetectSafeSearch(c.ctx, image, nil)
	if err != nil {
		return nil, fmt.Errorf("vision: DetectSafeSearch -- %s", err)
	}

	return props, nil
}

func (c *VisionClient) Classify(file io.Reader) (map[string]float64, error) {
	response := make(map[string]float64)

	props, err := c.detectSafeSearch(file)
	if err != nil {
		return response, err
	}

	response["Adult"] = float64(props.Adult)
	response["Medical"] = float64(props.Medical)
	response["Racy"] = float64(props.Racy)
	response["Spoof"] = float64(props.Spoof)
	response["Violence"] = float64(props.Violence)

	return response, nil
}

// NewClient start the Google Vision and return the client for classification
func NewVisionClient() (*VisionClient, error) {
	credOption := option.WithCredentialsFile(CREDENTIAL_FILE_PATH)
	ctx := context.Background()

	cli, err := gv.NewImageAnnotatorClient(ctx, credOption)
	if err != nil {
		return nil, fmt.Errorf("vision: NewImageAnnotatorClient -- %s", err)
	}

	return &VisionClient{engine: cli, ctx: ctx}, nil
}
