/*5.Счастливые билеты
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано шестизначное число, и сумма первых трёх цифр равна сумме последних трёх,
 то этот билет считается счастливым.
2. Сложный — если сумма чётных цифр билета равна сумме нечётных цифр билета, то билет считается счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их большее количество на заданном интервале.

Входные параметры:границы min и max
Выход: элемент структурного типа с информацией о победившем методе и количестве счастливых билетов для каждого способа подсчёта.
*/
package task5


const ticketSize = 6

func GetWinerTicketResult(min, max int) WinMethod{
	winer := WinMethod{easy:  easyMethod(min, max),
						hard: hardMethod(min, max)}
	if (winer.easy>winer.hard) {winer.winName = "easy"
	} else if (winer.easy<winer.hard) {winer.winName = "hard"
	} else {winer.winName = "both"}

	return winer
}

type  WinMethod struct {
winName string
easy    int
hard    int
}

func getNumbers(min, max int)[]int{
	sl := make([]int, (max+1)-min)
	i:=0
	for min != max+1 {
		sl[i] = min
		min++
		i++
	}
	return sl
}

func convertToSlice(number int) ([] int){
	sl:=make([]int, ticketSize)
	for i := ticketSize-1; i >= 0; i--{
		sl[i] = number%10
		number = number/10
	}
	return sl
}

func easyMethod(min, max int) int{
	count := 0
	for min != max+1 {
		if isLuckyEasy(min){
			count++
		}
		min++
	}
	return count
	}


func isLuckyEasy(number int) bool{
	sl := convertToSlice(number)
	lp :=0
	rp:=0
	for i:=0; i < len(sl); i++{
		if i<ticketSize/2{
			lp += sl[i]
		}
		if i>=ticketSize/2{
			rp+=sl[i]
		}
	}
	if lp == rp {return true} else{
		return false
	}
}

func hardMethod (min, max int) int{
	count := 0
	for min != max+1 {
		if isLuckyHard(min){
			count++
		}
		min++
	}
	return count
}

func isLuckyHard(number int) bool{
	sl := convertToSlice(number)
	ev :=0
	unev:=0
	for i:=0; i < len(sl); i++{
		if sl[i]%2 ==0{
			ev += sl[i]
		}
		if sl[i]%2 !=0{
			unev +=sl[i]
		}
	}
	if ev == unev {return true} else{
		return false
	}
}