package producer

import (
	"context"
	"log"
	"math/rand"

	"github.com/segmentio/kafka-go"
)

// randomly create a message that is to be sent to the kafka broker
var msgs = []string{"m1", "m2", "m3", "m4", "m5", "m6"}

func CreateMessage() string {
	randNum := rand.Intn(5)
	return msgs[randNum]
}

// function to write the messages to kafka
func WriteKafka() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"), // address of the kafka cluster or broker
		Topic:    "mytopic1",
		Balancer: &kafka.LeastBytes{},
	}

	// infinite loop so as the messages are sent in 1 sec intervals
	for {
		err := w.WriteMessages(context.Background(),
			// kafka.Message{
			// 	Key:   []byte("Key-A"),
			// 	Value: []byte("Hello World!"),
			// },
			// kafka.Message{
			// 	Key:   []byte("Key-B"),
			// 	Value: []byte("One!"),
			// },
			// kafka.Message{
			// 	Key:   []byte("Key-C"),
			// 	Value: []byte("Two!"),
			// },
			kafka.Message{
				Key:   []byte("Key-D"),
				Value: []byte(CreateMessage()),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)

			if err := w.Close(); err != nil {
				log.Fatal("failed to close writer:", err)
			}
		}
	}
}
