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

	// u := models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "texttext"

	u, _ := models.GetUser(1)
	u.Name = "updateUser"
	u.DeleteUser()
	u, _ = models.GetUser(1)

	fmt.Println(u)
}
