package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/badtripdude/gh-analyzer/internal/app"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	service *app.Service
}

func NewConsumer(brokers []string, topic, groupID string, s *app.Service) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &Consumer{reader: r, service: s}
}

func (c *Consumer) Start(ctx context.Context) error {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Got message: %s\n", string(m.Value))

		if err := c.service.HandleAnalysisResults(m.Value); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}
