package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
}