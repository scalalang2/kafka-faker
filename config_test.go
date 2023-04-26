package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDecodeConfig(t *testing.T) {
	content := `
kafka:
  brokers: [localhost:9092]
generators:
  - topic: topic1
    partitionKey: none
    schema: |-
      {
        "from_address": "::ethereum_address()",
        "to_address": "::ethereum_address()",
        "amount": "::number(1,10)",
        "timestamp": "::timestamp()"
      }
    number: 100
    loop: true
    delay: 5s
`

	expected := &Config{
		Kafka: KafkaConfig{
			Brokers: []string{"localhost:9092"},
		},
		Generators: []GeneratorConfig{
			{
				Topic:        "topic1",
				PartitionKey: "none",
				Schema: Schema{
					"from_address": &Func{
						Name: "ethereum_address",
					},
					"to_address": &Func{
						Name: "ethereum_address",
					},
					"amount": &Func{
						Name: "number",
						Args: []string{"1", "10"},
					},
					"timestamp": &Func{
						Name: "timestamp",
					},
				},
				Number: 100,
				Loop:   true,
				Delay:  time.Second * 5,
			},
		},
	}

	config, err := DecodeConfig([]byte(content))
	require.NoError(t, err)
	require.Equal(t, expected, config)
}
