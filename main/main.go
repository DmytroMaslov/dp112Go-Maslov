package main


import (
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"fmt")

func main() {
//Task1
	fmt.Println("Task1:")
	l:=5
	w:= 5
	c:="w"
	err:= task1.IsValid(l,w,c)
	if err !=nil{
		fmt.Println("error:", err)
	} else {
		task1.Task1(5,5,"w")
	}

	//task2
	fmt.Println("Task2:")
	en1 := task2.Envelope{S1: 10, S2:5}
	en2 := task2.Envelope{S1: 15, S2: 20}
	err = task2.IsValid(en1, en2)
	if err != nil{
		fmt.Println("error:", err)
	} else {
		rerult := task2.GetSmallEnvelope(en1, en2)
		fmt.Println(rerult)
	}


//task 3
	fmt.Println("Task3:")
	tr1:=task3.Triagle{Name:"first", A:10, B:10, C:10}
	tr2:=task3.Triagle{Name:"second", A:20, B:20, C:20}
	tr3:=task3.Triagle{Name:"third", A:15, B:15, C:15}
	tr4:=task3.Triagle{Name:"four", A:11.045, B:15.8263, C:15.234234}
	s:= []task3.Triagle{tr1, tr2, tr3, tr4}
	err = task3.IsValid(s)
	if err != nil{
		fmt.Println("error:", err)
	} else {
		name := task3.SortTriagle(s)
		fmt.Println(name)
	}


//task4
	fmt.Println("Task4:")
	var n uint64 = 12332145789987
	err = task4.IsValid(n)
	if err != nil{
		fmt.Println("error:", err)
	} else {
		fmt.Println(task4.GetMaxPalidrom(n))
	}


//task5
	fmt.Println("Task5:")
	min :=0
	max:= 999999
	err = task5.IsValid(min, max)
	if err != nil{
		fmt.Println("error:", err)
	} else {
	winMethod := task5.GetWinerTicketResult(0, 999999)
	fmt.Println(winMethod)
	}
//task6
	fmt.Println("Task6:")
	lengRow := 11
	minSquare := 11
	err = task6.IsValid(lengRow, minSquare)
	if err != nil{
		fmt.Println("error:", err)
	} else {
		task6.WriteToFileNumbersRow(11, 11)
	}

//task7
	fmt.Println("Task7:")
	err = task7.CalculateFibonachiRow()
	if err !=nil{
		fmt.Println(err)
	}

}



