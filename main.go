package main

import (
	"fmt"
	consumer "kafka-again/consumer"
	producer "kafka-again/producer"
	"time"
)

func main() {
	fmt.Println("Startig the GO Producer")

	go producer.WriteKafka()
	go consumer.StartKafka() // separate go routine

	fmt.Println("Kafka has been started")

	time.Sleep(20 * time.Second)
}

/*
docs :

This application will have a single kafka broker, one kafka producer and one kafka consumer
Randomly messages will be created and sent to the broker.
single topic : mytopic1
broker address : localhost:9092
*/
