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
	log.Println("A random message to send is created")
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
		// random message created that is to be sent to the topic : mytopic1 in the kafka
		msg := CreateMessage()
		// sending the message to the kafka server
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-D"),
				Value: []byte(msg),
			},
		)
		log.Println("Message sent to the kafka")

		if err != nil {
			log.Fatal("failed to write messages:", err)

			if err := w.Close(); err != nil {
				log.Fatal("failed to close writer:", err)
			}
		}
	}
}
