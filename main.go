package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponderWriter, r *http.Request) {
	template, err := template.ParseFiles("views/index.html", "views/footer.html", "views/header.html")
	if err != nil {
		fmt.Fprinf(w, err.Error())
	}
	template.ExecuteTemplate(w, "index", nil)
}

func initHttpServer() {
	http.Handle("/assets/", http.StripPrefix("/assets/"http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(':3000', nil)
}

func main() {
	fmt.Println("Starting GO server on port :3000")
	initHttpServer()
}