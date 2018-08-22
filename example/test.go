package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/lroc/talib"
)

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
func testFloat() {
	strNum := "3.845"
	num, err := strconv.ParseFloat(strNum, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%f\n", num)
	fmt.Printf("%T\n", num)
	var num1 float64
	num1 = Round(3845.0/1000, 1)
	fmt.Printf("%f\n", num1)
	fmt.Printf("%T\n", num1)
}
func mytime() {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	start := "2008-10-01"
	end := "2018-01-01"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	layout = "2006-01-02"
	t_start, err := time.Parse(layout, start)

	if err != nil {
		fmt.Println(err)
	}
	t_end, err := time.Parse(layout, end)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t_end)
	fmt.Println(t_end.Sub(t_start))
	days := t_end.Sub(t_start).Hours() / 24
	times := math.Floor(days / 640)
	fmt.Println(days)
	fmt.Println(times)
	a := 3
	fmt.Println(t_start.Add(time.Hour * 24 * time.Duration(a)))
}
func arrayAppend() {
	dst := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}
	src := [][]int{{11, 22, 33}, {22, 33, 44}, {33, 44, 55}, {44, 55, 66}}
	a := append(dst, src...)
	fmt.Printf("%v\n", a)
	r := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l := []int{10, 11, 12, 13, 14}
	b := append(r, l...)
	fmt.Printf("%v\n", b)
}

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
