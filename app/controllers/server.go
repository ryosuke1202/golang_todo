package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo/app/models"

	"todo/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (models.Session, error) {
	cookie, err := r.Cookie("_cookie")
	var session models.Session
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("Invalid Session")
		}
	}
	return session, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top) // urlを登録して"/"にアクセスするとtopというハンドラーの処理を実行する
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
