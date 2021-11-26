package main

import (
	"blog-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Articles{})

	// Create
	db.Create(&models.Articles{Heading: "Headind...", Body: "Some Body... Some Body"})

	// Read
	var article models.Articles
	db.First(&article, 1) // find product with id 1
	// db.First(&article, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&article).Update("Heading", "Heading2....")

	// Delete - delete product
	// db.Delete(&article)
}
