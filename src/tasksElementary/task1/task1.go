package task1

import (
	"errors"
	"fmt"
	"encoding/json"
)
type Task1 struct {
	Name string
}
type ChessBoard struct {
	Lenght int
	Width  int
	Symbol string
}


func Run(param []byte) (string, error){
	var ch = new (ChessBoard)
	err := json.Unmarshal(param, &ch)
	if err != nil{
		return "", errors.New("can't unmarshal data")
	}
	if err = IsValid(ch.Lenght, ch.Width, ch.Symbol); err != nil{
		return "", err
	}
	return writeGrid(ch.Lenght, ch.Width, ch.Symbol), nil
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

func writeGrid (l int, w int, c string) string{
	i:= 0
	var space = " "
	var char = c
	var newLine = "\n"
	var grid = make([]string, 0)
	for i<l {
		j :=0
		for j<w {
			if i%2 == 0{
				if j%2 == 0{
					grid = append(grid, char)
				} else {
					grid = append(grid, space)
				}

			}
			if i%2 == 1 {
				if j%2 == 0{
					grid = append(grid, space)
				} else {
					grid = append(grid, char)
				}
			}
			j++
		}
		grid = append(grid, newLine)
		//fmt.Print("\n")
		i++
	}
	return fmt.Sprint(grid)
}