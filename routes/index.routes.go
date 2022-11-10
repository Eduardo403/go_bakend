package routes

import "net/http"

func HomeHandele(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}