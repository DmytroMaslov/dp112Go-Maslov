package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"tasksElementary/task5"
)


func main(){
	http.HandleFunc("/", handlerTask5)
	http.ListenAndServe("localhost:8080", nil)
}


func handlerTask5 (w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.PostForm)
	inputS := r.Form["js"]
	fmt.Println(inputS)
	jsonData :=[]byte(inputS[0])
	var elementoryTasks = new(task5.TicketRange)
	err := json.Unmarshal(jsonData, &elementoryTasks)
	if err != nil{
		fmt.Fprintf(w, "Error iput %s", err)
	}
	winMethod :=task5.Run(elementoryTasks.Min, elementoryTasks.Max)
	fmt.Println(winMethod)
	fmt.Fprint(w, "URL.path=%s", winMethod)
}