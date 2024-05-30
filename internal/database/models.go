package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Description sql.NullString `json:"description"`
}
