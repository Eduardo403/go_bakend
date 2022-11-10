package models

import "gorm.io/gorm"

type Taks struct {
	gorm.Model

	Title string `gorm:"not null;unique_index" `
	Descriction string 
	 Done bool `gorm:"default:false"`
	UserID  uint
}