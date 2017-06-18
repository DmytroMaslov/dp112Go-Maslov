//TODO:Добавить валидацию параметров, проверку на то что треуголькник
package task3

import (
	"math"
	"sort"
)

func SortTriagle(triaglSlice []Triagle) *[]string{

	tl:=make([]triagleList, len(triaglSlice))
	i:=0
	for _,t:=range triaglSlice {
		tl[i] = triagleList{t, t.GetArea()}
		i++
	}
	sort.Sort(sort.Reverse(byArea(tl)))
	name:=make([]string, len(tl))

	i = 0
	for _, v :=range tl{
		name[i] = v.Name
		i++
	}
	return &name
}

type Triagle struct {
	Name string
	A float64
	B float64
	C float64
}

func (t *Triagle) GetArea () float64{
	p := (t.A+t.B+t.C)/2
	return math.Sqrt(p*(p-t.A)*(p-t.B)*(p-t.C))
}
type triagleList struct {
	//если наследоваться через ссылку *Triagle то в ответе только последний треугольник
	Triagle
	area float64
}

type byArea []triagleList

func(a byArea) Len() int {
	return len(a)
}
func (a byArea) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a byArea) Less(i, j int) bool {
	return a[i].area < a[j].area
	}