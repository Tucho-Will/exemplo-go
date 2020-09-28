package domain

//importa o module
import uuid "github.com/satori/go.uuid"

//Struct é uma estrutura de dados
type Patient struct {
	BaseEntity
	Name string `json:"patient_name" gorm:"typpe:varchar(120)"`
	Cid  string `json:"patient_cid" gorm:"typpe:varchar(35)"`
	Code string `json:"-" gorm:"typpe:varchar(80);unique_index"`
}

//Prepara o paciente para ser persistido
func (patient *Patient) Prepare() error {

	patient.Code = uuid.NewV4().String()
	patient.Cid = "Teste-10"

	return nil
}
