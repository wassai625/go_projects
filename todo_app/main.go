package main

import (
	"fmt"
	// "log"
	// "os/user"

	// "log"
	"todo_app/app/controllers"
	"todo_app/app/models"
	// "todo_app/config"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// log.Println(err)
	// }

	// valid, _ := session.CheckSession()
	// fmt.PrintIn(valid)
}
