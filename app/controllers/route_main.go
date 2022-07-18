package controllers

import (
	"log"
	"net/http"

	"todo/app/models"
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

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}

}

func todoSave(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		todo := &models.Todo{ID: id, Content: content, UserId: user.ID}
		if err := todo.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		if err := todo.DeleteTodo(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
