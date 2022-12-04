package main

import (
	"fmt"
	"log"

	"github.com/onainadapdap1/golang-gorm-postgresql/initializers"
	"github.com/onainadapdap1/golang-gorm-postgresql/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		fmt.Println("something wrong when trying to migrate user model.")
	}
	fmt.Println("? Migration complete")
}

