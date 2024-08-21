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


***************RUN THE APPLICATION******************
1. open terminal and start zookeeper : run the below cmd
/usr/local/opt/kafka/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties

2. open another terminal and create a kafka topic or whatever stuff you want
then start the kafka server
/usr/local/opt/kafka/bin/kafka-server-start /usr/local/etc/kafka/server.properties

3. go run main.go

for more reference
https://medium.com/@bectorhimanshu/setting-up-installing-and-starting-the-kafka-server-on-macos-7d505e32039a

*/
