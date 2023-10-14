package models

import (
	"time"
)

type User struct {
	Id        string     `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	LastName  string     `gorm:"not null"`
	FullName  string     `gorm:"not null";"omitempty"`
	Email     string     `gorm:"not null"`
	BirthDate *time.Time `gorm:"not null"`
}
