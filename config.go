package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Kafka    KafkaConfig     `yaml:"kafka"`
	Messages []MessageConfig `yaml:"messages"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

type MessageConfig struct {
	Topic        string        `yaml:"topic"`
	PartitionKey string        `yaml:"partitionKey"`
	Schema       Schema        `yaml:"schema"`
	Number       int           `yaml:"number"`
	Loop         bool          `yaml:"loop"`
	Delay        time.Duration `yaml:"delay"`
}

// LoadConfig read a config from a file
func LoadConfig(filename string) (*Config, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return DecodeConfig(contents)
}

func DecodeConfig(content []byte) (*Config, error) {
	var config Config

	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
