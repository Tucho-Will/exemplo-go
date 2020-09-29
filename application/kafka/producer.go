package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func Producer() {
	fmt.Println("Producer:::1.0")
	syncProducer, err := sarama.NewSyncProducer([]string{KafkaServer}, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg := &sarama.ProducerMessage{
			Topic: KafkaTopic,
			Value: sarama.ByteEncoder("MESSAGE" + time.Now().Format(time.RFC3339)),
		}

		_, _, err = syncProducer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		println(msg)
		time.Sleep(time.Second)
	}
}
