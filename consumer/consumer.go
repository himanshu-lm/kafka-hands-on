package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// this is the kafka consumer
func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "mytopic1",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Some error occured : %v", err)
			continue
		}
		fmt.Printf("The message received by the consumer is : %v\n", string(msg.Value))
		log.Println("Message received from kafka to consumer")
	}
}
