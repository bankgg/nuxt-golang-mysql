package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "root:password@tcp(db:3306)/sample?charset=utf8mb4&parseTime=true"
	var db *gorm.DB

	for {
		log.Println("Sleeping for 5 secs")
		time.Sleep(time.Second * 5)

		d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			continue
		}

		db = d
		break
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "test1", Email: "test1@email.com"})

	user := new(User)
	db.First(user)

	db.Delete(user)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s (%s), id: %v", user.Name, user.Email, user.ID),
		})
	})
	r.Run()
}
