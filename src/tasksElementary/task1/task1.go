package task1

import "fmt"
import "errors"


func Task1 (l int, w int, c string) {

	writeGrid(l, w, c)

}
func IsValid (l int, w int, c string) (error){
	if l< 2 {
		return errors.New("incorrect l")
	}
	if w <2{
		return errors.New("incorrect w")
	}
	if len(c) == 0{
		return errors.New("empty char symbol")
	}
	return nil
}

func writeGrid (l int, w int, c string){
	i:= 0

	for i<l {
		j :=0
		for j<w {
			if i%2 == 0{
				if j%2 == 0{
					fmt.Print(c)
				} else {
					fmt.Print(" ")
				}

			}
			if i%2 == 1 {
				if j%2 == 0{
					fmt.Print(" ")
				} else {
					fmt.Print(c)
				}
			}
			j++
		}
		fmt.Print("\n")
		i++
	}
}