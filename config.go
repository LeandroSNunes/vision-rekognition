package main

import "github.com/spf13/viper"

type Config struct {
	AwsRegion          string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AwsRegion:          viper.GetString("aws_region"),
		AwsAccessKeyID:     viper.GetString("aws_access_key_id"),
		AwsSecretAccessKey: viper.GetString("aws_secret_access_key"),
	}

	return cfg, nil
}
