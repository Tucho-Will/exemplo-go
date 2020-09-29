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

import "github.com/Tucho-Will/exemplo-go/application/kafka"

func main() {

	//GORoutine é similar a uma thread
	go kafka.Producer()

	consumer := kafka.Consumer{}
	consumer.Consume()

}
