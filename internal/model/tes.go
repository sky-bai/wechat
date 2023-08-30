package model

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	Mobile    string `gorm:"colum:mobile;type:varchar(11)" json:"mobile"`
}

// TableName MiniProgram's table name
func (*PersonAddress) TableName() string {
	return "person_addresses"
}
