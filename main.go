package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	t, _ := models.GetTodo(1)
	fmt.Println(t)
}