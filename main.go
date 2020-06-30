package main

import (
	"fmt"
	"go-algo/arrays"
)

func main() {

	//t := design.NewTicTacToe(2)
	res := arrays.ColumnWith1([][]int{
		{0, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}) //x
	fmt.Printf("result %v %v\n", res, "")
	//res, err = t.Play(0,1) //0
	//fmt.Printf("result %v %v\n", res, err)
	//res, err = t.Play(1, 1) //x
	//fmt.Printf("result %v %v\n", res, err)
	//res, err = t.Play(0, 1) //o
	//fmt.Printf("result %v %v\n", res, err)
	//res, err = t.Play(1, 0) //o
	//fmt.Printf("result %v %v\n", res, err)

}
