package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	firstname string `gorm:"not null"`
	lastname  string `gorm:"not null" `
	Email     string `gorm:"not null" `
	taks   []Taks
}