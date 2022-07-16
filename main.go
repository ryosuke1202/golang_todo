package main

import (
	"fmt"
	"todo/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("testtest")

	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "text@example.com"
		u.PassWord = "password"
		// fmt.Println(u)

		u.CreateUser()
	*/

	u, _ := models.GetUser(2)
	fmt.Println(u)
	u.Name = "henkou2"
	u.Email = "henkou2@henkou.com"

	u.UpdateUser()
	u, _ = models.GetUser(u.ID)
	fmt.Println(u)

}
