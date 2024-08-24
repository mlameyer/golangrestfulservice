package model

import (
	"gorm.io/gorm"
	"time"
)

type Carrier struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(256);not null"`
	InService bool      `json:"in_service" gorm:"type:boolean;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
