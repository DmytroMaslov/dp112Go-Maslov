package main

import (
	"fmt"
	"task1"
	"task2"
	"task3"
	//"task4"
	"task5"
	"task6"
	"task7"
)

func main() {
//Task1
	fmt.Println("Task1:")
	ok :=task1.Task1(5,5,"w")
	if ok !=nil{
		fmt.Println("Task1- error:", ok)
	}

	//task2
	fmt.Println("Task2:")
	en1 := task2.Envelope{S1: 10, S2:5}
	en2 := task2.Envelope{S1: 15, S2: 20}
	rerult := task2.GetSmallEnvelope(en1, en2)
	fmt.Println(rerult)

//task 3
	fmt.Println("Task3:")
	tr1:=task3.Triagle{Name:"first", A:10, B:10, C:10}
	tr2:=task3.Triagle{Name:"second", A:20, B:20, C:20}
	tr3:=task3.Triagle{Name:"third", A:15, B:15, C:15}
	tr4:=task3.Triagle{Name:"four", A:11.045, B:1115.8263, C:15.234234}
	s:= []task3.Triagle{tr1, tr2, tr3, tr4}
	name := task3.SortTriagle(s)
	fmt.Println(name)

//task5
	fmt.Println("Task5:")
	winMethod := task5.GetWinerTicketResult(0, 999999)
	fmt.Println(winMethod)
//task6
	fmt.Println("Task6:")
	task6.WriteToFileNumbersRow(11,11, "test2.txt")


//task7
	fmt.Println("Task7:")
	er := task7.CalculateFibonachiRow()
	if er !=nil{
		fmt.Println(er)
	}
}



