package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDecodeConfig(t *testing.T) {

	content := `
kafka:
  brokers: localhost:9092
generators:
  - topic: topic1
    partitionKey: none
    template: hello
    number: 100
    loop: true
    delay: 5s
`

	expected := &Config{
		Kafka: KafkaConfig{
			Brokers: "localhost:9092",
		},
		Generators: []GeneratorConfig{
			{
				Topic:        "topic1",
				PartitionKey: "none",
				Template:     Template("hello"),
				Number:       100,
				Loop:         true,
				Delay:        time.Second * 5,
			},
		},
	}

	config, err := DecodeConfig([]byte(content))
	require.NoError(t, err)
	require.Equal(t, expected, config)
}
