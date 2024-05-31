package database

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Author      string `json:"author" gorm:"not null"`
	Genre       string `json:"genre"`
	Description string `json:"description" gorm:"defualt:'No description'"`
}
