package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	handler1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-1\n")
	}
	handler2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-2\n")
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/foo/", handler1)
	http.HandleFunc("/bar/", handler2)
	http.HandleFunc("/signup/", signup)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
