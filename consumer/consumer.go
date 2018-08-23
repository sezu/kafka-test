package main

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"syscall"
	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "kafka-test"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.42.65:9092"},
		GroupID:   "consumer-group-id",
		Topic:     topic,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		// CommitInterval: time.Second, // flushes commits to Kafka every second
	})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Got %s signal. Aborting...\n", sig)
			run = false
		default:
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				fmt.Printf("Consumer error: %v\n", err)
				run = false
				break;
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}

	r.Close()
}
