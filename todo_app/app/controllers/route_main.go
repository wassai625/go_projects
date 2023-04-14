package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/models/views/templates/top.html")
	if err != nil {
		fmt.Println("aaaaa")
	}
	t.Execute(w, "hello")
}
