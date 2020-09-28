package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	consumer:= Consumer{}
	consumer.Consume()
}

type Consumer struct {
}

const KafkaServer2 = "localhost:9092"
const KafkaTopic2= "meu-topico-legal"


func (c *Consumer) Consume() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	group, err := sarama.NewConsumerGroup([]string{KafkaServer2}, "my-group", config)
	if err != nil {
		panic(err)
	}

	go func() {
		for err := range group.Errors() {
			panic(err)
		}
	}()

	func() {
		ctx := context.Background()
		for {
			topics := []string{KafkaTopic2}
			err := group.Consume(ctx, topics, c)
			if err != nil {
				fmt.Printf("kafka consume failed: %v, sleeping and retry in a moment\n", err)
				time.Sleep(time.Second)
			}
		}
	}()
}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("consumed a message: %v\n", string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
