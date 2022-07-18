package main

import (
	"fmt"
	"todo/app/controllers"
	"todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	// user, _ := models.GetUserByEmail("test@test.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
