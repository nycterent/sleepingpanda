package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User simple test gorm
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.Create(&User{Username: "test", Password: "test", Email: "test@test.lt"})

	var user User
	db.First(&user, 1)
	db.First(&user, "username = ?", "test")
	db.Model(&user).Update("Username", "test")
	fmt.Println(user)
	db.Delete(&user)
}
