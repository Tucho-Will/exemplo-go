package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type BaseEntity struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

func (baseEntity *BaseEntity) BeforeCreate(scope *gorm.Scope) error {

	err := scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during generated id")
	}

	return nil
}
