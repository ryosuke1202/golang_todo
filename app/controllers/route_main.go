package controllers

import (
	"log"
	"net/http"
)

// この中の引数はお決まりの書き方。こうすればハンドラーとして登録できる
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top") // ログインしていない場合
	} else {
		http.Redirect(w, r, "/todos", 302) // ログインしている場合
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodoByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}

}
