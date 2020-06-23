package main

import (
	"fmt"
	"go-algo/graph"
)

func main() {

	fmt.Printf("result %v\n", graph.CanVisitAllRooms([][]int{
		{1},
		{2},
		{3},
		{},
	}))

	fmt.Printf("result %v\n", graph.CanVisitAllRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}))

	//fmt.Printf("result %v\n", graph.GetKthPowerValue(1, 1, 1))

}
