package task1

import "fmt"
import "errors"


func Task1(l int, w int, c string) (error) {
	if l< 2 {
		return errors.New("incorect l")
	}
	if w <2{
		return errors.New("incorect w")
	}

	firstGrid(l,w,c)
	secondGrid(l, w,c )
	return nil

}
func firstGrid(l int, w int, c string ){
	for i := 0; i<l; i++{
		for j := 0; j<w; j++{
			if i%2 == 0 {
				if j%2 ==0 {
					fmt.Print(c)
				}
				if j%2==1{
					fmt.Print(" ")
				}
			}
			if i%2 == 1 {
				if j%2 == 0 {
					fmt.Print(" ")
				}
				if j%2 == 1 {
					fmt.Print(c)
				}
			}
		}
		fmt.Print("\n")
	}
}

func secondGrid (l int, w int, c string){
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