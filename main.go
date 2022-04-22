package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 109})

	// Read
	var product Product
	db.First(&product, "code = ?", "D42")
	/*
	 *  db.First(&product, 1)                 // Finds product with int primary key 1
	 *  db.First(&product, "code = ?", "D42") // Finds product with code D42
	 *
	 *  // Update
	 *  db.Model(&product).Update("Price", 200) // single field
	 *  db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	 *
	 *  // Deletes
	 *  db.Delete(&product, 1)
	 */
}
