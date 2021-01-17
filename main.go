package main

import (
	"fmt"

	"./app/controllers"
	"./app/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
