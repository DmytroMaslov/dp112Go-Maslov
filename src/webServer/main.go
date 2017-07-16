/*
Contract
respons
one task:
.../task/N
{width:int, height: int, symbol: string}
[{ "width": 8, "height": 5}, {"width": 6, "height": 9 }]
request:
{task: N, resp: string, reason: string}  //если ок по resp заполнен reason пустой

allTask:
.../tasks


*/
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
	http.HandleFunc("/tasks", handlerTasks)
	http.ListenAndServe(":8080", nil)
}
//{task: N, resp: string, reason: string}
type Respons struct {
	Task string `json: task`
	Resp string `json:"resp"`
	Reason string `json:"reason"`
}
type Request struct {
	Task string
	Data json.RawMessage
}

type RespAll struct {
	Name string
	Resp Respons
}


func handlerTask(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.RequestURI)
	indx := s.LastIndex(r.RequestURI, "/")+1
	task := r.RequestURI[indx:]
	inpData, err := ioutil.ReadAll(r.Body)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()
	res, err := taskManager.RunTask(task, inpData)
	var resp = new(Respons)
	resp.Task = task
	resp.Resp = res
	fmt.Println(resp.Resp)
	if err != nil {resp.Reason = err.Error()}
	json.NewEncoder(w).Encode(resp)

}

func handlerTasks (w http.ResponseWriter, r *http.Request){
	fmt.Println("Body", r.Body)
	inp, err := ioutil.ReadAll(r.Body)
	fmt.Println(inp)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()

	allTasks :=make(map[string]json.RawMessage)
	fmt.Println("unmarhale")
	err = json.Unmarshal(inp, &allTasks)

	if err !=nil{
		fmt.Fprintf(w, err.Error())
	}
	respAll := make([]Respons, len(allTasks))
	fmt.Println("all", allTasks)
	i:=0
	for k, v := range allTasks{
		res, err := taskManager.RunTask(k, []byte(v))
		var resp = new(Respons)
		resp.Task = k
		resp.Resp = res
		if err != nil {resp.Reason = err.Error()}
		respAll[i] = *resp
		i++
	}
	json.NewEncoder(w).Encode(respAll)

}


func IndexHandler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
