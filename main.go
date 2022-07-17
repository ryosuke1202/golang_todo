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

	/*
		u, _ := models.GetUser(2)
			fmt.Println(u)
			u.Name = "henkou2"
			u.Email = "henkou2@henkou.com"

			u.UpdateUser()
			u, _ = models.GetUser(u.ID)
			fmt.Println(u)


			u.DeleteUser()
			user, _ := models.GetUser(3)
			user.CreateTodo("3回目のTodoです")
			fmt.Println(user.ID)

			todo, err := models.GetTodo(user.ID)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(todo)

			todos, _ := models.GetTodos()
			for _, v := range todos {
				fmt.Println(v)
			}
			user, _ := models.GetUser(4)
			todos, _ := user.GetTodoByUser()
			for _, v := range todos {
				fmt.Println(v)
			}
			todo, _ := models.GetTodo(1)
			todo.Content = "更新しました"
			todo.UpdateTodo()

			todo, _ := models.GetTodo(1)
			todo.DeleteTodo()
	*/
}
