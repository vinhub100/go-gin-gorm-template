package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Articles struct {
	gorm.Model
	Heading string `gorm:"size:255"`
	Body    string
}
