//package main
//
//import (
//	"fmt"
//	"github.com/Tucho-Will/exemplo-go/application/repositories"
//	"github.com/Tucho-Will/exemplo-go/domain"
//	"github.com/Tucho-Will/exemplo-go/framework/utils"
//	"log"
//)
//
//func main() {
//
//	db := utils.ConnectDb()
//
//	patient := domain.Patient{
//		BaseEntity: domain.BaseEntity{},
//		Name:       "Fagner",
//	}
//
//	patientRepo := repositories.PatientRepositoryDb{Db:db}
//
//	result, err := patientRepo.Insert(&patient)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(result)
//}

package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

const KafkaServer = "localhost:9092"
const KafkaTopic = "meu-topico-legal"

func main() {
	fmt.Println("Producer:::1.0")
	syncProducer, err := sarama.NewSyncProducer([]string{KafkaServer}, nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{
			Topic: KafkaTopic,
			Value: sarama.ByteEncoder(fmt.Sprintf(`hello word %t`, i) + time.Now().Format(time.RFC3339)),
		}

		_, _, err = syncProducer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		println(msg)
	}
}
