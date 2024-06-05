package models

import (
	"github.com/VCulev/BookStore-Management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	var err error
	db, err = config.Connect()
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic("failed to auto migrate database")
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}
