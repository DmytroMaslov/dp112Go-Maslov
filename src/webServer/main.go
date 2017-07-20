package main

import (
	"net/http"
	"webServer/handlers"
)


func main(){
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/task/", handlers.HandlerTask)
	http.HandleFunc("/tasks", handlers.HandlerTasks)
	http.ListenAndServe(":8080", nil)
}

