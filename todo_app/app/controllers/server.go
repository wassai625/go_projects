package controllers

import (
	"io"
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
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

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
