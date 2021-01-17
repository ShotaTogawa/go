package main

import (
	"fmt"

	"./app/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "texttext"

	u.CreateUser()

	user, _ := models.GetUser(2)
	user.CreateTodo("First todo")
	fmt.Println(user)
}
