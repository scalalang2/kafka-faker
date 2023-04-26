# Kafka Faker 
`Kafka Faker` is a fake message generator for kafka. It generates fake messages based on a schema and sends them to a kafka topic.

It's useful for a service that consumes messages from kafka, stream processing with tools (e.g. Apache Flink) or testing.

## Configuration
```yaml
kafka:
  brokers: [localhost:9092]
messages:
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
```

## Run
```shell
$ kafka-faker -c config.yaml
```

## Supported Types
- JSON

## LICENSE
MIT