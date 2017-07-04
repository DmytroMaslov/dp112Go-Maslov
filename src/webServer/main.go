package main

import (
	"net/http"
	"fmt"
	//"tasksElementary/task5"
	"encoding/json"
	"io/ioutil"
)


func main(){
	http.HandleFunc("/", handlerTask5)
	http.ListenAndServe("localhost:8080", nil)
}
type Resp struct {
	Resp string
	Error string
}
type Request struct {
	Name string
	Param []byte
}
func handlerTask5 (w http.ResponseWriter, r *http.Request){
	inp, err := ioutil.ReadAll(r.Body)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()
	var req = new(Request)
	err = json.Unmarshal(inp, &req)
	fmt.Println(req.Name)
	//fmt.Println(inp)
	//inputS := r.Form["js"]
	//jsonData :=[]byte(inputS[0])
	/*str, err:=task5.Run(inp)
	fmt.Println(str)
	var resp = new(Resp)
	resp.Resp = str
	fmt.Println(resp.Resp)
	if err != nil {resp.Error = err.Error()}
	json.NewEncoder(w).Encode(resp)*/


}