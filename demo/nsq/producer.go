package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topicName := "topic"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	for i := 0; i < 99999; i++ {
		messageBody := []byte(fmt.Sprintf("hello %d", i))
		err = producer.Publish(topicName, messageBody)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(1))
	}

	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()
}
