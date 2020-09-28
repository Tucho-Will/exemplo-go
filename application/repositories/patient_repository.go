package repositories

import (
	"github.com/Tucho-Will/exemplo-go/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type PatientRepository interface {
	Insert(patient *domain.Patient) (*domain.Patient, error)
}

type PatientRepositoryDb struct {
	Db *gorm.DB
}

func (repo PatientRepositoryDb) Insert(patient *domain.Patient) (*domain.Patient, error) {

	err := patient.Prepare()

	if err != nil {
		log.Fatalf("Error during the user prepare: %e", err)
		return patient, err
	}

	err = repo.Db.Create(patient).Error

	if err != nil {
		log.Fatalf("Error to persist patient: %e", err)
		return patient, err
	}

	return patient, nil

}
