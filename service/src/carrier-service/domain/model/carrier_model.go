package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Carrier struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(256);not null"`
	Address   string    `json:"address" gorm:"type:varchar(256);not null;default:empty"`
	Active    bool      `json:"active" gorm:"type:boolean;not null;default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Carrier) NewCarrier(name string, address string, active bool) error {
	if active != true {
		return errors.New("carrier is not active")
	}

	if address == "" {
		return errors.New("carrier address is empty")
	}

	c.Name = name
	c.Address = address
	c.Active = active
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return nil
}

func (c *Carrier) UpdateCarrierAddress(address string) error {
	if address == "" {
		return errors.New("carrier address is empty")
	}

	c.Address = address

	return nil
}

func (c *Carrier) UpdateCarrierActiveStatus(activeStatus bool) {
	c.Active = activeStatus
}
