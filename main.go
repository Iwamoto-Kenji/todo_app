package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	user, _ := models.GetUser(1)

	todos, _ := user.GetTodosByUser()
	for _,v := range todos {
		fmt.Println(v)
	}
}