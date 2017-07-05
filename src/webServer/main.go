package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	s "strings"
	"tasksElementary/taskManager"
	"encoding/json"
)


func main(){
	http.HandleFunc("/", handlerTask5)
	http.ListenAndServe("localhost:8080", nil)
}
type Resp struct {
	Resp string
	Error string
}

func handlerTask5 (w http.ResponseWriter, r *http.Request){
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