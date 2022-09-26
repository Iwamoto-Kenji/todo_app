package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	u :=&models.User{}

	u.Name = "test"
	u.Email = "test@example.com"
	u.Name = "test"

	u.CreateUser()
}