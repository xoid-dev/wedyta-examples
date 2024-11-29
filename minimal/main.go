package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pa-pe/wedyta"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User demo model
type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:255" json:"name"`
}

// TableName for gorm
func (User) TableName() string {
	return "user"
}

func main() {
	// wedyta does not depend on any specific database type; using SQLite for a quick in-memory setup
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Prepare demo table & data
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database")
	}
	db.Create(&User{Name: "Alice"})
	db.Create(&User{Name: "Bob"})

	// Initialize Gin router
	router := gin.Default()

	// minimal wedyta setup
	wedyta.New(router, db, nil)

	// Run the server on port 8080
	panic(router.Run(":8080"))

	// Visit http://127.0.0.1:8080/wedyta/Users to see the rendered table
}
