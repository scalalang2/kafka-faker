package main

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/inconshreveable/log15"
)

// Entry represents a message to be sent to Kafka.
// only Value is encoded to JSON. Topic and Key are used to define the destination and partition.
type Entry struct {
	Topic string
	Key   string
	Value string
}

// Sender is a worker that sends messages to Kafka.
type Sender struct {
	producer *kafka.Producer
	box      chan *Entry
}

func NewSender(cfg KafkaConfig) (*Sender, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
	})
	if err != nil {
		return nil, err
	}

	// reporting error
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Error("failed to send message", "error", ev.TopicPartition)
				}
			}
		}
	}()

	return &Sender{
		producer: p,
		box:      make(chan *Entry, 100),
	}, nil
}

// Send asynchronously sends a message to Kafka.
func (s *Sender) Send(msg *Entry) {
	s.box <- msg
}

func (s *Sender) Run(ctx context.Context) {
	cnt := 0
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			s.producer.Flush(15 * 1000)
			log.Info("Sender is stopped by context")
			return
		case <-ticker.C:
			log.Info("Sent fake message", "count", cnt)
			cnt = 0
		case msg := <-s.box:
			cnt++
			if err := s.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &msg.Topic, Partition: kafka.PartitionAny},
				Value:          []byte(msg.Value),
			}, nil); err != nil {
				log.Error("failed to send message to kafka", "error", err, "entry", msg)
				panic(err)
			}
		}
	}
}
