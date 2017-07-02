package main

/*
import (
	"tasksElementary/task2"
	"tasksElementary/task3"
	"encoding/json"
	"fmt"
)

package tools
import (
"encoding/json"
"fmt"
"task3"
"task2"
)
type Params struct {
	Task2 task2.Envelope
	Task3 []task3.Triagle
}

func ParseJson () {

	//TODO read from file Json
	var jsonP = []byte(`{
	"Task2":
	{"S1":10,"S2":20},
	"Task3":
	[{"Name":"one","A":10,"B":15,"C":20},
	{"Name":"two","A":12,"B":16,"C":22},
	{"Name":"tree","A":13,"B":17,"C":24}]
	}`)
	var param Params
	err :=json.Unmarshal(jsonP, &param)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v", param)

}
*/