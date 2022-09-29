package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	user, _ := models.GetUser(1)
	user.CreateTodo("testTodo2")

	todos, _ := models.GetTodos()
	for _,v := range todos {
		fmt.Println(v)
	}
}