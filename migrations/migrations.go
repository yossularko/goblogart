package main

import (
	"goblogart/inits"
	"goblogart/models"
	"log"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	log.Println("running migrations")
	if err := inits.DB.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Println("err migrations: ", err.Error())
		return
	}

	log.Println("success migrations")
}
