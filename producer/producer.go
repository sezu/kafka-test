package main

import (
	"fmt"
	"time"
	"context"
	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "kafka-test"

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: 	[]string{"192.168.42.65:9092"},
		Topic: 		topic,
		Balancer: 	&kafka.LeastBytes{},
	})

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
			case _ = <-ticker.C:

			fmt.Println("Writing")

			w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key"),
					Value: []byte("This is a test."),
				},
			)
		}
	}

	w.Close()
}