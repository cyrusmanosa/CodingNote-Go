package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Set a new HTTP request multiplexer
	mux := httprouter.New()
	// Listen to root path
	mux.GET("/", index)
	// Listen to CSS assets
	mux.ServeFiles("/css/*filepath", http.Dir("public/css"))
	// Listen to JavaScript assets
	mux.ServeFiles("/js/*filepath", http.Dir("public/js"))
	// Set the parameters for a HTTP server
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	// Listen requests
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, "My Application with CSS")
}
