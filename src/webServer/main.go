package main

import (
	"net/http"
	"fmt"
	"tasksElementary/task5"
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
	r.ParseForm()
	inputS := r.Form["js"]
	fmt.Println(inputS)
	jsonData :=[]byte(inputS[0])
	str, err:=task5.Run(jsonData)
	fmt.Println(str)
	var resp = new(Resp)
	resp.Resp = str
	fmt.Println(resp.Resp)
	if err != nil {resp.Error = err.Error()}
	json.NewEncoder(w).Encode(resp)


}