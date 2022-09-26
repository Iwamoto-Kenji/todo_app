package main

import (
	"fmt"
	"todo_app/app/models"

)

func main() {
	fmt.Println(models.Db)
	u, _ := models.GetUser(1)
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	
	fmt.Println(u)
}