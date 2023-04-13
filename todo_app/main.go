package main

import (
	"fmt"

	// "log"
	"todo_app/app/models"
	"todo_app/app/models/controllers"
	// "todo_app/config"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
