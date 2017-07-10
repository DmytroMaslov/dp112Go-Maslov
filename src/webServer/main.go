package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	s "strings"
	"tasksElementary/taskManager"
	"encoding/json"
	"html/template"
)


func main(){
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/task/", handlerTask)
	http.ListenAndServe(":8080", nil)
}
type Resp struct {
	Resp string
	Error string
}
type Request struct {
	Task string
	Data json.RawMessage
}


func handlerTask(w http.ResponseWriter, r *http.Request){
	pos := s.LastIndex(r.RequestURI, "/")
	task := r.RequestURI[pos+1:]
	fmt.Println(task)
	inp, err := ioutil.ReadAll(r.Body)
	fmt.Println(inp)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()
	res, err := taskManager.RunTask(task, inp)
	var resp = new(Resp)
	resp.Resp = res
	fmt.Println(resp.Resp)
	if err != nil {resp.Error = err.Error()}
	json.NewEncoder(w).Encode(resp)

}

func IndexHandler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
