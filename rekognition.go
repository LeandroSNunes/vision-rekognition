package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type RekogClient struct {
	engine *rekognition.Rekognition
}

func (c *RekogClient) Classify(file io.Reader) (map[string]float64, error) {
	bytes, err := fileForBytes(file)
	if err != nil {
		return nil, err
	}
	image := rekognition.Image{Bytes: bytes}
	input := rekognition.DetectModerationLabelsInput{Image: &image}
	labels, err := c.engine.DetectModerationLabels(&input)
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)
	for _, label := range labels.ModerationLabels {
		result[*label.Name] = *label.Confidence
	}

	return result, nil
}

func NewRekognitionClient(region, access_key_id, secret_access_key string) (*RekogClient, error) {
	credValues := credentials.Value{
		AccessKeyID:     access_key_id,
		SecretAccessKey: secret_access_key,
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credValues),
	})

	if err != nil {
		return nil, fmt.Errorf("rekog: new session -- %s", err)
	}

	engine := rekognition.New(sess)
	return &RekogClient{engine: engine}, nil
}

func fileForBytes(file io.Reader) ([]byte, error) {
	return ioutil.ReadAll(file)
}
