package main

import (
	"./app/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// u := &models.User{}
	// u.Name = "apple"
	// u.Email = "test@example.com"
	// u.PassWord = "texttext"

	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("apple todo")
	// fmt.Println(user)

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
