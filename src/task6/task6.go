/*6.	Числовая последовательность
Вывестив файл
через запятую
ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше (больше) заданного m.
Входные параметры: длина и значение минимального квадрата
Выход: nil если сохранение удалось и err в противном случае
*/
package task6

import (
	"math"
	"os"
	"errors"
	"strconv"
	"fmt"
)

func WriteToFileNumbersRow(lengRow, minSquare int, name string){
	numberRow := getNumbersRow(lengRow, minSquare)
	fmt.Println(numberRow)
	writeInFile(name, numberRow)
}

func getNumbersRow (lengRow, minSquare int) (numbersRow []float64){

	var i float64 =1
	n:=0
	for n < lengRow {
		if (math.Pow(i,2) > float64(minSquare)){
			numbersRow = append (numbersRow, i)
			n++
		}
		i++
	}
	return numbersRow
}
func writeInFile(name string, numbersRow []float64)(error){

	file, err := os.Create(name)
	if err !=nil {
	return errors.New("Error with create file")
	}
	defer file.Close()
	var str string
	for _,s  :=range numbersRow{
		str += strconv.FormatFloat(s, 'f', 0, 64)
		str += ","
	}
	file.WriteString(str[:len(str)-1])
	return nil

}
